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
#     'Token':'1|1701630097|27369556dc2fe3428a6372d41787589a',
#     'UserID':1,
# }

# r = requests.post(history_url, json=history_data)
# print(r.text)


# Favorites API

# record_favorites_url = host + '/record-favorites'
# record_favorites_data = {
#     'Token':'1|1701630097|27369556dc2fe3428a6372d41787589a',
#     'UserID':1,
#     'VideoID':2,
# }

# r = requests.post(record_favorites_url, json=record_favorites_data)
# print(r.text)

# favorites_url = host + '/favorites'
# favorites_data = {
#     'Token':'1|1701630097|27369556dc2fe3428a6372d41787589a',
#     'UserID':1,
# }

# r = requests.post(favorites_url, json=favorites_data)
# print(r.text)