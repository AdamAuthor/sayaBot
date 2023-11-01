DB_URL=postgresql://root:secret@localhost:5432/tgpostgres?sslmode=disable

postgres:
	docker run --name tgpostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it tgpostgres createdb --username=root --owner=root bot

dropdb:
	docker exec -it postgres dropdb tgpostgres

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir pkg/migration -seq tgpostgres

server:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown server