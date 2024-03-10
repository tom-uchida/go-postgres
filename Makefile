include .env

up:
	docker compose up -d
	docker ps -a

down:
	docker compose down
	docker ps -a

atlas-hash:
	atlas migrate hash \
	--dir "file://db/sql"

# スキーマの適用
atlas-schema-init:
	atlas schema apply \
	--url "postgres://${DB_USER}:${DB_PASS}@localhost:5432/${DB_NAME}?search_path=public&sslmode=disable" \
	--dev-url "docker://postgres/15/dev?search_path=public" \
	--file file://db/sql/schema.sql

# マイグレーションファイルの生成
atlas-migrate-diff:
	atlas migrate diff initial \
	--dir file://db/migrations \
	--to file://db/sql/schema.sql \
	--dev-url "docker://postgres/15/dev?search_path=public" \
	--format '{{ sql . "  " }}'

# # マイグレーションの実行
# atlas-migrate-apply:
# 	atlas migrate apply \
# 	--dir file://db/migrations \
#   	--url "postgres://${DB_USER}:${DB_PASS}@localhost:5432/${DB_NAME}?search_path=public&sslmode=disable" \
#   	--dry-run

atlas-schema-clean:
	atlas schema clean \
	--url "postgres://${DB_USER}:${DB_PASS}@localhost:5432/${DB_NAME}?search_path=public&sslmode=disable"

sqlboiler:
	sqlboiler psql -c ./sqlboiler.toml -o ./db/models --no-tests

gui:
	open http://localhost:8081/

run:
	go run ./