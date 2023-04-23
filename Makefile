postgres:
	docker run --name postgres15 --network flockmanager-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=85dilanwest -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root flock_manager

dropdb:
	docker exec -it postgres15 dropdb flock_manager

migrateup:
	migrate -path db/migration -database "postgresql://root:%7CFY(Kf3-_P42-YC4uui_zTDr%3Fk0D@flock-manager.citdezvqnc1s.us-east-2.rds.amazonaws.com:5432/flock_manager" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:85dilanwest@localhost:5432/flock_manager?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:85dilanwest@localhost:5432/flock_manager?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:85dilanwest@localhost:5432/flock_manager?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test  -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ngenohkevin/flock_manager/db/sqlc Store


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1

#	migrate -path db/migration -database "postgresql://root:%7CFY(Kf3-_P42-YC4uui_zTDr%3Fk0D@flock-manager.citdezvqnc1s.us-east-2.rds.amazonaws.com:5432/flock_manager" -verbose up