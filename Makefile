# Image URL to use all building/pushing image targets
IMG ?= ishenle/mundilfari:latest
DB_URL = postgresql://myuser:mypass@localhost:5432/mundilfari?sslmode=disable

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: postgres
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mypass -d postgres:12-alpine

.PHONY: createdb
createdb:
	docker exec postgres12 createdb --username=myuser --owner=myuser mundilfari

.PHONY: dropdb
dropdb:
	docker exec postgres12 dropdb -U myuser mundilfari

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

.PHONY: migrateup1
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

.PHONY: migratedown1
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test
test:
	go test -v -count=1 -cover --short ./...

.PHONY: server
server: fmt vet
	go run main.go

.PHONY: mock
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/lushenle/mundilfari/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/lushenle/mundilfari/worker TaskDistributor

.PHONY: docker-build
docker-build:
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: docker
docker: docker-build docker-push

.PHONY: db_docs
db_docs:
	dbdocs build doc/db.dbml

.PHONY: db_schema
db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

.PHONY: proto
proto:
	rm -rf pb/*.go
	rm -rf doc/swagger/*.swagger.jon
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=mundilfari \
        proto/*.proto

.PHONY: evans
evans:
	evans --host localhost --port 9090 -r repl

.PHONY: redis
redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine
