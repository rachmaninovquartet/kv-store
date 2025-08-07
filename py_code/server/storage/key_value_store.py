from abc import ABC, abstractmethod
from typing import Any, List, Optional


"""
Abstract base store for the key-value store
"""
class KeyValueStore(ABC):
    @abstractmethod
    async def store(self, key: str, value: Any, ttl: Optional[int] = None) -> bool:
        pass
    
    @abstractmethod
    async def retrieve(self, key: str) -> Optional[Any]:
        pass
    
    @abstractmethod
    async def delete(self, key: str) -> bool:
        pass
    
    @abstractmethod
    async def exists(self, key: str) -> bool:
        pass