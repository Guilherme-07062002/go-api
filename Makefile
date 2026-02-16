# Variáveis
BINARY_NAME=go-api
MAIN_PATH=./main.go

.PHONY: all build test clean run swag wire tidy

# Comando padrão: limpa, gera código e compila
all: tidy swag wire build

# Sincroniza os módulos do Go (adiciona faltantes e remove lixo)
tidy:
	@echo "Cleaning modules..."
	go mod tidy

# Gera a documentação do Swagger
swag:
	@echo "Generating Swagger docs..."
	swag init --parseDependency --parseInternal

# Gera o código de Injeção de Dependência (Wire)
wire:
	@echo "Generating DI container..."
	cd infra/config/wire && wire

# Compila o binário do projeto
build:
	@echo "Building binary..."
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Roda os testes com cobertura (padrão de mercado)
test:
	@echo "Running tests..."
	go test -v -cover ./...

# Roda a aplicação com o ciclo completo de automação
dev: swag wire tidy
	@echo "Starting application in dev mode..."
	go run $(MAIN_PATH)

# Limpa binários e arquivos temporários
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -rf docs/