import uvicorn
from fastapi import FastAPI
from langserve import add_routes
from chain import question_chain

app = FastAPI()

add_routes(
    app,
    question_chain,
    path="/chat",
)

if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=True)
