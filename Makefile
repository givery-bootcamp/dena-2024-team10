.PHONY: up
up:
	docker compose up backend frontend db

.PHONY: init
init:
	docker compose up --build

.PHONY: docs
docs:
	docker compose up swagger-editor swagger-ui swagger-api

