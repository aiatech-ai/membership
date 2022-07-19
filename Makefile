DB_URL=mysql://root:@tcp(localhost:3306)/shiny?multiStatements=true

init:
	docker-compose up mysql
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
.PHONY: migrateup migratedown init