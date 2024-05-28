.PHONY: up
up:
	docker compose up backend frontend db

.PHONY: init
init:
	docker compose up --build

.PHONY: docs
docs:
	docker compose up swagger-editor swagger-ui

.PHONY: test_repository
test_repository:
	cd backend && DB_HOSTNAME=127.0.0.1 go test -v ./test/repositories/...
