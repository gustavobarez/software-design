.PHONY: start

DB_DSN := postgresql://postgres:example@localhost:5432/database?sslmode=disable
JET_OUT := ./adapter/output
JET_TMP := ./jetgen

start:
	@$(MAKE) --no-print-directory docker-run
	@$(MAKE) --no-print-directory wait-db
	@$(MAKE) --no-print-directory migrate-up
	@$(MAKE) --no-print-directory jet
	@$(MAKE) --no-print-directory run

run:
	@echo "==> (1/1) Iniciando servidor..."
	@go run main.go

docker-run:
	@echo "==> (1/1) Iniciando banco de dados"
	@docker compose up -d

docker-down:
	@echo "==> (1/1) Desligando banco de dados"
	@docker compose down

jet:
	@echo "==> (1/5) Limpando geração anterior..."
	@rm -rf $(JET_TMP)

	@echo "==> (2/5) Gerando código Jet..."
	@jet -dsn="$(DB_DSN)" -ignore-tables=schema_migrations -path=$(JET_TMP)

	@echo "==> (3/5) Criando pasta final..."
	@mkdir -p $(JET_OUT)

	@echo "==> (4/5) Movendo apenas conteúdo dentro de public/..."
	@cp -R $(JET_TMP)/*/public/* $(JET_OUT)/

	@echo "==> (5/5) Removendo pasta temporária..."
	@rm -rf $(JET_TMP)

migrate-up:
	@migrate -database "$(DB_DSN)" -path ./migrations up

wait-db:
	@echo "==> (1/2) Aguardando Postgres ficar pronto..."
	@until docker inspect --format='{{.State.Health.Status}}' postgres 2>/dev/null | grep -q healthy; do \
		sleep 1; \
	done
	@echo "==> (2/2) Postres pronto"
