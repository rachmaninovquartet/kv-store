from pydantic import BaseModel
from typing import List, Any, Optional


"""
Request model for a key-value pair
"""
class KeyValue(BaseModel):
    key: str
    value: Any
    ttl: Optional[int] = None


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