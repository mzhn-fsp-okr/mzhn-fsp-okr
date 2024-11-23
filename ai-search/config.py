import os
from dotenv import load_dotenv

load_dotenv()

HUGGINGFACE_MODEL = os.getenv("HUGGINGFACE_MODEL", None)

# В данной задаче rscr/vikhr_nemo_12b:Q4_K_M на RX 6750 XT хорошо себя показала
OLLAMA_MODEL = os.getenv("OLLAMA_MODEL", None)

GIGACHAT_SCOPE = os.getenv("GIGACHAT_SCOPE", "GIGACHAT_API_PERS")
GIGACHAT_MODEL = os.getenv("GIGACHAT_MODEL", "GigaChat")
GIGACHAT_CREDENTIALS = os.getenv("GIGACHAT_CREDENTIALS", None)
