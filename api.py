import requests

host = 'http://127.0.0.1:8080'

# LOGIN API

# login_url = host + '/login'
# login_data = {
#     'email':'a@test.com',
#     'password':'123456',
# }

# r = requests.post(login_url, json=login_data)
# print(r.text)

# REGISTER API

# register_url = host + '/register'
# register_data = {
#     'email':'a@test.com',
#     'password':'123456',
# }

# r = requests.post(register_url, json=register_data)
# print(r.text)

# VIDEO API

# login_url = host + '/login'
# login_data = {
#     'email':'a@test.com',
#     'password':'e10adc3949ba59abbe56e057f20f883e',
# }

# print('Login to get token')
# r = requests.post(login_url, json=login_data)
# print(r.text)

# user_video_url = host + '/user-video'
# user_video_data = {
#     'Token':'1|1701501957|60025be94df7d7bd75331513e184e9d1',
#     'UserID':1,
#     'VideoID':2,
#     'Code':1,
# }

# print('Record the video that user liked')
# r = requests.post(user_video_url, json=user_video_data)
# print(r.text)

# user_video_list_url = host + '/user-video-list'
# user_video_list_data = {
#     'Token':'1|1701501957|60025be94df7d7bd75331513e184e9d1',
#     'UserID':1,
#     'Code':2,
# }

# print('Get the video that user liked')
# r = requests.post(user_video_list_url, json=user_video_list_data)
# print(r.text)