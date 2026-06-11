# Makefile сервиса Shinobi.
# Список целей — `make help`.

SERVICE_NAME := shinobi-service
MODULE       := github.com/PianyCoder/naruto-kage-crm/shinobi-service
BIN_DIR      := bin
MAIN_PKG     := ./cmd/shinobi_service

.PHONY: help build build-check run fmt lint test gen clean

# Цель по умолчанию — справка.
.DEFAULT_GOAL := help

help: ## 📜 Свиток техник: показать доступные цели
	@echo "🍥 忍 $(SERVICE_NAME) — доступные техники:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'

build: ## 🗡 Сборка: выковать бинарь в bin/
	@echo "🗡 忍 Сборка: компилируем $(MAIN_PKG) -> $(BIN_DIR)/$(SERVICE_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(SERVICE_NAME) $(MAIN_PKG)
	@echo "✅ 🗡 Сборка завершена: клинок готов — $(BIN_DIR)/$(SERVICE_NAME)"

build-check: ## 👁 Проверка компиляции: go build без артефактов (шаринган видит всё)
	@echo "👁 Проверка компиляции: go build ./... (без артефактов)..."
	@go build ./...
	@echo "✅ 👁 Компиляция чистая: гендзюцу не обнаружено"

run: ## 🍃 Запуск: поднять сервис локально
	@echo "🍃 Запуск: шиноби покидает Коноху — go run $(MAIN_PKG)..."
	@go run $(MAIN_PKG)

fmt: ## 🌀 Форматирование: gofumpt + goimports
	@echo "🌀 Форматирование: крутим расенган — gofumpt..."
	@gofumpt -l -w .
	@echo "🌀 Форматирование: выравниваем импорты — goimports..."
	@goimports -l -w .
	@echo "✅ 🌀 Форматирование завершено: поток чакры выровнен"

lint: ## 🐸 Линтер: golangci-lint
	@echo "🐸 Линтер: мудрец-жаба изучает свитки — golangci-lint run..."
	@golangci-lint run
	@echo "✅ 🐸 Линтер пройден: мудрец доволен"

test: ## ⚡ Тесты: go test ./...
	@echo "⚡ Тесты: Чидори! — go test ./..."
	@go test ./...
	@echo "✅ ⚡ Тесты пройдены: все цели поражены"

gen: ## 🦊 Генерация: Go-код из proto (protoc)
	@echo "🦊 Генерация: печать призыва — api/shinobi/shinobi.proto -> pkg/api/generated/..."
	@protoc --proto_path=api \
		--go_out=. --go_opt=module=$(MODULE) \
		--go-grpc_out=. --go-grpc_opt=module=$(MODULE) \
		api/shinobi/shinobi.proto
	@echo "✅ 🦊 Генерация завершена: Курама отозвался на призыв"

clean: ## 🔥 Очистка: удалить артефакты сборки (bin/)
	@echo "🔥 Очистка: Катон — Гокакю но дзюцу! Удаляем $(BIN_DIR)/..."
	@rm -rf $(BIN_DIR)
	@echo "✅ 🔥 Очистка завершена: пепел развеян"
