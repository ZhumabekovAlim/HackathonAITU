from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
import openai
import uvicorn

app = FastAPI()


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  
    allow_credentials=True,
    allow_methods=["GET", "POST", "PUT", "DELETE"],  
    allow_headers=["*"],  
)

openai_api_key = "sk-proj-9kLlLOwSPSj6qSnfni0ET3BlbkFJWUTrdYb2XxH7mNOBxd2s"

with open("sight_details.txt", "r", encoding="utf-8") as file:
    dataset_text = file.read()

context = dataset_text

class Query(BaseModel):
    question: str

@app.post("/query/")
async def answer_query(query: Query):
    try:
        response = openai.ChatCompletion.create(
            model="gpt-4-turbo",  
            messages=[{"role": "system", "content": context+"(дай ответ максимум в 110 слов)"}, {"role": "user", "content": query.question}],
            max_tokens=500,
            api_key=openai_api_key
        )
        answer_text = response.choices[0].message['content'].strip()
        return {"answer": answer_text}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
