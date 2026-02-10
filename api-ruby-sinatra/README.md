# 💎 Todo API — Ruby / Sinatra + Sequel

A minimal CRUD API to learn REST fundamentals.

## Quick Start

```bash
bundle install
ruby app.rb
```

Server runs at **http://localhost:4567**

## Try It

```bash
# Create
curl -s -X POST http://localhost:4567/todos \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk"}' | ruby -rjson -e 'puts JSON.pretty_generate(JSON.parse(STDIN.read))'

# List all
curl -s http://localhost:4567/todos | ruby -rjson -e 'puts JSON.pretty_generate(JSON.parse(STDIN.read))'

# Get one
curl -s http://localhost:4567/todos/1 | ruby -rjson -e 'puts JSON.pretty_generate(JSON.parse(STDIN.read))'

# Update
curl -s -X PUT http://localhost:4567/todos/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy oat milk","done":true}' | ruby -rjson -e 'puts JSON.pretty_generate(JSON.parse(STDIN.read))'

# Delete
curl -X DELETE http://localhost:4567/todos/1
```

## What You'll Learn

| Concept | Where to look |
|---------|---------------|
| **Routing** | `get "/todos"`, `post "/todos"`, etc. |
| **Path params** | `params[:id]` from `:id` in the URL |
| **JSON parsing** | `json_body` helper (manual — no magic!) |
| **JSON response** | `.to_json` on hashes/arrays |
| **ORM** | `Todo < Sequel::Model` — maps class ↔ table |
| **Status codes** | `status 201`, `halt 404`, `status 204` |

## Comparison to Python

- Sinatra is more **explicit** — you manually parse JSON and set content types
- FastAPI auto-validates with Pydantic; Sinatra leaves validation to you
- Ruby blocks (`do...end`) replace Python decorators (`@app.get`)
