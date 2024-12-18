from langchain_core.language_models import BaseChatModel

from config import (
    HUGGINGFACE_MODEL,
    
    GIGACHAT_SCOPE,
    GIGACHAT_CREDENTIALS,
    GIGACHAT_MODEL,
    
    OLLAMA_MODEL,
    OLLAMA_BASE_URL
)

chat: BaseChatModel = None

if HUGGINGFACE_MODEL:
    from langchain_huggingface import ChatHuggingFace, HuggingFacePipeline

    llm = HuggingFacePipeline.from_model_id(
        model_id=HUGGINGFACE_MODEL,
        task="text-generation",
        pipeline_kwargs=dict(
            max_new_tokens=1024,
            temperature=0.1,
            do_sample=True,
        ),
    )

    chat = ChatHuggingFace(llm=llm)

elif OLLAMA_MODEL:
    from langchain_ollama import ChatOllama

    chat = ChatOllama(
        model=OLLAMA_MODEL,
        temperature=0.1,
        max_tokens=1024,
        base_url=OLLAMA_BASE_URL
    )

else:
    from langchain_gigachat import GigaChat

    chat = GigaChat(
        scope=GIGACHAT_SCOPE,
        credentials=GIGACHAT_CREDENTIALS,
        model=GIGACHAT_MODEL,
        verify_ssl_certs=False,
    )
