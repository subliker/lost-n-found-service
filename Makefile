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
	find logs -type f -name "*.log" -delete

clean-all: clean-logs
	docker compose down -v

set-example:
	cat .env.example > .env

start-example: set-example up