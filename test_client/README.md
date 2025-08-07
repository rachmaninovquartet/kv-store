First install the requirements file #TODO save current requirements
`
pip install -r requirements.txt
`

The test client relies on an env variable for the server url:
`
SERVER_URL=http://localhost:8000 python client.py
`

otherwise start like:
`
python client.py
`
