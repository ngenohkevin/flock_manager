include app.env

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} --network ${NETWORK} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:15-alpine

createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

dropdb:
	docker exec -it ${DB_DOCKER_CONTAINER} dropdb ${DB_NAME}

migrateup:
	migrate -path db/migration -database "postgresql://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database"postgresql://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test  -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ngenohkevin/flock_manager/db/sqlc Store


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1

