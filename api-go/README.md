# 🐹 Todo API — Go / net/http + database/sql

A minimal CRUD API to learn REST fundamentals — **zero frameworks**.

## Quick Start

```bash
go mod tidy
go run .
```

Server runs at **http://localhost:8080**

## Try It

```bash
# Create
curl -s -X POST http://localhost:8080/todos \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk"}' | jq

# List all
curl -s http://localhost:8080/todos | jq

# Get one
curl -s http://localhost:8080/todos/1 | jq

# Update
curl -s -X PUT http://localhost:8080/todos/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy oat milk","done":true}' | jq

# Delete
curl -X DELETE http://localhost:8080/todos/1
```

## What You'll Learn

| Concept | Where to look |
|---------|---------------|
| **Routing** | `http.HandleFunc` + manual `r.Method` switch |
| **Path params** | `strings.TrimPrefix` to extract ID (no router!) |
| **JSON decode** | `json.NewDecoder(r.Body).Decode(&input)` |
| **JSON encode** | `json.NewEncoder(w).Encode(data)` |
| **SQL queries** | `db.Query`, `db.QueryRow`, `db.Exec` |
| **Error handling** | Explicit `if err != nil` everywhere |
| **Struct tags** | `` `json:"title"` `` controls JSON field names |
| **Pointer trick** | `*string` distinguishes "not sent" from "empty" |

## Comparison to Python & Ruby

- Go has **no ORM** here — you write SQL directly and `Scan` into structs
- Go has **no router** — you dispatch on `r.Method` manually
- Go has **no auto-validation** — you check everything yourself
- More code, but every line is **explicit and predictable**
- This is why Go is popular for production APIs — no hidden magic
