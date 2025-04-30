# go-junior

Создать .env file где будет информации с Postgresql:
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=0000
DB_NAME=check
API_PORT=8080

запускаем код с помощью 
```bash
go run cmd/go-junior/main.go 

```
или через билд 
```bash 
go build -o j ./cmd/go-junior/main.go

```
and just:

```
./j
```