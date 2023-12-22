
# Perseus

## What style is this project?
Light and Clear

## How to start the server?

start the server:
```
go build
./Perseus
```
compile the program:
```
go build
```
clean the compile file:
```
go clean
```
compile the program for Linux:
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```
start the server for Linux:
```
sudo lsof -i :8080
kill [PID]
git
nohup ./perseus/Perseus > perseus.log 2>&1 &
```

## What API do we have?

### Login API
`/login`  
Authenticate users and set token.

### Register API
`/register`  
User register and set token.  

### Reset API
`/reset-email`  
Send reset email.  
`/reset`  
Reset User password.  

### Verify API
`/verify-email`  
Send verify email.  
`/verify`  
Veryify user email.  

### Video API
`/video-list`  
Get video list.  

### Favorites API
`/remove-favorites`  
Remove user's favorites.  
`/record-favorites`  
Record user's favorites.  
`/favorites`  
Get user's favorites list.  

### History API
`/record-history`  
Record user's history.  
`/history`  
Get user's history list.  

### Points API
`/points`  
Get user's points.  
`/already-checkin`  
Check if already checked in today.  
`/checkin`  
Check in and add points.  

### Admin API
`/upload-video-info`  
Upload video info for admin.  

## What tables are in the database?

`user_info`
```
CREATE TABLE user_info (
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
password VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
activated BOOLEAN,
vip BOOLEAN,
points INT UNSIGNED NOT NULL
);
```

`video_info`
```
CREATE TABLE video_info (
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(255) NOT NULL,
type INT UNSIGNED NOT NULL,
total_number INT UNSIGNED NOT NULL,
base_url VARCHAR(255) NOT NULL,
hot BOOLEAN NOT NULL DEFAULT FALSE
);

```

`user_history`
```
CREATE TABLE user_history (
user_id INT UNSIGNED NOT NULL,
video_id INT UNSIGNED NOT NULL,
episode INT UNSIGNED NOT NULL,
watch_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES user_info(id),
FOREIGN KEY (video_id) REFERENCES video_info(id)
);
```

`user_like`
```
CREATE TABLE user_like (
user_id INT UNSIGNED NOT NULL,
video_id INT UNSIGNED NOT NULL,
watch_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES user_info(id),
FOREIGN KEY (video_id) REFERENCES video_info(id)
);
```

`activities`
```
CREATE TABLE activities (
user_id INT UNSIGNED NOT NULL,
type INT UNSIGNED NOT NULL,
points INT UNSIGNED NOT NULL,
time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES user_info(id)
);
```

## Who are the authors of this project?

Minglei Li
