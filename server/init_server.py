import os
import redis
from fastapi import Depends
from services.key_value_service import KeyValueService
from storage.key_value_store import KeyValueStore
from storage.redis_store import RedisStore
from storage.in_memory_store import InMemoryStore


"""
Get the storage type for startup (memory or redis)
"""
def get_storage_type() -> str:
    return os.getenv("STORAGE_TYPE", "memory")  


"""
Create the store based on the storage type
"""
def create_store() -> KeyValueStore:
    storage_type = get_storage_type()
    
    if storage_type == "redis":
        redis_client = redis.Redis(
            host=os.getenv("REDIS_HOST", "localhost"),
            port=int(os.getenv("REDIS_PORT", 6379)),
            db=int(os.getenv("REDIS_DB", 0)),
            decode_responses=True
        )
        return RedisStore(redis_client)
    else:
        return InMemoryStore()


"""
Global store instance
"""
store_instance = None


"""
Get the store
"""
def get_store() -> KeyValueStore:
    global store_instance
    if store_instance is None:
        store_instance = create_store()
    return store_instance


"""
Get the service
"""
def get_service(store: KeyValueStore = Depends(get_store)) -> KeyValueService:
    return KeyValueService(store)