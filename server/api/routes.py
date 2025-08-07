from fastapi import APIRouter, Depends, HTTPException
from services.key_value_service import KeyValueService
from models.responses import MessageResponse, ValueResponse, KeyValue
from init_server import get_service, get_storage_type, get_store
from storage.redis_store import RedisStore
from storage.in_memory_store import InMemoryStore

# Create a router instead of a separate app
router = APIRouter()


"""
Set a key-value pair
"""
@router.post("/set", response_model=MessageResponse)
async def store_key_value(
    item: KeyValue,
    service: KeyValueService = Depends(get_service)
):
    success = await service.set_key_value(item.key, item.value, item.ttl)
    if not success:
        raise HTTPException(status_code=500, detail="Failed to store key-value pair")
    return MessageResponse(message=f"Key '{item.key}' stored successfully")


"""
Get a value by key
"""
@router.get("/get/{key}", response_model=ValueResponse)
async def retrieve_value(
    key: str,
    service: KeyValueService = Depends(get_service)
):
    value = await service.get_value(key)
    if value is None:
        raise HTTPException(status_code=404, detail=f"Key '{key}' not found")
    return ValueResponse(value=value)


"""
Delete a key-value pair
"""
@router.delete("/delete/{key}", response_model=MessageResponse)
async def delete_key(
    key: str,
    service: KeyValueService = Depends(get_service)
):
    """Delete a key-value pair"""
    success = await service.delete_key(key)
    if not success:
        raise HTTPException(status_code=404, detail=f"Key '{key}' not found")
    return MessageResponse(message=f"Key '{key}' deleted successfully")


"""
Check if a key exists
"""
@router.get("/exists/{key}")
async def check_key_exists(
    key: str,
    service: KeyValueService = Depends(get_service)
):
    """Check if a key exists"""
    exists = await service.key_exists(key)
    return {"key": key, "exists": exists}


"""
Root endpoint
"""
@router.get("/")
async def root():
    return {"message": "KV Service is running"}
