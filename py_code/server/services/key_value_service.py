from storage.key_value_store import KeyValueStore
from typing import Any, List, Optional


"""
Service layer for the key-value store
"""
class KeyValueService:
    def __init__(self, store: KeyValueStore):
        self.store = store
    
    async def set_key_value(self, key: str, value: Any, ttl: Optional[int] = None) -> bool:
        return await self.store.store(key, value, ttl)
    
    async def get_value(self, key: str) -> Optional[Any]:
        return await self.store.retrieve(key)
    
    async def delete_key(self, key: str) -> bool:
        return await self.store.delete(key)
    
    async def get_all_keys(self) -> List[str]:
        return await self.store.list_keys()
    
    async def key_exists(self, key: str) -> bool:
        return await self.store.exists(key)