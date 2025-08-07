import uvicorn
import redis
from fastapi import FastAPI
from contextlib import asynccontextmanager
from init_app import get_storage_type, get_store
from server.storage.redis_store import RedisStore
# Import routes to include them in the main app
from server.routes import router


"""
Lifespan for the app
"""
@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup
    storage_type = get_storage_type()
    print(f"Starting with {storage_type} storage")
    
    # Test connection for Redis
    if storage_type == "redis":
        try:
            store = get_store()
            if isinstance(store, RedisStore):
                store.redis_client.ping()
                print("Redis connection successful")
        except redis.RedisError as e:
            print(f"Redis connection failed: {e}")
            raise
    
    yield
    
    print("Shutting down...")


app = FastAPI(title="Key-Value Store API", lifespan=lifespan)

# Include routes from routes.py
app.include_router(router)

if __name__ == "__main__":
    uvicorn.run("app:app", host="0.0.0.0", port=8000, reload=True) 