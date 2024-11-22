import schedule
import time
from logger import logging
from process_pdf import process_pdf


def run_task():
    """
    Выполняет основную задачу с обработкой ошибок.
    """
    try:
        logging.info("Запуск задачи process_pdf")
        process_pdf()
        logging.info("Задача process_pdf выполнена успешно")
    except Exception as e:
        logging.error(f"Ошибка при выполнении process_pdf: {e}", exc_info=True)

def main():
    """
    Основная функция запуска планировщика задач.
    """    
    schedule.every().day.at("00:00").do(run_task)

    logging.info("Планировщик запущен. Ожидание задач...")

    while True:
        try:
            # Выполняем запланированные задачи
            schedule.run_pending()
        except Exception as e:
            logging.error(f"Ошибка в планировщике: {e}", exc_info=True)
        time.sleep(1)

if __name__ == "__main__":
    main()
