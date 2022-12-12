migrate-up-mysql:
	@migrate -path ./migrations/mysql -database "mysql://root:secret@tcp(localhost:3306)/sample" -verbose up
migrate-down-mysql:
	@migrate -path ./migrations/mysql -database "mysql://root:secret@tcp(localhost:3306)/sample" -verbose down