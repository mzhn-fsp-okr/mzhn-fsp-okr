import json
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
from logger import logging
import os
from config import FORCE_PARSE, FORCE_DOWNLOAD

camelot.logger.setLevel(logging.ERROR)

PEOPLE_DESCRIPTION_KEYWORDS = [
    "женщины",
    "юниорки",
    "девушки",
    "девочки",
    "мужчины",
    "юниоры",
    "юноши",
    "мальчики",
]


def extract_gender_age(
    description: str,
) -> List[Dict[str, Optional[List[Optional[int]]]]]:
    """
    Извлекает поло-возрастные характеристики из описания соревнования.
    """
    gender_mapping = {
        "женщины": False,
        "юниорки": False,
        "девушки": False,
        "девочки": False,
        "мужчины": True,
        "юниоры": True,
        "юноши": True,
        "мальчики": True,
    }

    desc_lower = description.lower()
    genders_found = {
        gender
        for term, gender in gender_mapping.items()
        if re.search(rf"\b{term}\b", desc_lower)
    }

    if not genders_found:
        return []

    age_patterns = [
        r"от\s+(?P<from>\d+)\s+лет(?:\s+и\s+старше)?",
        r"до\s+(?P<to>\d+)\s+лет",
        r"(?P<range_from>\d+)\s*-\s*(?P<range_to>\d+)\s+лет",
        r"(?P<exact>\d+)\s+лет(?:\s+и\s+старше)?",
    ]

    combined_pattern = "|".join(age_patterns)
    age_regex = re.compile(combined_pattern, re.IGNORECASE)

    ages_found = []
    for match in age_regex.finditer(description):
        if match.group("range_from") and match.group("range_to"):
            ages_found.append(
                [int(match.group("range_from")), int(match.group("range_to"))]
            )
        elif match.group("from"):
            ages_found.append([int(match.group("from")), None])
        elif match.group("to"):
            ages_found.append([None, int(match.group("to"))])
        elif match.group("exact"):
            age = int(match.group("exact"))
            ages_found.append([age, age])

    unique_ages = []
    seen = set()
    for age in ages_found:
        age_tuple = tuple(age)
        if age_tuple not in seen:
            unique_ages.append(age)
            seen.add(age_tuple)

    if not unique_ages:
        unique_ages = [[None, None]]

    return [
        {"gender": gender, "age": age}
        for gender in genders_found
        for age in unique_ages
    ]


def post_process(event: SportEvent):
    event.name = event.name.strip()
    event.description = event.description.strip()
    event.location = event.location.strip()

    # Инициализация переменной для хранения имени
    start_index = -1

    # Поиск первого символа, который не является строчной буквой или пробелом
    for index, char in enumerate(event.description):
        if char.islower():
            start_index = index
            break

    # Если нашли начало имени
    if start_index != -1:
        event.name += event.description[:start_index].strip()
        event.name = event.name.strip()
        # Убираем имя из описания
        event.description = event.description[start_index:].strip()
        logging.debug(f"Название события извлечено и перенесено в 'name': {event.name}")
    else:
        # Если не найдено подходящего текста
        logging.debug(
            "Не найдено подходящего текста для названия, описание остается без изменений."
        )

    event.gender_age_info = extract_gender_age(event.description)


def process_page_range(file_path, page_range: str) -> List[SportEvent]:
    logging.debug(f"Начало обработки диапазона страниц: {page_range}")
    tables = camelot.read_pdf(
        file_path,
        flavor="stream",
        pages=page_range,
        row_tol=5,
        edge_tol=100,
        columns=["107,340,467,725"],
    )
    logging.debug(f"Извлечено таблиц: {len(tables)} в диапазоне {page_range}")

    events = []
    sport_type, sport_subtype = "", ""
    page_event_counters = {}

    for table in tables:
        page_number = table.page
        if page_number not in page_event_counters:
            page_event_counters[page_number] = 0

        logging.debug(f"Обработка таблицы на странице {page_number}")

        for row in table.data:
            logging.debug(f"Данные строки: {row}")
            if len(row) < 5:
                row += [""] * (5 - len(row))
            col1, col2, col3, col4, col5 = row[:5]

            if col1.isdigit():
                page_event_counters[page_number] += 1
                current_event = SportEvent(
                    id=col1,
                    sport_type=sport_type,
                    sport_subtype=sport_subtype,
                    page_number=page_number,
                    event_order=page_event_counters[page_number],
                    description="",
                    location="",
                )
                events.append(current_event)
                logging.debug(
                    f"Создано новое событие: ID={current_event.id}, Стр={current_event.page_number}, Порядок={current_event.event_order}"
                )
            elif col1.isupper():
                sport_type = col1
                logging.debug(f"Найден новый тип спорта: {sport_type}")
                continue
            elif col1:
                sport_subtype = col1
                logging.debug(f"Найден новый подтип спорта: {sport_subtype}")
                continue

            if current_event:
                if col2:
                    if col2.isupper() and not current_event.name:
                        current_event.name += col2 + " "
                    else:
                        current_event.description += col2 + " "

                if col3:
                    if not current_event.dates.from_:
                        current_event.dates.from_ = col3
                    elif not current_event.dates.to:
                        current_event.dates.to = col3

                if col4:
                    current_event.location += col4 + " "

                if col5 and not current_event.participants:
                    try:
                        current_event.participants = int(col5)
                    except ValueError:
                        current_event.participants = 0
                        logging.warning(
                            f"Некорректное значение участников '{col5}' на странице {page_number}, событие ID {current_event.id}"
                        )

    # Очистка строковых полей
    for event in events:
        post_process(event)

    logging.debug(
        f"Обработка диапазона {page_range} завершена. Всего событий: {len(events)}"
    )
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
        },
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


def download_pdf(url: str) -> str:
    logging.info(f"Скачивание PDF по ссылке: {url}")
    response = requests.get(url, verify=False, headers={"User-Agent": "Mozilla/5.0"})
    if response.status_code == 200:
        with tempfile.NamedTemporaryFile(delete=False, suffix=".pdf") as temp_pdf:
            temp_pdf.write(response.content)
            logging.info("PDF успешно скачан")
            return temp_pdf.name
    else:
        logging.error(f"Не удалось скачать PDF. Код ошибки: {response.status_code}")
        raise Exception(f"HTTP {response.status_code}")


CACHE_FILE = "cache.json"


def load_cache():
    if os.path.exists(CACHE_FILE):
        with open(CACHE_FILE, "r", encoding="utf-8") as f:
            return json.load(f)
    return {}


def save_cache(cache):
    with open(CACHE_FILE, "w", encoding="utf-8") as f:
        json.dump(cache, f)


def process_pdf():
    file_path = None
    try:
        # Загружаем кеш
        cache = load_cache()

        # Получаем дату обновления и ссылку на PDF
        # update_date, link = get_link()
        update_date = "14.11.2024"
        link = ""

        # Проверяем, изменилась ли дата обновления
        if (
            not FORCE_DOWNLOAD
            and cache.get("update_date") == update_date
            and os.path.exists(cache.get("file_path", ""))
        ):
            logging.info("PDF не изменился с последнего скачивания.")
            file_path = cache["file_path"]

            if not FORCE_PARSE:
                return
        else:
            if os.path.exists(cache.get("file_path", "")):
                logging.info("Удаляем старый PDF файл.")
                os.remove(cache.get("file_path", ""))

            # Скачиваем новый PDF
            file_path = download_pdf(link)
            cache["update_date"] = update_date
            cache["file_path"] = file_path
            save_cache(cache)
            logging.info("Скачан новый PDF файл.")

        # Читаем PDF и определяем количество страниц
        reader = PdfReader(file_path)
        total_pages = len(reader.pages)

        # Разбиваем страницы на диапазоны для многопроцессорной обработки
        num_workers = os.cpu_count()
        page_ranges = split_pages(total_pages, num_workers)

        logging.info(f"Начата обработка. Количество потоков: {num_workers}")

        # Список всех извлеченных событий
        all_events = []
        with ProcessPoolExecutor(max_workers=num_workers) as executor:
            futures = [
                executor.submit(process_page_range, file_path, pr) for pr in page_ranges
            ]

            for future in as_completed(futures):
                try:
                    all_events.extend(future.result())
                except Exception as e:
                    logging.error(f"Ошибка при обработке диапазона страниц: {e}")

        # Сортируем события по номеру страницы и порядку
        all_events.sort(key=lambda e: (e.page_number, e.event_order))

        # Постобработка для заполнения отсутствующих типов и подтипов спорта
        last_sport_type, last_sport_subtype = "", ""
        for event in all_events:
            if not event.sport_type:
                event.sport_type = last_sport_type
            else:
                last_sport_type = event.sport_type

            if not event.sport_subtype:
                event.sport_subtype = last_sport_subtype
            else:
                last_sport_subtype = event.sport_subtype

        # Преобразование событий в словари для JSON
        events_dict = [sport_event_to_dict(event) for event in all_events]

        # Сохранение в JSON файл
        with open("results.json", "w", encoding="utf-8") as json_file:
            json.dump(events_dict, json_file, ensure_ascii=False, indent=4)

        # Отправляем события через gRPC
        send_events_via_grpc(all_events)
        logging.info(f"Всего отправлено событий: {len(all_events)}")
    except Exception as e:
        logging.error(f"Общая ошибка в процессе обработки PDF: {e}")
