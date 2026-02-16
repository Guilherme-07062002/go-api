# Go API - Clean Architecture

Exemplo de API REST para gerenciamento de álbuns implementada com Clean Architecture.

## Stack

- **Go** 1.24.4 + **Gin** (HTTP framework)
- **PostgreSQL** 16 + **GORM** (ORM)
- **JWT** (autenticação)
- **Swagger** (documentação)
- **Wire** (injeção de dependências)
- **Docker** + **Docker Compose**

## Estrutura

```
├── controllers/       # Handlers HTTP
├── usecases/          # Lógica de negócio
├── domain/            # Entidades, DTOs, interfaces
├── infra/             # Implementações (DB, security, config)
└── docs/              # Swagger docs (auto-gerado)
```

## Execução

### Docker
```bash
docker-compose up -d
```

### Local
```bash
make dev        # Desenvolvimento
make dev-watch  # Com hot reload (requer air)
make test       # Testes
```

## Documentação

Swagger disponível em: `http://localhost:8080/swagger/index.html`

## Comandos Make utilitários

```bash
make swag    # Gerar docs Swagger
make wire    # Gerar DI container
make tidy    # Limpar dependências
make build   # Compilar binário
```