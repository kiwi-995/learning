import anthropic
from rich.console import Console
from rich.markdown import Markdown

client = anthropic.Anthropic()
console = Console()

MODEL = "claude-haiku-5-5"
MAX_TOKENS = 16000
MESSAGES = []

def main_loop():
    while True:
        prompt = input("> ")
        
        if prompt == "exit":
            return
        
        MESSAGES.append({
            "role": "user",
            "content": prompt
        })

        # send prompt to server and receive answer
        answer = call_inference_endpoint()
        
        answer_reasoning = ""
        answer_text = "".join([t.text for t in answer.content if t.type == "text"])
        
        console.print(Markdown(answer_text))
        
        print("DEBUG:")
        console.print(answer)

        MESSAGES.append({
            "role": "assistant",
            "content": answer.content
        })


def call_inference_endpoint():
    return client.messages.create(model=MODEL, max_tokens=MAX_TOKENS, messages=MESSAGES)

if __name__=="__main__":
    main_loop()
