DB_URL=mysql://admin:abc123@tcp(localhost:3306)/shiny?multiStatements=true

init:
	docker-compose up mysql -d
uninit: 
	docker-compose down -v
up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
.PHONY: migrateup migratedown init uninit