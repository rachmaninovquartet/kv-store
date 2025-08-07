from fastapi import APIRouter, HTTPException, Query, Depends
import httpx
import asyncio
import os

router = APIRouter()

# Default server URL - will be overridden by dependency
DEFAULT_SERVER_URL = "http://localhost:8000"


"""
Get the server URL from the environment or use default
"""
def get_server_url() -> str:
    """Get server URL from environment or use default"""
    return os.getenv("SERVER_URL", DEFAULT_SERVER_URL)


"""
Test deletion of a key from the server
"""
@router.get("/test_deletion")
async def test_deletion(server_url: str = Depends(get_server_url)):
    """Test deletion by making HTTP request to main server's delete endpoint"""
    async with httpx.AsyncClient() as client:
        try:
            # First set a key to delete
            set_response = await client.post(
                f"{server_url}/set",
                json={"key": "test_delete_key", "value": "test_value"}
            )
            
            if set_response.status_code != 200:
                raise HTTPException(status_code=500, detail="Failed to set key for deletion test")
            
            # check if the key exists
            exists_response = await client.get(f"{server_url}/exists/test_delete_key")
            if exists_response.status_code != 200:
                raise HTTPException(status_code=500, detail="Failed to check if key exists")
            if exists_response.json()["exists"] is False:
                raise HTTPException(status_code=500, detail="Key does not exist")
            
            # Now delete the key
            delete_response = await client.delete(f"{server_url}/delete/test_delete_key")
            
            if delete_response.status_code == 200:
                return {"message": "Test deletion successful", "details": delete_response.json()}
            else:
                return {"message": "Test deletion failed", "status": delete_response.status_code}
                
        except httpx.RequestError as e:
            raise HTTPException(status_code=500, detail=f"Failed to connect to main server: {str(e)}")


"""
Test overwriting a value in the server
"""
@router.get("/test_overwrite")
async def test_overwrite(server_url: str = Depends(get_server_url)):
    """Test overwrite by making HTTP request to main server's set endpoint"""
    async with httpx.AsyncClient() as client:
        try:
            # Set a key with initial value
            set_response1 = await client.post(
                f"{server_url}/set",
                json={"key": "test_overwrite_key", "value": "initial_value"}
            )
            
            if set_response1.status_code != 200:
                raise HTTPException(status_code=500, detail="Failed to set initial key")
            
            # check if the key exists
            exists_response = await client.get(f"{server_url}/exists/test_overwrite_key")
            if exists_response.status_code != 200:
                raise HTTPException(status_code=500, detail="Failed to check if key exists")
            if exists_response.json()["exists"] is False:
                raise HTTPException(status_code=500, detail="Key does not exist")
            
            # Overwrite the same key with new value
            set_response2 = await client.post(
                f"{server_url}/set",
                json={"key": "test_overwrite_key", "value": "overwritten_value"}
            )
            
            if set_response2.status_code == 200:
                # Get the value to confirm overwrite
                get_response = await client.get(f"{server_url}/get/test_overwrite_key")
                
                if get_response.status_code == 200:
                    return {
                        "message": "Test overwrite successful", 
                        "final_value": get_response.json()["value"]
                    }
                else:
                    return {"message": "Test overwrite failed - couldn't retrieve value"}
            else:
                return {"message": "Test overwrite failed", "status": set_response2.status_code}
                
        except httpx.RequestError as e:
            raise HTTPException(status_code=500, detail=f"Failed to connect to main server: {str(e)}")


"""
Test getting a value from the server
"""
@router.get("/test_get/{key}")
async def test_get(key: str, server_url: str = Depends(get_server_url)):
    """Test getting a value from the main server"""
    async with httpx.AsyncClient() as client:
        try:
            response = await client.get(f"{server_url}/get/{key}")
            
            if response.status_code == 200:
                return {"message": "Test get successful", "value": response.json()["value"]}
            elif response.status_code == 404:
                return {"message": "Key not found", "key": key}
            else:
                return {"message": "Test get failed", "status": response.status_code}
                
        except httpx.RequestError as e:
            raise HTTPException(status_code=500, detail=f"Failed to connect to main server: {str(e)}")

