createdb_users:
	docker exec -it users_db createdb --username=db_user --owner=db_user users

createdb_pizzeria:
	docker exec -it pizzeria_db createdb --username=db_user --owner=db_user pizzeria

dropdb_user:
	docker exec -it users_db dropdb --username=db_user users

dropdb_pizzeria:
	docker exec -it pizzeria_db dropdb --username=db_user pizzeria

migrateup_users:
	migrate -path ./internal/db/migrations/users -database "postgres://db_user:db_pass@localhost:5431/users?sslmode=disable" up

migrateup_pizzeria:
	migrate -path ./internal/db/migrations/pizzeria -database "postgres://db_user:db_pass@localhost:5432/pizzeria?sslmode=disable" up

migratedown_pizzeria:
	migrate -path ./internal/db/migrations/pizzeria -database "postgres://db_user:db_pass@localhost:5432/pizzeria?sslmode=disable" down 1

migratedown_users:
	migrate -path ./internal/db/migrations/users -database "postgres://db_user:db_pass@localhost:5431/users?sslmode=disable" down 1

.PHONY:createdb_pizzeria createdb_users dropdb_pizzeria dropdb_user migratedown_pizzeria migratedown_users migrateup_pizzeria migrateup_users