DB_URL=postgresql://adam:A05092003a@localhost:5432/telegram?sslmode=disable

postgres:
	docker run --name tgpostgres -p 5432:5432 -e POSTGRES_USER=adam -e POSTGRES_PASSWORD=A05092003a -d postgres

createdb:
	docker exec -it tgpostgres createdb --username=adam --owner=adam telegram

dropdb:
	docker exec -it tgpostgres dropdb -U adam telegram

migrateup:
	migrate -path pkg/db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path pkg/db/migrations -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir pkg/db/migrations init_schema

server:
	go run main.go


.PHONY: postgres createdb dropdb migrateup migratedown new_migration server