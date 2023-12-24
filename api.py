import requests

host = 'http://127.0.0.1:8080'

# LOGIN API

# login_url = host + '/login'
# login_data = {
#     'email':'a@test.com',
#     'password':'e10adc3949ba59abbe56e057f20f883e',
# }

# r = requests.post(login_url, json=login_data)
# print(r.text)

# History API

# record_history_url = host + '/record-history'
# record_history_data = {
#     'Token':'1|1701630097|27369556dc2fe3428a6372d41787589a',
#     'UserID':1,
#     'VideoID':2,
#     'Episode':3,
# }

# r = requests.post(record_history_url, json=record_history_data)
# print(r.text)

# history_url = host + '/history'
# history_data = {
#     'Token':'1|1702855882|d5392ff4ac3bc2d45ba950b9e15c0f3c',
#     'UserID':1,
# }

# r = requests.post(history_url, json=history_data)
# print(r.text)


# Favorites API

# record_favorites_url = host + '/record-favorites'
# record_favorites_data = {
#     'Token':'1|1702855882|d5392ff4ac3bc2d45ba950b9e15c0f3c',
#     'UserID':1,
#     'VideoID':5,
# }

# r = requests.post(record_favorites_url, json=record_favorites_data)
# print(r.text)

# remove_favorites_url = host + '/remove-favorites'
# remove_favorites_data = {
#     'Token':'1|1702855882|d5392ff4ac3bc2d45ba950b9e15c0f3c',
#     'UserID':1,
#     'VideoID':5,
# }

# r = requests.post(remove_favorites_url, json=remove_favorites_data)
# print(r.text)

# favorites_url = host + '/favorites'
# favorites_data = {
#     'Token':'1|1702855882|d5392ff4ac3bc2d45ba950b9e15c0f3c',
#     'UserID':1,
# }

# r = requests.post(favorites_url, json=favorites_data)
# print(r.text)

# Points API

# points_url = host + '/points'
# points_data = {
#     'Token':'1|1701910411|894979dbbebc3eeb86c7b21e799fd40d',
#     'UserID':1,
# }

# r = requests.post(points_url, json=points_data)
# print(r.text)

# checkin_url = host + '/checkin'
# checkin_data = {
#     'Token':'1|1701910411|894979dbbebc3eeb86c7b21e799fd40d',
#     'UserID':1,
# }

# r = requests.post(checkin_url, json=checkin_data)
# print(r.text)

# Episodes API

# episodes_url = host + '/episodes'
# episodes_data = {
#     'Token':'1|1703389604|a5e6df9347234579865b33895f089645',
#     'UserID':1,
# }

# r = requests.post(episodes_url, json=episodes_data)
# print(r.text)

# unlock_episode_url = host + '/unlock-episode'
# unlock_episode_data = {
#     'Token':'1|1703389604|a5e6df9347234579865b33895f089645',
#     'UserID':1,
#     'VideoID':1,
#     'Episode':8,
# }

# r = requests.post(unlock_episode_url, json=unlock_episode_data)
# print(r.text)

# Admin API

# upload_url = host + '/upload-video-info'
# upload_data = {
#     'Admin':'testtesttest',
#     'Name':'CEO 2',
#     'Type':2,
#     'TotalNumber':2,
#     'BaseUrl':'https://dc4ef1i295q51.cloudfront.net/test5',
# }

# r = requests.post(upload_url, json=upload_data)
# print(r.text)

# video_url = host + '/video-list'

# r = requests.get(video_url)
# print(r.text)

# stripe_url = host + '/create-stripe-payment'
# stripe_data = {
# 	'ID':1,
# 	'Amount':100,
# 	'ProductID':'price_1OQqxYLvs8YNyX8sRMRaBbcN',
# 	'SuccessURL':'http://localhost:3000/',
# 	'CancelURL':'http://localhost:3000/',
# }

# r = requests.post(stripe_url, json=stripe_data)
# print(r.text)