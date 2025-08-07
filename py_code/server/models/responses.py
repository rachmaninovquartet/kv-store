from pydantic import BaseModel
from typing import Any


"""
Response model for a value
"""
class ValueResponse(BaseModel):
    value: Any


"""
Response model for a message
"""
class MessageResponse(BaseModel):
    message: str