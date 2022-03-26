install:
	go install github.com/swaggo/swag/cmd/swag@latest

generate:
	swag init -ot go --parseDependency

docker-clean:
	docker system prune -a

docker-build:
	docker buildx build --platform linux/amd64 -t kedai-usersvc .
	docker tag kedai-usersvc alganbr/kedai-usersvc
	docker push alganbr/kedai-usersvc

docker-rebuild:
	make docker-clean
	make docker-build

migrate-up:
	migrate -path internal/migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=disable" -verbose down