from pydantic import BaseModel
from typing import Optional

class TodoBase(BaseModel):
    title: Optional[str] = None
    done: Optional[bool] = None

class TodoCreate(TodoBase):
    title: str

class TodoUpdate(TodoBase):
    pass

class TodoResponse(TodoBase):
    id: int
    class Config:
        orm_mode = True
