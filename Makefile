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
pizzas:
	go run cmd/pizzas/pizzas_main.go

mock_users:
	mockgen -package mockdb_users -destination internal/db/mock/users_mock/store_users.go github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_users Store
mock_pizzas:
	mockgen -package mockdb_pizzas -destination internal/db/mock/pizzas_mock/store_pizzas.go github.com/madalingrecescu/PizzaDelivery/internal/db/sqlc_pizzas Store

proto:
	rm -f internal/pb/*.go
	protoc --proto_path=internal/proto --go_out=internal/pb --go_opt=paths=source_relative \
        --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
        internal/proto/*.proto

grpc:
	go run cmd/gRPC/gRPC_main.go

evans:
	evans --host localhost --port 3001 -r repl
.PHONY:evans proto pizzas mock_pizzas mock_users users createdb_pizzas createdb_users dropdb_pizzas dropdb_user migratedown_pizzas migratedown_users migrateup_pizzas migrateup_users generate_go_server_code_users generate_go_client_code_users sqlc test