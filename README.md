
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
nohup ./perseus/Perseus > perseus.log 2>&1 &
```

## What API do we have?

### Login API
`/login`  
Authenticate users and set cookies.

### Register API
`/register-email`  
User register and send verification email.  
`/register-activate`   
Email verification passed.  

### Reset API
`/password-email`  
Send reset email.  
`/password-reset`  
User reset password.  

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

## Who are the authors of this project?

Minglei Li
