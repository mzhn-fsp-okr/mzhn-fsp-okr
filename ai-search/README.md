# Конфигурация проекта `ai-search`

## :gear: Библиотеки

- **Langchain-Core** - Библиотека для создания интеллектуальных цепочек и управления потоками обработки данных.
- **Langchain-Gigachat** - Интеграция Langchain с GigaChat.
- **Langserve** - Утилита для развертывания моделей LangChain как веб-сервисов.
- **FastAPI** - Современный инструмент для создания API.
- **Langchain-Community** - Пакет для использования и настройки моделей из сообщества Langchain.
- **Lark** - Инструмент для создания парсеров грамматик.
- **SSE-Starlette** - Поддержка Server-Sent Events для приложений FastAPI.
- **Uvicorn** - Высокопроизводительный сервер для FastAPI.

### Опциональные зависимости
- **Langchain-Ollama** - Интеграция Langchain с Ollama.
- **Ollama** - Для работы с моделями Ollama.
- **Sentence-Transformers** - Библиотека для использования моделей HuggingFace.
- **Langchain-Huggingface** - Интеграция моделей Huggingface с Langchain.

### Dev-зависимости
- **Black** - Инструмент для форматирования Python-кода.
- **Matplotlib** - Для визуализации данных.
- **IPyKernel** - Интеграция с Jupyter Notebook.

---

## :hammer_and_wrench: Структура проекта

```
.
├── chain.py
├── chat_model.py
├── config.py
├── Dockerfile
├── main.py
├── poetry.lock
├── pyproject.toml
└── README.md
```

### Конфигурация

- **HUGGINGFACE_MODEL**:
  - **Описание**: ID модели HuggingFace, используемой в проекте.
  - **Значение по умолчанию**: `None`.

- **OLLAMA_MODEL**:
  - **Описание**: Название модели Ollama для использования.
  - **Значение по умолчанию**: `None`.

- **OLLAMA_BASE_URL**:
  - **Описание**: Базовый URL-адрес API Ollama.
  - **Значение по умолчанию**: `None`.

- **GIGACHAT_SCOPE**:
  - **Описание**: Область доступа для использования GigaChat API.
  - **Значение по умолчанию**: `"GIGACHAT_API_PERS"`.

- **GIGACHAT_MODEL**:
  - **Описание**: Название модели GigaChat для работы.
  - **Значение по умолчанию**: `"GigaChat"`.

- **GIGACHAT_CREDENTIALS**:
  - **Описание**: Путь к файлу или данные для аутентификации в GigaChat.
  - **Значение по умолчанию**: `None`. 
