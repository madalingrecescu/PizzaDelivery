createdb_users:
	docker exec -it users_db createdb --username=db_user --owner=db_user users

dropdb_users:
	docker exec -it users_db dropdb --username=db_user users

migrateup_users:
	migrate -path ./internal/db/migrations/users -database "postgres://db_user:db_pass@localhost:5431/users?sslmode=disable" up

migratedown_users:
	migrate -path ./internal/db/migrations/users -database "postgres://db_user:db_pass@localhost:5431/users?sslmode=disable" down 1

createdb_pizzas:
	docker exec -it pizzas_db createdb --username=db_user --owner=db_user pizzas

dropdb_pizzas:
	docker exec -it pizzas_db dropdb --username=db_user pizzas

migrateup_pizzas:
	migrate -path ./internal/db/migrations/pizzas -database "postgres://db_user:db_pass@localhost:5432/pizzas?sslmode=disable" up

migratedown_pizzas:
	migrate -path ./internal/db/migrations/pizzas -database "postgres://db_user:db_pass@localhost:5432/pizzas?sslmode=disable" down 1

generate_go_server_code_users:
	swagger generate server -A pizzadelivery -f ./configs/swagger/users_swagger.yaml -t ./internal/swagger

generate_go_client_code_users:
	swagger generate client -f ./configs/swagger/users_swagger.yaml -t ./internal/swagger

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
users:
	go run cmd/users/users_main.go

mock_users:
	mockgen -package mockdb_users -destination internal/db/mock/store.go github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users Store

.PHONY:mock_users users createdb_pizzas createdb_users dropdb_pizzas dropdb_user migratedown_pizzas migratedown_users migrateup_pizzas migrateup_users generate_go_server_code_users generate_go_client_code_users sqlc test