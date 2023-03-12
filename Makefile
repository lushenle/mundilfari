# Image URL to use all building/pushing image targets
IMG ?= ishenle/simplebank:latest

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
	docker exec postgres12 createdb --username=myuser --owner=myuser simplebank

.PHONY: dropdb
dropdb:
	docker exec postgres12 dropdb -U myuser simplebank

.PHONY: migrateup
migrateup:
	migrate -path db/migration -database "postgresql://myuser:mypass@localhost:5432/simplebank?sslmode=disable" -verbose up

.PHONY: migratedown
migratedown:
	migrate -path db/migration -database "postgresql://myuser:mypass@localhost:5432/simplebank?sslmode=disable" -verbose down

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: test
test:
	go test -v -count=1 -cover ./...

.PHONY: server
server: fmt vet
	go run main.go

.PHONY: docker-build
docker-build:
	docker build -t ${IMG} .

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: docker
docker: docker-build docker-push
