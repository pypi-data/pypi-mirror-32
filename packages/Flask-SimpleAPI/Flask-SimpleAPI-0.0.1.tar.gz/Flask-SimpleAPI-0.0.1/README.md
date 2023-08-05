# Simple way to create API

# user.py
def GET_index():
    return {'name': 'User'}
    

def GET_info(str_id):
    return {'user_id': str_id}


# app.py
from flask_simpleapi
import user

api = flask_simpleapi(__name__)

api.register_api(user)

api.run('0.0.0.0', host=5000, debug=True)


# This API will generate 2 endpoints
/user/index
/user/index/<str:str_id>
