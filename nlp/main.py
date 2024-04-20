from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import openai

app = FastAPI()
openai_api_key = "sk-proj-DooqRHAy1VuviqipRvg3T3BlbkFJYjgvYzv87LKheKm1OUow"

# Load your dataset from a plain text file
with open("sight_details.txt", "r", encoding="utf-8") as file:
    dataset_text = file.read()

# Create a string that combines all data from your dataset
context = dataset_text  # Here we assume the text file is already formatted correctly

class Query(BaseModel):
    question: str

@app.post("/query/")
async def answer_query(query: Query):
    try:
        # Send the question and context to the GPT model
        response = openai.ChatCompletion.create(
            model="gpt-4-turbo",  # Specify the appropriate ChatGPT model
            messages=[{"role": "system", "content": context+"(дай ответ максимум в 110 слов)"}, {"role": "user", "content": query.question}],
            max_tokens=500,
            api_key=openai_api_key
        )
        answer_text = response.choices[0].message['content'].strip()
        return {"answer": answer_text}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
