# Build stage
FROM golang:1.24.4-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Instalar swag CLI para gerar documentação
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copiar código fonte
COPY . .

# Gerar documentação Swagger
RUN swag init

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest

# Instalar certificados SSL
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar binário do estágio de build
COPY --from=builder /app/main .

# Copiar arquivo .env-example (opcional)
COPY --from=builder /app/.env-example .

# Expor porta padrão
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]
