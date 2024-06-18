postgres: 
	docker run --name postgres16 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres16 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./db/sqlc/
mock:
	mockgen -destination=db/mock/store.go github.com/muhammadsaman77/simplebank/db/sqlc Store
.PHONY: createdb dropdb postgres migrateup  migratedown sqlc test







