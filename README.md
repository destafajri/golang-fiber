# Go-Lang Fiber Project

Sample Go-Lang Fiber Project Structure

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
- `installing migrator tools` download from [golang migrate](https://github.com/golang-migrate/migrate) in release page
- run
```
make migrate-up
```
- run the golang server
```
make run
```

## Framework

- Web : GoFiber
- Validation : Go-Ozzo
- Configuration : GoDotEnv
- Database : MongoDB, Postgre(Supabase), MySQL

## Architecture

Controller -> Service -> Repository

## Addition 

- Entity is representing database table
- Model is representing payload and response
- Helper is representing anything what you need to help your coding process
