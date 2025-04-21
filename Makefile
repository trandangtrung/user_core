ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

DB_URL=postgresql://root:secret@localhost:5432/CORE_USER?sslmode=disable
PATH_DB=db/migration


migrateup:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path "$(PATH_DB)" -database "$(DB_URL)" -verbose down 1

dev:
	go run main.go dev

production:
	go run main.go production

.PHONY: migrateup migratedown migrateup1 migratedown1 dev production

include ./hack/hack-cli.mk
include ./hack/hack.mk