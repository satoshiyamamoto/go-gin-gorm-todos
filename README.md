# go-gin-gorm-todos

Todo app written in Golang.

Implemented with reference to Ekta Garg's [blog post](https://medium.com/@_ektagarg/golang-a-todo-app-using-gin-980ebb7853c8)

## Getting Started

Running MySQL with Docker.

### Prerequisites

```
docker run -d \
  -e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
  -e MYSQL_ROOT_HOST=% \
  -e MYSQL_DATABASE=test \
  -p 3306:3306 \
  mysql:5.7.23
```

### RUN the server

```
go run main.go
```
