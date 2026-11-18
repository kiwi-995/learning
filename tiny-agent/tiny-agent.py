import anthropic
from rich.console import Console
from rich.markdown import Markdown

client = anthropic.Anthropic()
console = Console()

MODEL = "claude-haiku-5-5"
MAX_TOKENS = 16000
MESSAGES = []

def main_loop():
    loop = True

    while loop:
        loop = get_input()

        if loop:
            return_output()


def get_input():
    prompt = input("> ")
    
    if prompt == "exit":
        return False
    
    MESSAGES.append({
        "role": "user",
        "content": prompt
    })

    return True


def return_output():
    # send prompt to server and receive answer
    answer = call_inference_endpoint()
    
    console.print(Markdown("# DEBUG"))
    console.print(answer.content)

    for block in answer.content:
        if block.type == "thinking":
            console.print(Markdown("# THINKING"))
            console.print(Markdown(block.thinking))
    
        elif block.type == "text":
            console.print(Markdown("# ANSWER"))
            console.print(Markdown(block.text))

    MESSAGES.append({
        "role": "assistant",
        "content": answer.content
    })


def call_inference_endpoint():
    return client.messages.create(
        model=MODEL,
        max_tokens=MAX_TOKENS,
        messages=MESSAGES,
        thinking={"type": "adaptive", "display": "summarized"}
    )

if __name__=="__main__":
    main_loop()
