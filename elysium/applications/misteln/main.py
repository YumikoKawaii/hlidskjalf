import uvicorn
from app import initialize

application = initialize()

if __name__ == '__main__':
    uvicorn.run("main:application", host="0.0.0.0", port=8000, reload=True, access_log=False)
