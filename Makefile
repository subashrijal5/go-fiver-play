run:
	reflex -r '.go|.html|.css' -s go run .

migrate:
	go run ./utils/database/migrate.go
