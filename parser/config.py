import os
from dotenv import load_dotenv

load_dotenv()

# Адрес event-service
EVENT_SERVICE_ADDRESS = os.getenv("EVENT_SERVER_ADDRESS")

# Уровень логирования
LOG_LEVEL = os.getenv("LOG_LEVEL", "ERROR").upper()

# Количество потоков при парсинге pdf
WORKERS_COUNT = int(os.getenv("WORKERS_COUNT", os.cpu_count()))

# Выполнять парсинг pdf при запуске приложения
PROCESS_ON_STARTUP = os.getenv("PROCESS_ON_STARTUP", "FALSE").upper() == "TRUE"

# Парсить, даже если есть в cache
FORCE_PARSE = os.getenv("FORCE_PARSE", "FALSE").upper() == "TRUE"

# Скачивать, даже если есть в cache
FORCE_DOWNLOAD = os.getenv("FORCE_DOWNLOAD", "FALSE").upper() == "TRUE"