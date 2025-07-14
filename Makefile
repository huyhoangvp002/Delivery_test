DB_URL=postgresql://root:secret@localhost:6543/Delivery?sslmode=disable
postgres:
	sudo docker run --name postgresql1 -p 6543:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	sudo docker exec -it postgresql1 createdb --username=root --owner=root Delivery
dropdb:
	sudo docker exec -it postgresql1 dropdb  Delivery
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
sqlc:
	sqlc generate	
run:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc mock run db