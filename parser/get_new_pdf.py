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


# Настройка Selenium
def init_driver():
    options = Options()
    options.add_argument("--headless=new")
    return webdriver.Chrome(
        service=Service(ChromeDriverManager().install()), options=options
    )


# Открытие URL
def open_page(driver, url):
    logging.info(f"Открываем страницу: {url}")
    driver.get(url)


# Поиск заголовка "II часть ЕКП"
def find_ii_part_header(driver):
    try:
        return WebDriverWait(driver, 10).until(
            EC.presence_of_element_located(
                (By.XPATH, "//h2[contains(text(), 'II часть ЕКП')]")
            )
        )
    except Exception as e:
        logging.error(f"Не удалось найти заголовок 'II часть ЕКП': {e}")
        driver.quit()
        exit()


# Поиск контейнера для "II часть ЕКП"
def find_ii_part_container(header_element):
    try:
        return header_element.find_element(
            By.XPATH, "./ancestor::div[@class='wrapper']"
        )
    except Exception as e:
        logging.error(f"Не удалось найти контейнер для 'II часть ЕКП': {e}")
        exit()


# Поиск всех каталогов
def find_year_containers(container):
    try:
        return container.find_elements(By.CSS_SELECTOR, "div.cursor-pointer")
    except Exception as e:
        logging.error(f"Не удалось найти годовые каталоги: {e}")
        exit()


# Обработка каталога
def process_catalog(container, index):
    try:
        is_expanded = False
        try:
            container.find_element(
                By.XPATH,
                "./following-sibling::div[contains(@class, 'select-anymation-show')]",
            )
            is_expanded = True
            logging.info(f"Каталог {index} уже развернут")
        except:
            pass

        if not is_expanded:
            try:
                container.click()
                logging.info(f"Разворачиваем каталог {index}")
                time.sleep(3)
            except Exception as e:
                logging.error(f"Не удалось развернуть каталог {index}: {e}")

        return container.find_elements(By.CSS_SELECTOR, "div.file-item")
    except Exception as e:
        logging.error(f"Ошибка при обработке каталога {index}: {e}")
        return []


# Извлечение ссылок из файлов
def extract_link(file_items, index):
    pdf_links = set()
    for file_index, file_item in enumerate(file_items, start=1):
        try:
            text_elem = file_item.find_element(By.TAG_NAME, "p")

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
                    logging.info(f"Найдена ссылка: {href}")
                    return update_date, href

        except Exception as e:
            logging.error(
                f"Ошибка при обработке file-item {file_index} в каталоге {index}: {e}"
            )
    return pdf_links


# Сохранение ссылок в файл
def save_links(pdf_links, filename):
    with open(filename, "w", encoding="utf-8") as f:
        for link in pdf_links:
            f.write(f"{link}\n")
    logging.info(f"Ссылки сохранены в файл '{filename}'")


def get_link():
    driver = init_driver()
    url = "https://www.minsport.gov.ru/activity/government-regulation/edinyj-kalendarnyj-plan/"

    open_page(driver, url)

    ii_part_header = find_ii_part_header(driver)
    ii_part_container = find_ii_part_container(ii_part_header)

    year_containers = find_year_containers(ii_part_container)
    container = year_containers[0]

    file_items = process_catalog(container, 1)
    update_date, link = extract_link(file_items, 1)

    return update_date, link


# Основной процесс
def main():
    driver = init_driver()
    url = "https://www.minsport.gov.ru/activity/government-regulation/edinyj-kalendarnyj-plan/"
    pdf_links = set()

    try:
        open_page(driver, url)

        ii_part_header = find_ii_part_header(driver)
        ii_part_container = find_ii_part_container(ii_part_header)

        year_containers = find_year_containers(ii_part_container)
        logging.info(
            f"Найдено {len(year_containers)} годовых каталогов в 'II части ЕКП'"
        )

        year_containers = year_containers[:1]  # Ограничиваем для теста

        for index, container in enumerate(year_containers, start=1):
            file_items = process_catalog(container, index)
            links = extract_link(file_items, index)
            pdf_links.update(links)

        save_links(pdf_links, "pdf_links_II_ECP.txt")

    finally:
        driver.quit()
        logging.info("Браузер закрыт")


if __name__ == "__main__":
    main()
