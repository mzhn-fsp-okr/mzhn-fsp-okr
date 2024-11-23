import os
from dotenv import load_dotenv

load_dotenv()

HUGGINGFACE_MODEL = os.getenv("HUGGINGFACE_MODEL", None)

GIGACHAT_SCOPE = os.getenv("GIGACHAT_SCOPE", "GIGACHAT_API_PERS")
GIGACHAT_CREDENTIALS = os.getenv("GIGACHAT_CREDENTIALS", None)
GIGACHAT_MODEL = os.getenv("GIGACHAT_MODEL", "GigaChat-Pro")
