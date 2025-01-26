include ./apps/backend/.env.local
.PHONY: sqlc

migrate_up:
	cd apps/backend && migrate -path src/db/migration -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate_down:
	cd apps/backend && migrate -path src/db/migration -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down


sqlc:
	cd apps/backend && sqlc generate


test:
	go test -v -cover ./...

server:
	cd apps/backend && air -c .air.toml

protoc_backend:
	cd apps/backend && protoc -I . \
    --go_out ./src/_generated/ --go_opt paths=source_relative \
    --go-grpc_out ./src/_generated/ --go-grpc_opt paths=source_relative \
    proto/v1/*.proto

protoc_gateway:
	cd apps/backend && protoc -I . --grpc-gateway_out ./src/_generated \
    --grpc-gateway_opt paths=source_relative \
    proto/v1/*.proto
