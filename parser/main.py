import schedule
from process_pdf import process_pdf
import time

def main():
    process_pdf()
    schedule.every().day.at("00:00", "Europe/Moscow").do(process_pdf)
    while True:
        schedule.run_pending()
        time.sleep(1)

if __name__ == "__main__":
    main()
