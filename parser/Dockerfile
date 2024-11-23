FROM python:3.12-slim AS builder

ENV POETRY_VIRTUALENVS_IN_PROJECT=true \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    curl \
    && rm -rf /var/lib/apt/lists/*

RUN pip install poetry==1.8.3

WORKDIR /app

COPY vendored/*-linux_x86_64.whl /app/vendored/

COPY pyproject.toml poetry.lock* /app/

RUN poetry install --no-root --no-interaction --no-ansi

COPY . /app

# ====== Stage 2: Runtime ======
FROM python:3.12-slim

ENV POETRY_VIRTUALENVS_IN_PROJECT=true \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

RUN apt-get update && apt-get install -y --no-install-recommends \
    gnupg2 \
    wget \
    curl

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list'

RUN apt-get update && \
    apt-get install -y --no-install-recommends google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

RUN pip install get-chrome-driver==1.3.20 && get-chrome-driver --auto-download --extract

RUN pip install poetry==1.8.3

WORKDIR /app

COPY --from=builder /app/.venv /app/.venv

COPY --from=builder /app /app

CMD ["poetry", "run", "python", "main.py"]
