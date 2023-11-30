
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
git
sudo lsof -i :8080
kill [PID]
nohup ./perseus/Perseus > perseus.log 2>&1 &
```

## What API do we have?

### Login API
`/login`  
Authenticate users and set cookies.

### Register API
`/register`  
User register and set cookies.  

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
`/user-video`  
Record the video that user liked/watched.  
`/user-video-list`  
Get video list that user liked/watched.  

## What tables are in the database?

`user_info`
```
CREATE TABLE user_info (
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
password VARCHAR(255) NOT NULL,
email VARCHAR(255) NOT NULL,
activated BOOLEAN,
vip BOOLEAN
);
```

`video_info`
```
CREATE TABLE video_info (
id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(255) NOT NULL,
type INT UNSIGNED NOT NULL,
total_number INT UNSIGNED NOT NULL,
base_url VARCHAR(255) NOT NULL
);
```

`user_history`
```
CREATE TABLE user_history (
user_id INT UNSIGNED NOT NULL,
video_id INT UNSIGNED NOT NULL,
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
PRIMARY KEY (user_id, video_id),
FOREIGN KEY (user_id) REFERENCES user_info(id),
FOREIGN KEY (video_id) REFERENCES video_info(id)
);
```

## Who are the authors of this project?

Minglei Li
