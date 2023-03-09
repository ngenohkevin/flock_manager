postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=85dilanwest -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root flock_manager

dropdb:
	docker exec -it postgres15 dropdb flock_manager