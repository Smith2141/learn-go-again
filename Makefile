.PHONY: run

up:
	@docker compose -f docker-compose.yml up -d

down:
	@docker compose down

ps:
	@docker compose ps

logs:
	@docker compose logs -f

run:
	@docker compose exec app go run main.go