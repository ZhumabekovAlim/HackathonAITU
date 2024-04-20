import streamlit as st
import json
import requests

def load_data(filename):
    try:
        with open(filename, 'r', encoding='utf-8') as file:
            data = json.load(file)
        return data
    except UnicodeDecodeError:
        with open(filename, 'r', encoding='latin1') as file:
            data = json.load(file)
        return data
    except Exception as e:
        st.error(f"Error loading JSON file: {e}")
        return None 

def query_model(prompt):
    try:
        response = requests.post(
            'https://api.mistral.ai/v0.1/predict',
            headers={'Authorization': 'Bearer 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'},
            json={'model': 'mistralai/Mixtral-8x7B-Instruct-v0.1', 'prompt': prompt}
        )
        if response.status_code == 200:
            return response.json()
        else:
            st.error(f"Failed to fetch data from API: {response.status_code}, {response.text}")
            return None
    except requests.exceptions.RequestException as e:
        st.error(f"Request failed: {e}")
        return None

def main():
    st.title('Interact with Mistral AI Model')
    
    data = load_data(r'C:\Users\Alim\GolandProjects\awesomeProject\nlp\sight_details.json')
    if data:
        st.write('Loaded data:', data)
    else:
        st.write("Failed to load data.")
    
    user_input = st.text_area("Enter your prompt:", value="Type your prompt here...")
    
    if st.button('Send Prompt to Model'):
        model_response = query_model(user_input)
        if model_response:
            st.write('Model Response:', model_response)
        else:
            st.write("Failed to get a valid response from the model.")

if __name__ == "__main__":
    main()
