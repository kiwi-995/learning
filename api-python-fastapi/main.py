"""
Minimal Todo API — FastAPI + SQLAlchemy
========================================
Run:   uvicorn main:app --reload
Docs:  http://localhost:8000/docs  (auto-generated Swagger UI!)

Key concepts demonstrated:
  • Routing           — @app.get / .post / .put / .delete
  • Path parameters   — /todos/{todo_id}
  • Request body      — Pydantic models auto-parse JSON
  • Response model    — FastAPI serialises to JSON for you
  • Dependency inject. — Depends(get_db) provides a DB session
  • ORM               — SQLAlchemy maps Python classes ↔ DB tables
"""

from fastapi import FastAPI, Depends, HTTPException
from pydantic import BaseModel
from sqlalchemy import create_engine, Column, Integer, String, Boolean
from sqlalchemy.orm import declarative_base, sessionmaker, Session

# ─── Database setup ───────────────────────────────────────────────
# SQLite stores everything in a single file — no server needed.
DATABASE_URL = "sqlite:///./todos.db"

engine = create_engine(
    DATABASE_URL,
    connect_args={"check_same_thread": False},  # needed for SQLite + threads
)
SessionLocal = sessionmaker(bind=engine)
Base = declarative_base()


# ─── ORM Model ────────────────────────────────────────────────────
# This class IS the database table. SQLAlchemy creates it for us.
class Todo(Base):
    __tablename__ = "todos"

    id = Column(Integer, primary_key=True, index=True)
    title = Column(String, nullable=False)
    done = Column(Boolean, default=False)


# Create the table if it doesn't exist yet
Base.metadata.create_all(bind=engine)


# ─── Pydantic Schemas ─────────────────────────────────────────────
# Schemas define WHAT the API accepts and returns (separate from DB).
# FastAPI uses these to:
#   1. Validate incoming JSON automatically
#   2. Generate the OpenAPI docs
#   3. Serialise responses to JSON

class TodoCreate(BaseModel):
    """What the client sends when creating a todo."""
    title: str
    done: bool = False


class TodoUpdate(BaseModel):
    """What the client sends when updating a todo."""
    title: str | None = None
    done: bool | None = None


class TodoResponse(BaseModel):
    """What the API returns."""
    id: int
    title: str
    done: bool

    model_config = {"from_attributes": True}  # allows ORM → Pydantic conversion


# ─── Dependency: DB session ───────────────────────────────────────
# FastAPI calls this function for every request that needs a DB.
# `yield` ensures the session is closed after the request finishes.
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


# ─── App & Routes ─────────────────────────────────────────────────
app = FastAPI(title="Todo API", version="0.1.0")


@app.get("/todos", response_model=list[TodoResponse])
def list_todos(db: Session = Depends(get_db)):
    """Return all todos."""
    return db.query(Todo).all()


@app.get("/todos/{todo_id}", response_model=TodoResponse)
def get_todo(todo_id: int, db: Session = Depends(get_db)):
    """Return a single todo by ID."""
    todo = db.query(Todo).get(todo_id)
    if not todo:
        raise HTTPException(status_code=404, detail="Todo not found")
    return todo


@app.post("/todos", response_model=TodoResponse, status_code=201)
def create_todo(data: TodoCreate, db: Session = Depends(get_db)):
    """Create a new todo."""
    todo = Todo(title=data.title, done=data.done)
    db.add(todo)
    db.commit()
    db.refresh(todo)  # reload to get the auto-generated ID
    return todo


@app.put("/todos/{todo_id}", response_model=TodoResponse)
def update_todo(todo_id: int, data: TodoUpdate, db: Session = Depends(get_db)):
    """Update an existing todo (partial update)."""
    todo = db.query(Todo).get(todo_id)
    if not todo:
        raise HTTPException(status_code=404, detail="Todo not found")

    if data.title is not None:
        todo.title = data.title
    if data.done is not None:
        todo.done = data.done

    db.commit()
    db.refresh(todo)
    return todo


@app.delete("/todos/{todo_id}", status_code=204)
def delete_todo(todo_id: int, db: Session = Depends(get_db)):
    """Delete a todo."""
    todo = db.query(Todo).get(todo_id)
    if not todo:
        raise HTTPException(status_code=404, detail="Todo not found")

    db.delete(todo)
    db.commit()
