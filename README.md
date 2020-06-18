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
└── services
    ├── auth.go
    ├── mocks
    │   └── users.go
    ├── users.go
    └── user_test.go

```

## Structure
### Config
All the config files are placed under `config`. Since I'm doing development. I only create development files here.

### Endpoint
Endpoint is registering routes for endpoints, parsing parameters from request and returning JSON response.

### Services
There are many single services like `user` and `auth`. In this layer, it will collect data from db, compose it and then return to endpoint layer
`mocks` contains mocked functions for unit test.

### Models
This layer is for defining gorm models and doing basic serializing issue.

### Lib
There are several functions here for parsing config files, initializing database and unifying error handling.

#### Middlewares
`middlewares` contains a middleware for jwt including basic authorization and validation.

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
I'm using MySQL for this project.

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