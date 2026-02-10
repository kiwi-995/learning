from sqlalchemy.orm import Session
from . import models, schemas
from typing import List, Optional

def get_todos(db: Session) -> List[models.Todo]:
    return db.query(models.Todo).all()

def get_todo(db: Session, todo_id: int) -> Optional[models.Todo]:
    return db.query(models.Todo).filter(models.Todo.id == todo_id).first()

def create_todo(db: Session, todo: schemas.TodoCreate) -> models.Todo:
    db_todo = models.Todo(title=todo.title, done=todo.done or False)
    db.add(db_todo)
    db.commit()
    db.refresh(db_todo)
    return db_todo

def update_todo(db: Session, db_todo: models.Todo, todo: schemas.TodoUpdate) -> models.Todo:
    if todo.title is not None:
        db_todo.title = todo.title
    if todo.done is not None:
        db_todo.done = todo.done
    db.commit()
    db.refresh(db_todo)
    return db_todo

def delete_todo(db: Session, db_todo: models.Todo):
    db.delete(db_todo)
    db.commit()
