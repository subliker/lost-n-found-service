.PHONY: build up down rebuild clean-logs

build:
	docker compose build

up:
	docker compose up --build -d

restart:
	docker compose restart

down:
	docker compose down

rebuild:
	docker compose up -d --build $(c)

clean-logs:
	rm -f ./logs/*.log

set-example:
	cat .env.example > .env

start-example: set-example up