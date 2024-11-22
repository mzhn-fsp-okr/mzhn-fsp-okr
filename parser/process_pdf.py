from get_new_pdf import get_link
import requests
import tempfile
from PyPDF2 import PdfReader
from concurrent.futures import ProcessPoolExecutor, as_completed
from typing import Dict, List, Optional
import camelot
import re
from lib import SportEvent, sport_event_to_dict
from grpc_communitcation import send_events_via_grpc

def extract_gender_age(description: str) -> List[
    Dict[str, Optional[List[Optional[int]]]]
]:
    """
    Извлекает поло-возрастные характеристики из описания соревнования.

    :param description: Строка с описанием соревнования.
    :return: Список словарей с ключами 'gender' и 'age'.
             'gender' принимает значения "ЖЕН" или "МУЖ",
             'age' представляет собой список [min_age, max_age], где min_age и max_age могут быть целыми числами или None.
    """
    gender_mapping = {
        'женщины': False,
        'юниорки': False,
        'девушки': False,
        'девочки': False,
        
        'мужчины': True,
        'юниоры': True,
        'юноши': True,
        'мальчики': True,
    }

    # Приведение описания к нижнему регистру для унификации
    desc_lower = description.lower()

    # Извлечение гендерных терминов
    genders_found = set()
    for term, gender in gender_mapping.items():
        # Используем границы слова для точного соответствия
        if re.search(r'\b' + re.escape(term) + r'\b', desc_lower):
            genders_found.add(gender)
    
    # Если гендерные категории не найдены, возвращаем пустой список
    if not genders_found:
        return []

    # Определение регулярных выражений для возрастных категорий с именованными группами
    age_patterns = [
        r'от\s+(?P<from>\d+)\s+лет(?:\s+и\s+старше)?',  # от X лет или от X лет и старше
        r'до\s+(?P<to>\d+)\s+лет',                     # до X лет
        r'(?P<range_from>\d+)\s*-\s*(?P<range_to>\d+)\s+лет',  # X-Y лет
        r'(?P<exact>\d+)\s+лет(?:\s+и\s+старше)?',        # X лет или X лет и старше
    ]

    # Объединение всех шаблонов в одно регулярное выражение
    combined_pattern = '|'.join(age_patterns)
    age_regex = re.compile(combined_pattern, re.IGNORECASE)

    # Поиск всех возрастных категорий в описании
    ages_found = []
    for match in age_regex.finditer(description):
        if match.group('range_from') and match.group('range_to'):
            # Формат X-Y лет
            min_age = int(match.group('range_from'))
            max_age = int(match.group('range_to'))
            ages_found.append([min_age, max_age])
        elif match.group('from'):
            # Формат 'от X лет' или 'от X лет и старше'
            min_age = int(match.group('from'))
            ages_found.append([min_age, None])
        elif match.group('to'):
            # Формат 'до X лет'
            max_age = int(match.group('to'))
            ages_found.append([None, max_age])
        elif match.group('exact'):
            # Формат 'X лет' или 'X лет и старше'
            age = int(match.group('exact'))
            ages_found.append([age, age])
    
    # Удаление дубликатов
    unique_ages = []
    seen = set()
    for age in ages_found:
        age_tuple = tuple(age)
        if age_tuple not in seen:
            unique_ages.append(age)
            seen.add(age_tuple)
    
    # Если возрастные категории не найдены, устанавливаем [None, None]
    if not unique_ages:
        unique_ages = [[None, None]]
    
    # Создание списка словарей с гендером и возрастом
    result = []
    for gender in genders_found:
        for age in unique_ages:
            result.append({
                "gender": gender,
                "age": age
            })
    
    return result

def process_page_range(file_path, page_range: str) -> List[SportEvent]:
    # print(f"[Process {page_range}] Начало обработки диапазона страниц")
    tables = camelot.read_pdf(
        file_path,
        flavor='stream', 
        pages=page_range,
        row_tol=10, 
        edge_tol=5000
    )
    # print(f"[Process {page_range}] Извлечено таблиц: {len(tables)}")
    
    events = []
    current_event = None
    sport_type, sport_subtype = '', ''
    
    # Словарь для отслеживания порядка событий на каждой странице
    page_event_counters = {}

    for table in tables:
        page_number = table.page  # Получаем номер страницы
        if page_number not in page_event_counters:
            page_event_counters[page_number] = 0  # Инициализируем счётчик для страницы
        
        # print(f"[Process {page_range}] Обработка таблицы на странице {page_number}")
        
        for row in table.data:
            # print(f"[Process {page_range}] Данные строки: {row}")
            if len(row) < 5:
                row += [''] * (5 - len(row))
            col1, col2, col3, col4, col5 = row[:5]
            if col1.isdigit():
                page_event_counters[page_number] += 1  # Увеличиваем порядок события
                current_event = SportEvent(
                    id=col1,
                    sport_type=sport_type,
                    sport_subtype=sport_subtype,
                    page_number=page_number,
                    event_order=page_event_counters[page_number],
                    description='',
                    location=''
                )
                events.append(current_event)
                # print(f"[Process {page_range}] Создано новое событие: ID={current_event.id}, Стр={current_event.page_number}, Порядок={current_event.event_order}")
            elif col1.isupper():
                sport_type = col1
                # print(f"[Process {page_range}] Найден новый тип спорта: {sport_type}")
                continue
            elif col1:
                sport_subtype = col1
                # print(f"[Process {page_range}] Найден новый подтип спорта: {sport_subtype}")
                continue
            
            if current_event:
                if col2:
                    if col2.isupper() and not current_event.name:
                        current_event.name += col2 + " "
                    else:
                        current_event.description += col2 + " "

                if col3:
                    if current_event.dates.from_:
                        current_event.dates.to = col3
                    else:
                        current_event.dates.from_ = col3

                if col4:
                    current_event.location += col4 + " "

                if col5 and not current_event.participants:
                    try:
                        current_event.participants = int(col5)
                    except ValueError:
                        current_event.participants = 0
                        print(f"[Process {page_range}] Предупреждение: Некорректное значение участников '{col5}' на странице {page_number}, событие ID {current_event.id}")

    # Очистка строковых полей
    for event in events:
        event.name = event.name.strip()
        event.description = event.description.strip()
        event.location = event.location.strip()
        event.gender_age_info = extract_gender_age(event.description)

    # print(f"[Process {page_range}] Обработка завершена. Всего событий: {len(events)}")
    return events

def downlod_pdf(url):
    response = requests.get(
        url,
        verify=False,
        headers={
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
            "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
            "Accept-Language": "en-US,en;q=0.5",
            "Accept-Encoding": "gzip, deflate, br",
            "Connection": "keep-alive",
            "Upgrade-Insecure-Requests": "1",
            "Sec-Fetch-Dest": "document",
            "Sec-Fetch-Mode": "navigate",
            "Sec-Fetch-Site": "none",
            "Sec-Fetch-User": "?1",
        }
    )
    if response.status_code == 200:
        with tempfile.NamedTemporaryFile(delete=False, suffix=".pdf") as temp_pdf:
            temp_pdf.write(response.content)
            return temp_pdf.name
    else:
        print(f"Не удалось скачать файл. Код ошибки: {response.status_code}")

def split_pages(total_pages: int, num_workers: int) -> List[str]:
    pages_per_worker = total_pages // num_workers
    page_ranges = []
    for i in range(num_workers):
        start = i * pages_per_worker + 1
        if i == num_workers - 1:
            end = total_pages
        else:
            end = (i + 1) * pages_per_worker
        page_ranges.append(f"{start}-{end}")
    return page_ranges

def process_pdf():
    # _, link = get_link()

    # file_path = downlod_pdf(link)

    file_path = '/home/maxim/dev/mzhn-chmp/parser/test.pdf'

    reader = PdfReader(file_path)
    total_pages = len(reader.pages)

    num_workers = 12
    page_ranges = split_pages(total_pages, num_workers)

    all_events = []

    with ProcessPoolExecutor(max_workers=num_workers) as executor:
        # print("Отправка задач в пул процессов...")
        futures = [
            executor.submit(
                process_page_range, file_path, pr
            ) for pr in page_ranges
        ]

        for future in as_completed(futures):
            try:
                events = future.result()
                print(f"Главный процесс: Получено {len(events)} событий из задачи")
                all_events.extend(events)
            except Exception as e:
                print(f"Ошибка при обработке диапазона страниц: {e}")

    print(f"Всего собрано событий: {len(all_events)}")
    print("Сортировка событий по номеру страницы и порядку...")
    # Сортировка событий по номеру страницы и порядку на странице
    all_events.sort(key=lambda e: (e.page_number, e.event_order))

    # Постобработка для заполнения отсутствующих sport_type и sport_subtype
    print("Начало постобработки для заполнения отсутствующих типов спорта...")
    last_sport_type = ''
    last_sport_subtype = ''
    for event in all_events:
        if not event.sport_type:
            event.sport_type = last_sport_type
        else:
            last_sport_type = event.sport_type
        
        if not event.sport_subtype:
            event.sport_subtype = last_sport_subtype
        else:
            last_sport_subtype = event.sport_subtype

    send_events_via_grpc(all_events)
