# Go-Lang Fiber Project

Sample Go-Lang Fiber Project

## How to use
- Please clone or download this repository.
- Prepare postgres database, or use docker, you can type
```
docker-compose up
```
OR
```
docker-compose up -d
```
- add .env file to setup your database connection
```
POSTGRES_URL="user=postgres password=[your-password] host=[your-host] port=5432 dbname=postgres"
KEY_JWT="[type your secret key here]"
```
- run the golang server
```
go run main.go
```

## Framework

- Web : GoFiber
- Validation : Go-Ozzo
- Configuration : GoDotEnv
- Database : MongoDB

## Architecture

Controller -> Service -> Repository