DATABASE_HOST = 
DATABASE_NAME = 
DATABASE_USER = 
DATABASE_PASSWORD = 
DATABASE_PORT = 

up:
	migrate -database "postgresql://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" -path db/migrations up

down:
	migrate -database "postgresql://$(DATABASE_USER):$(DATABASE_PASSWORD)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable" -path db/migrations down

run:
	go run .

doc:
	swag init