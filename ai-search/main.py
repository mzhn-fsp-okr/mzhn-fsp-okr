import os
from fastapi import FastAPI
from langserve import add_routes
from langchain.chains.query_constructor.schema import AttributeInfo
from langchain_huggingface import ChatHuggingFace, HuggingFacePipeline
from langchain.chains.query_constructor.base import (
    StructuredQueryOutputParser,
    get_query_constructor_prompt,
)
from langchain_core.prompts import ChatPromptTemplate
from langchain_gigachat import GigaChat
from langchain_core.messages.ai import AIMessage
from langchain_core.output_parsers import JsonOutputParser
from langchain_gigachat import GigaChat

metadata_field_info = [
    AttributeInfo(
        name="gender",
        description="Пол. Вернуть True, если мужской",
        type="boolean",
    )
]

app = FastAPI()


chat_model = GigaChat(
    scope=os.getenv("GIGACHAT_SCOPE", "GIGACHAT_API_PERS"),
    credentials=os.getenv("GIGACHAT_CREDENTIALS"),
    model="GigaChat",
    verify_ssl_certs=False
)

# llm = HuggingFacePipeline.from_model_id(
#     model_id="Vikhrmodels/Vikhr-Qwen-2.5-0.5b-Instruct",
#     task="text-generation",
#     pipeline_kwargs=dict(
#         max_new_tokens=512,
#         temperature=0.7,
#         do_sample=True,
#     ),
# )
# 
# def qwen_fix(message: AIMessage) -> AIMessage:
#     message.content = message.content.split("<|im_start|>assistant\n")[-1]
#     return message
# 
# chat_model = ChatHuggingFace(llm=llm) | qwen_fix


SYSTEM_TEMPLATE="""Извлеки из вопроса пользователя необходимые поля, указанные в структуре ниже.
Поля могут быть указаны или отсутствовать. Если поле отсутствует, его значение должно быть null.
Верни ответ в формате JSON. Никаких пояснений, только JSON.

Описание полей: 
- "name" - название события: КУБОК, ПЕРВЕНСТВО, ЧЕМПИОНАТ и другое.
- "sex" - укажи true для мужчин/мальчиков/юниоров/юношей, false для женщин/девочек/юниоров/юниорок.
- "min_age" и "max_age" - целые числа, границы возраста. Если указан точный возраст, а не диапазон (от/до) - указывать одинаковое значение.
Если написано только от,старше и прочее - указывай только min_age.
Если написано только до,младше и прочее - указывай только max_age.
- "start_date" и "end_date" - строки, в формате DD.MM.YYYY. Временной промежуток для поиска.
Если явно не указано start_date - сегодняшний день. Сегодняшний день указан в контексте.
Если явно не указано end_date - null.

Если поля не указаны - необходимо их вернуть со значением null.

Проверь всё несколько раз.


ПРИМЕРЫ

Вопрос: "Привет! Что ты можешь подсказать для пловца в ближайшее время?"
Ответ: {{"start_date":"23.11.2024","end_date": null,"name":null,"sport":"плавание","sex":true,"min_age": null,"max_age": null}}
Вопрос: "Какие есть чемпионаты по скалолазанию для девочки 10 лет в ближайший месяц?"
Ответ: {{"start_date":"23.11.2024","end_date":"23.12.2024","name":"ЧЕМПИОНАТ","sport":"скалолазание","sex":false,"min_age":10,"max_age":10}}
Вопрос: "Мероприятия на завтра"
Ответ: {{"start_date":"24.11.2024","end_date": null,"name":null,"sport":null,"sex":null,"min_age":null,"max_age":null}}
Вопрос: "Соревнование для борца от 18 лет"
Ответ: {{"start_date":"23.11.2024","end_date":null,"name":"СОРЕВНОВАНИЯ","sport":"борьба","sex":true,"min_age":18,"max_age":null}}

КОНТЕКСТ

Сегодня: 23.11.2024
"""
# 
# SYSTEM_TEMPLATE="""You are an assistant that transforms user queries into a structured JSON format for searching competitions. 
# Do not provide any explanations or answers other than the JSON structure specified below.
# 
# {{
#     "sportType": string | null     // Name of the sport with stemming. If not specified, use null.
#     "gender": boolean   | null     // true если мужчина, false если женщина. If not specified, use null.
#     "minAge": integer   | null     // Minimum age of participants. If unknown, use null.
#     "maxAge": integer   | null     // Maximum age of participants. If unknown, use null.
#     "location": string  | null     // Geographical locations for the search. If unknown, use null.
# }}
# 
# Инструкции:
# - Всегда трижды проверяйте правильность вашего ответа.
# - sportType содержит конкретный тип спорта, а не обобщенно
# - **Лучше указывать null значения вместо неправильных **.
# - Используйте информацию только из пользовательского запроса
# - Не предоставляйте никаких объяснений или ответов, кроме JSON
# - Правильно указывай gender
# 
# **Примеры**:
# ---
# User: Тренеровки по плаванию до 10 лет мальчик в Крыму
# Assistant:
# {{
#     "sportType": "плаван",
#     "gender": true,
#     "minAge": null,
#     "maxAge": 10,
#     "location": "КРЫМ",
# }}
# --- 
# User: Первенство федерального округа ДФО для женщины по борьбе. Ей 10 лет (молодежный состав).
# Assistant:
# {{
#     "sportType": "борьб",
#     "gender": false,
#     "minAge": 10,
#     "maxAge": 10,
#     "location": "РОССИЯ"
# }}
# ---
# User: Кубок Европы для юниорки от 18 лет по конному спорту. 
# Assistant:
# {{    
#     "sportType": "кон спор",
#     "gender": false,
#     "minAge": 18,
#     "maxAge": null,
#     "location": null
# }}
# ---
# User: Соревнования в Саратове
# Assistant:
# {{    
#     "sportType": null,
#     "gender": null,
#     "minAge": null,
#     "maxAge": null,
#     "location": "САРАТОВ"
# }}
# ---
# 
# A lot depends on this answer—triple-check it!
# """

prompt_template = ChatPromptTemplate.from_messages([
    ('system', SYSTEM_TEMPLATE),
    ('user', '{input}')
])

parser = JsonOutputParser()

query_constructor = prompt_template | chat_model | parser

add_routes(
    app,
    query_constructor,
    path="/chat",
)

import uvicorn

if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)