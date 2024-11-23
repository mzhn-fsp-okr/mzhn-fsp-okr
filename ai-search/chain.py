from datetime import datetime
from dateutil.relativedelta import *
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.output_parsers import JsonOutputParser
from langchain_core.runnables import RunnableParallel, RunnablePassthrough
from pydantic import BaseModel

from chat_model import chat

SYSTEM_TEMPLATE = """Извлеки из вопроса пользователя необходимые поля, указанные в структуре ниже.
Поля могут быть указаны или отсутствовать. 

Верни ответ в формате JSON. Никаких пояснений, только JSON.
Описание полей: 
- "name" - название события: КУБОК, ПЕРВЕНСТВО, ЧЕМПИОНАТ, СОРЕВНОВАНИЕ и другое. Одно слово.
- "sex" - укажи true для мужчин/мальчиков/юниоров/юношей, false для женщин/девочек/юниоров/юниорок. 
Если нельзя узнать из вопроса (например, дети/люди) - null
- "min_age" и "max_age" - целые числа, границы возраста. 
Если указан точный возраст, а не диапазон (от/до) - указывать одинаковое значение.
Если написано только от,старше и прочее - указывай только min_age.
Если написано только до,младше и прочее - указывай только max_age.
- "start_date" и "end_date" — строки в формате DD.MM.YYYY, представляющие временной промежуток для поиска. В запросе могут содержаться относительные даты, указанные словами, например: "полугодие", "квартал" и другие. Тщательно проверяй, есть ли в запросе такие указания.
При наличии в запросе относительного периода (например, "за последние 6 месяцев"), вычисляй "start_date" как дату начала этого периода, а "end_date" устанавливай как текущую дату ({current_date}).
ВАЖНО! Не забывай про "end_date"! Например, если спрашивают "за последний месяц", то "end_date" должна быть равна текущей дате ({current_date}).
Для справки:
Кварталы: 01.01 — 31.03; 01.04 — 30.06; 01.07 — 30.09; 01.10 — 31.12
Полугодия: 01.01 — 30.06; 01.07 — 31.12
Убедись, что для запросов с относительными датами корректно вычисляются и "start_date", и "end_date", и что "end_date" не остается пустым или равным null.

Если поля не указаны - необходимо их вернуть со значением null.

Проверь всё несколько раз.

ПРИМЕРЫ

Вопрос: "Привет! Что ты можешь подсказать для пловца в ближайшее время?"
Ответ: {{"start_date":"{current_date}","end_date": null,"name":null,"sport":"плавание","sex":true,"min_age": null,"max_age": null}}
Вопрос: "Какие есть чемпионаты по скалолазанию для девочки 10 лет на ближайший месяц?"
Ответ: {{"start_date":"{current_date}","end_date":"{current_date_plus_month}","name":"ЧЕМПИОНАТ","sport":"скалолазание","sex":false,"min_age":10,"max_age":10}}
Вопрос: "Соревнование для борца от 18 лет"
Ответ: {{"start_date":"{current_date}","end_date":null,"name":"СОРЕВНОВАНИЕ","sport":"борьба","sex":true,"min_age":18,"max_age":null}}
Вопрос: "Чемпионаты за последний месяц"
Ответ: {{"start_date":"{current_date_minus_month}","end_date":"{current_date}","name":"ЧЕМПИОНАТ","sport":"null","sex":null,"min_age":null,"max_age":null}}
Вопрос: "Учебно-тренеровочные мероприятия за этот год по айкидо"
Ответ: {{"start_date":"{current_start_year}","end_date":"{current_end_year}","name":"ТРЕНЕРОВОЧНЫЕ","sport":"айкидо","sex":null,"min_age":null,"max_age":null}}
"""

prompt_template = ChatPromptTemplate.from_messages(
    [("system", SYSTEM_TEMPLATE), ("user", "{input}")]
)

parser = JsonOutputParser()

def current_date():
    return datetime.now().strftime('%d.%m.%Y')

def current_start_year():
    return f"01.01.{datetime.now().year}"

def current_end_year():
    return f"31.12.{datetime.now().year}"

def current_date_plus_month():
    today = datetime.now()
    next_month = today + relativedelta(months=+1)
    return next_month.strftime('%d.%m.%Y')

def current_date_minus_month():
    today = datetime.now()
    prev_month = today + relativedelta(months=+1)
    return prev_month.strftime('%d.%m.%Y')

class InputType(BaseModel):
    input: str

question_chain = RunnableParallel(
    current_date=lambda x: current_date(),
    current_date_plus_month=lambda x: current_date_plus_month(),
    current_date_minus_month=lambda x: current_date_minus_month(),
    current_start_year=lambda x: current_start_year(),
    current_end_year=lambda x: current_end_year(),
    input=RunnablePassthrough(),
).with_types(
    input_type=InputType
) | prompt_template | chat | parser
