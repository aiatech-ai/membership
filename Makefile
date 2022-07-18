DB_URL=mysql://root:@tcp(localhost:3306)/shiny?multiStatements=true

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
