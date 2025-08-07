First install the requirements file #TODO save current requirements
`
pip install -r requirements.txt
`

The KV service offers 2 storage options, in memory and redis. Default is in memory, to use redis, start like this:
`
STORAGE_TYPE=redis python server.py
`

otherwise start like:
`
python server.py
`
