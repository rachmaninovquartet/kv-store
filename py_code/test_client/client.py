import os
import uvicorn
from fastapi import FastAPI
from contextlib import asynccontextmanager
from api.routes import router


"""
Lifespan for the app
"""
@asynccontextmanager
async def lifespan(app: FastAPI):
    print("Starting Test Client...")

    yield
    
    print("Shutting down Test Client...")


app = FastAPI(title="Test Client", lifespan=lifespan)

app.include_router(router)

if __name__ == "__main__":
    uvicorn.run("client:app", host="0.0.0.0", port=8002, reload=True) 