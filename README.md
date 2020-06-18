# Rest API
This is a very simple RESTful API service using gin, gorm and jwt.
The structure is like following
```
.
├── config
│   └── development
│       ├── config.json
│       └── jwtsecret.key
├── endpoint
│   ├── auth.go
│   └── users.go
├── lib
│   ├── config.go
│   ├── database
│   │   └── db.go
│   ├── error.go
│   └── middlewares
│       ├── authorized.go
│       └── jwt.go
├── models
│   └── users.go
```

## Compile and Run
Compile
```
make
```
Run
```
./restapi
```
Run test
```
make test
```
## API
### User
#### List User
> GET localhost:8080/users

Response
```
[
    {
        "id": 1,
        "name": "John",
        "age": 31,
        "city": "New York"
    },
    {
        "id": 2,
        "name": "Doe",
        "age": 22,
        "city": "Vancouver"
    }
]
```
#### Get User
> GET localhost:8080/user/1

Response
```
{
    "id": 1,
    "name": "John",
    "age": 31,
    "city": "New York"
}
```
### Auth
#### Register and get token
> POST localhost:8080/register

Request body
```
{
    "username": "violet",
    "password": "233"
}
```
Response body
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTMxMTEzODMsInVzZXIiOnsiQWdlIjowLCJDaXR5IjoiIiwiaWQiOjMsIm5hbWUiOiJ2aW9sZXQifX0.DFZ7REZuUlJife4cZ3_fSC94TV54R0Yemh6smq1zQt0",
    "user": {
        "Age": 0,
        "City": "",
        "id": 3,
        "name": "violet"
    }
}
```
#### Login and get token
> POST localhost:8080/login

Request body
```
{
    "username": "violet",
    "password": "233"
}
```
Response body
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTMxMTEzODMsInVzZXIiOnsiQWdlIjowLCJDaXR5IjoiIiwiaWQiOjMsIm5hbWUiOiJ2aW9sZXQifX0.DFZ7REZuUlJife4cZ3_fSC94TV54R0Yemh6smq1zQt0",
    "user": {
        "Age": 0,
        "City": "",
        "id": 3,
        "name": "violet"
    }
}
```
## DB schema
| Field | Description | Type | Null | Key |
| ------- | ------------- | ------ | ------ | ----- |
| id    | user id     | int  | No   | Primary |
| name  | user name   | varchar | No  | No |
| age   | user age    | int     | Yes | No |
| city  | city        | varchar | Yes | No |
| password_hash | hashed password | varchar | No | No |

## Authorized API
I use jwt middleware for listing user. It would require token at the first time and then store it in cookie. If no token provided, it'll return 401.

## Error Handling
| condition | status code |
|---|---|
| no config files | panic |
| config file is invalid | panic |
| cannot connect to DB when init | panic |
| cannot load jwt key files | panic |
| DB is down | 500 |
| user is not existed | 404 |
| user is conflict | 409 |
| user is not authorized | 401 |