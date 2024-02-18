include .env

up:
	docker compose up -d
	docker ps

down:
	docker compose down
	docker ps

hash:
	atlas migrate hash --dir "file://db/migrations" 

migrate:
	atlas migrate apply \
  	--url "postgres://${DB_USER}:${DB_PASS}@0.0.0.0:5432/${DB_NAME}?search_path=public&sslmode=disable" \
  	--dir file://db/migrations \
  	--dry-run

boiler:
	sqlboiler psql

gui:
	open http://localhost:8081/