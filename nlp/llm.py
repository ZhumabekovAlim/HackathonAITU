from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import openai

app = FastAPI()
openai_api_key = "sk-proj-DooqRHAy1VuviqipRvg3T3BlbkFJYjgvYzv87LKheKm1OUow"

class Query(BaseModel):
    question: str

@app.post("/query/")
async def answer_query(query: Query):
    try:
        response = openai.Completion.create(
            engine="gpt-4-turbo",  # Укажите актуальную модель ChatGPT
            prompt=query.question,
            max_tokens=150
        )
        answer_text = response.choices[0].text.strip()
        return {"answer": answer_text}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
