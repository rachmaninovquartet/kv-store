import redis
import json
from typing import Any, List, Optional
from storage.key_value_store import KeyValueStore


"""
Redis store for the key-value store
"""
class RedisStore(KeyValueStore):
    def __init__(self, redis_client: redis.Redis):
        self.redis_client = redis_client
    
    def _serialize_value(self, value: Any) -> str:
        return json.dumps(value)
    
    def _deserialize_value(self, value: str) -> Any:
        try:
            return json.loads(value)
        except json.JSONDecodeError:
            return value
    
    async def store(self, key: str, value: Any, ttl: Optional[int] = None) -> bool:
        try:
            serialized_value = self._serialize_value(value)
            if ttl:
                self.redis_client.setex(key, ttl, serialized_value)
            else:
                self.redis_client.set(key, serialized_value)
            return True
        except redis.RedisError:
            return False
    
    async def retrieve(self, key: str) -> Optional[Any]:
        try:
            value = self.redis_client.get(key)
            if value is None:
                return None
            return self._deserialize_value(value)
        except redis.RedisError:
            return None
    
    async def delete(self, key: str) -> bool:
        try:
            result = self.redis_client.delete(key)
            return result > 0
        except redis.RedisError:
            return False
    
    async def exists(self, key: str) -> bool:
        try:
            return bool(self.redis_client.exists(key))
        except redis.RedisError:
            return False