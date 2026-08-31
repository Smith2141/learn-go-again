.PHONY: up down shell dev debug build logs run run-env clean

# Управление контейнером
up:
	@docker compose up -d

down:
	@docker compose down

build:
	@docker compose build

logs:
	@docker compose logs -f

# Вход в контейнер
shell:
	@docker exec -it go-dev /bin/sh

# Запуск приложения
run:
	@docker compose exec app go run main.go

# Запуск с переменной окружения
run-env:
	@docker compose exec -e TEST_YEAR=1943 app go run main.go

# Запуск с переменной из аргумента
run-year:
	@docker compose exec -e TEST_YEAR=$(YEAR) app go run main.go

# Разработка с Air (исправленный)
dev:
	@docker compose exec app sh -c "air -c .air.toml"

# Отладка (исправленная)
debug:
	@docker compose exec app sh -c "pkill -f dlv 2>/dev/null; TEST_YEAR=1990 dlv debug --buildvcs=false --headless --listen=:2345 --api-version=2 --accept-multiclient"

# Отладка с интерактивным вводом
debug-interactive:
	@docker compose exec -it app sh -c "pkill -f dlv 2>/dev/null; dlv debug --buildvcs=false --headless --listen=:2345 --api-version=2 --accept-multiclient"

# Очистка
clean:
	@docker compose down -v
	@docker compose rm -f
	@rm -rf tmp/