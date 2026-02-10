# 🐍 Todo API — Python / FastAPI + SQLAlchemy

A minimal CRUD API to learn REST fundamentals.

## Quick Start

```bash
python -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
uvicorn main:app --reload
```

Open **http://localhost:8000/docs** — FastAPI generates interactive Swagger docs for free!

## Try It

```bash
# Create
curl -s -X POST http://localhost:8000/todos \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk"}' | python -m json.tool

# List all
curl -s http://localhost:8000/todos | python -m json.tool

# Get one
curl -s http://localhost:8000/todos/1 | python -m json.tool

# Update
curl -s -X PUT http://localhost:8000/todos/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy oat milk","done":true}' | python -m json.tool

# Delete
curl -X DELETE http://localhost:8000/todos/1
```

## What You'll Learn

| Concept | Where to look |
|---------|---------------|
| **Routing** | `@app.get`, `@app.post`, etc. decorators |
| **Path params** | `{todo_id}` in the URL |
| **Request validation** | `TodoCreate` / `TodoUpdate` Pydantic models |
| **Response serialisation** | `response_model=TodoResponse` |
| **Dependency injection** | `Depends(get_db)` |
| **ORM** | `Todo` class maps to a DB table |
| **HTTP status codes** | `201` for created, `204` for deleted, `404` for not found |
