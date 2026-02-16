# Runtime stage
FROM golang:1.24.4-alpine

# Instalar dependências necessárias
RUN apk add --no-cache git make ca-certificates

# Instalar ferramentas Go necessárias
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/google/wire/cmd/wire@latest

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar todo o código fonte e Makefile
COPY . .

# Expor porta padrão
EXPOSE 8080

# Comando para executar a aplicação usando Makefile
CMD ["make", "dev"]
