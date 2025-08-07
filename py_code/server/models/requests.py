from pydantic import BaseModel
from typing import Any, Optional

"""
Request model for a key-value pair
"""
class KeyValue(BaseModel):
    key: str
    value: Any
    ttl: Optional[int] = None
