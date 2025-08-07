from typing import Any, Dict, Optional
from storage.key_value_store import KeyValueStore


"""
In-memory store for the key-value store
"""
class InMemoryStore(KeyValueStore):
    def __init__(self):
        self._store: Dict[str, Any] = {}
    
    async def store(self, key: str, value: Any, ttl: Optional[int] = None) -> bool:
        # Note: In-memory implementation ignores TTL for simplicity
        # You could implement TTL with asyncio.create_task and asyncio.sleep
        self._store[key] = value
        return True
    
    async def retrieve(self, key: str) -> Optional[Any]:
        return self._store.get(key)
    
    async def delete(self, key: str) -> bool:
        if key in self._store:
            del self._store[key]
            return True
        return False
    
    async def exists(self, key: str) -> bool:
        return key in self._store