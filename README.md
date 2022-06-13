## Simple TODO API

A simple REST TODO API developed in Go. The TODOs are persisted on mysql database. The program can be executed on Docker or as a Go application locally.

It provides an API for manager TODO Tasks. 

These are the endpoints available:

| Operation      | Endpoint    | HTTP Verb |
|----------------|-------------|-----------|
| List all Todos | /todos      | GET       |
| Create Todo    | /todos      | POST      |
| Get Todo       | /todos/{id} | GET       |
| Update Todo    | /todos/{id} | PUT       |
| Delete Todo    | /todos/{id} | DELETE    |

## Build and Running

### As a Go app

```bash
$ go mod tidy

$ go run cmd/main.go
```

### As a Docker Container

```bash
$ docker-compose up -d
```

## Accessing API

```bash
$ curl http://localhost:8000/todos
```