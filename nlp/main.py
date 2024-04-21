from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import openai

app = FastAPI()
openai_api_key = "sk-proj-Qhx9pfRtzZNrqtg2YGNCT3BlbkFJ3yR3fMx5v9UwVmQeu7wr"

with open("sight_details.txt", "r", encoding="utf-8") as file:
    dataset_text = file.read()

context = dataset_text

class Query(BaseModel):
    question: str

@app.post("/query/")
async def answer_query(query: Query):
    return {"answer": "This is a test response"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
