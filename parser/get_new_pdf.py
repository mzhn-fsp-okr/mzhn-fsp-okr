import re
import time
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

from logger import logging


def init_driver():
    """
    Инициализация WebDriver с настройками для работы в фоновом режиме.
    """
    options = Options()
    options.add_argument("--window-size=1920,1080")
    options.add_argument("--headless=new")
    logging.debug("Инициализация WebDriver")
    return webdriver.Chrome(
        service=Service(ChromeDriverManager().install()), options=options
    )


def open_page(driver, url):
    """
    Открытие страницы в браузере.
    :param driver: Selenium WebDriver.
    :param url: URL страницы.
    """
    logging.debug(f"Открываем страницу: {url}")
    driver.get(url)


# Поиск заголовка "II часть ЕКП"
def find_ii_part_header(driver):
    """
    Поиск заголовка "II часть ЕКП" на странице.
    :param driver: Selenium WebDriver.
    :return: WebElement заголовка.
    """
    try:
        logging.debug("Ищем заголовок 'II часть ЕКП'")
        return WebDriverWait(driver, 10).until(
            EC.presence_of_element_located(
                (By.XPATH, "//h2[contains(text(), 'II часть ЕКП')]")
            )
        )
    except Exception as e:
        logging.error(f"Не удалось найти заголовок 'II часть ЕКП': {e}")
        driver.quit()


# Поиск контейнера для "II часть ЕКП"
def find_ii_part_container(header_element):
    """
    Поиск контейнера, связанного с заголовком "II часть ЕКП".
    :param header_element: WebElement заголовка.
    :return: WebElement контейнера.
    """
    try:
        logging.debug("Ищем контейнер для 'II часть ЕКП'")
        return header_element.find_element(
            By.XPATH, "./ancestor::div[@class='wrapper']"
        )
    except Exception as e:
        logging.error(f"Не удалось найти контейнер для 'II часть ЕКП': {e}")


# Поиск всех каталогов
def find_year_containers(container):
    """
    Поиск элементов, представляющих годовые каталоги.
    :param container: WebElement основного контейнера.
    :return: Список WebElements для годовых каталогов.
    """
    try:
        logging.debug("Ищем годовые каталоги")
        return container.find_elements(By.CSS_SELECTOR, "div.cursor-pointer")
    except Exception as e:
        logging.error(f"Не удалось найти годовые каталоги: {e}")
        exit()


def process_catalog(container, index):
    """
    Обработка одного каталога: проверка и разворачивание.
    :param container: WebElement каталога.
    :param index: Индекс каталога.
    :return: Список WebElements файлов в каталоге.
    """
    try:
        logging.debug(f"Обрабатываем каталог {index}")
        is_expanded = False

        # Проверяем, развернут ли каталог
        try:
            container.find_element(
                By.XPATH,
                "./following-sibling::div[contains(@class, 'select-anymation-show')]",
            )
            is_expanded = True
            logging.debug(f"Каталог {index} уже развернут")
        except:
            pass

        # Если не развернут, кликаем для разворачивания
        if not is_expanded:
            try:
                container.click()
                logging.debug(f"Разворачиваем каталог {index}")
                time.sleep(3)  # Ожидание завершения анимации
            except Exception as e:
                logging.error(f"Не удалось развернуть каталог {index}: {e}")

        # Возвращаем элементы файлов в каталоге
        return container.find_elements(By.CSS_SELECTOR, "div.file-item")
    except Exception as e:
        logging.error(f"Ошибка при обработке каталога {index}: {e}")
        return []


# Извлечение ссылок из файлов
def extract_link(file_items, index):
    """
    Извлечение ссылки на PDF из элементов файлов.
    :param file_items: Список WebElements файлов.
    :param index: Индекс каталога.
    :return: Дата обновления и ссылка на PDF, если найдены.
    """
    for file_index, file_item in enumerate(file_items, start=1):
        try:
            text_elem = file_item.find_element(By.TAG_NAME, "p")

            # Проверяем текст на соответствие требуемому шаблону
            if (
                "Единый календарный план межрегиональных, всероссийских и международных физкультурных мероприятий и спортивных мероприятий"
                not in text_elem.text
            ):
                continue

            text = text_elem.text
            date_pattern = r"\b\d{2}\.\d{2}\.\d{4}\b"
            dates = re.findall(date_pattern, text)

            if not dates:
                continue

            update_date = dates[0]

            for link in file_item.find_elements(By.TAG_NAME, "a"):
                href = link.get_attribute("href")
                if href and href.endswith(".pdf"):
                    logging.debug(f"Найдена ссылка: {href}")
                    return update_date, href

        except Exception as e:
            logging.error(
                f"Ошибка при обработке file-item {file_index} в каталоге {index}: {e}"
            )
    return None, None


# Сохранение ссылок в файл
def save_links(pdf_links, filename):
    with open(filename, "w", encoding="utf-8") as f:
        for link in pdf_links:
            f.write(f"{link}\n")
    logging.info(f"Ссылки сохранены в файл '{filename}'")


def get_link():
    """
    Главная функция для извлечения последней ссылки PDF из страницы.
    :return: Дата обновления и ссылка на PDF.
    """
    driver = init_driver()
    url = "https://www.minsport.gov.ru/activity/government-regulation/edinyj-kalendarnyj-plan/"

    try:
        open_page(driver, url)

        # Ищем заголовок и связанный контейнер
        ii_part_header = find_ii_part_header(driver)
        ii_part_container = find_ii_part_container(ii_part_header)

        # Находим первый контейнер и обрабатываем его
        year_containers = find_year_containers(ii_part_container)
        container = year_containers[0]

        # Извлекаем файлы и получаем ссылку
        file_items = process_catalog(container, 1)
        update_date, link = extract_link(file_items, 1)

        return update_date, link
    finally:
        driver.quit()
