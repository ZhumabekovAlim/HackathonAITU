from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
import openai
import uvicorn

app = FastAPI()


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # You can specify specific origins instead of "*" for security
    allow_credentials=True,
    allow_methods=["GET", "POST", "PUT", "DELETE"],  # Specify the HTTP methods you want to allow
    allow_headers=["*"],  # You can specify specific headers or use "*" to allow all
)

openai_api_key = "sk-proj-3QMinI9N0ofzcR8IQYpVT3BlbkFJcB9z9Rxqvr0SLBZHMbjY"

with open("sight_details.txt", "r", encoding="utf-8") as file:
    dataset_text = file.read()

context = dataset_text

class Query(BaseModel):
    question: str

@app.post("/query")
async def answer_query(query: Query):
    try:
        response = openai.ChatCompletion.create(
            model="gpt-4-turbo",  # Specify the appropriate ChatGPT model
            messages=[{"role": "system", "content": context+"(дай ответ максимум в 50 слов)"}, {"role": "user", "content": query.question}],
            max_tokens=500,
            api_key=openai_api_key
        )
        answer_text = response.choices[0].message['content'].strip()
        return {"answer": answer_text}
    except Exception as e:
           print("An error occurred:", e)  # логирование ошибки
    raise HTTPException(status_code=500, detail="Internal server error")

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
