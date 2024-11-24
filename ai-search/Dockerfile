# ====== Stage 1: Builder ======
FROM python:3.12-slim AS builder

# Определяем версию Poetry
ARG POETRY_VERSION=1.8.3

# Устанавливаем необходимые переменные окружения
ENV POETRY_VIRTUALENVS_IN_PROJECT=true \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

# Устанавливаем системные зависимости
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Устанавливаем Poetry
RUN pip install poetry==${POETRY_VERSION}

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы Poetry
COPY pyproject.toml poetry.lock* /app/

# Аргумент для указания групп зависимостей
ARG POETRY_GROUPS=""

# Установка зависимостей с проверкой на пустую переменную
RUN if [ -n "${POETRY_GROUPS}" ]; then \
        poetry install --no-root --no-interaction --no-ansi --with ${POETRY_GROUPS}; \
    else \
        poetry install --no-root --no-interaction --no-ansi; \
    fi

# Копируем остальной код проекта
COPY . /app

# ====== Stage 2: Runtime ======
FROM python:3.12-slim

# Определяем версию Poetry (должна совпадать с билд-стадией)
ARG POETRY_VERSION=1.8.3

# Устанавливаем переменные окружения
ENV POETRY_VIRTUALENVS_IN_PROJECT=true \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

# Устанавливаем Poetry
RUN pip install poetry==${POETRY_VERSION}

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем виртуальное окружение и проект из билдера
COPY --from=builder /app/.venv /app/.venv
COPY --from=builder /app /app

# Добавляем виртуальное окружение в PATH для упрощения команд
ENV PATH="/app/.venv/bin:$PATH"

# Опционально: удаляем ненужные зависимости, если они есть

# Указываем команду для запуска приложения
CMD ["python", "main.py"]
