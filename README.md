# Introduction

Starts web server on port 8080 to fetch data from a [fake users endpoint](https://jsonplaceholder.typicode.com/users) and do some operations over it.

Endpoints:

- `/live`: Health check endpoint
- `/api/v1/users`: Return list of users
- `/api/v1/nearby-user`?lat=123&lng=123: Returns the nearest users from the lat lng provided.


# Quick start

## Run application

[Air](https://github.com/air-verse/air) is a convenient tool to run go with hot reload:

`air "go run main.go"`

If we want to run the app without hot reload:

`go run main.go`

## Curl commands

### Get All Users
`curl localhost:8080/api/v1/users | jq .`

### Get Nearby User
`curl 'localhost:8080/api/v1/nearby?lat=-38.2386&lng=57.2232' | jq .`