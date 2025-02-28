# Etapa 1: Construção (Build Stage)
FROM golang:1.23 AS builder

# Atualizando o apt-get e instalando dependências
RUN apt-get update && apt-get install -y

# Definir o diretório de trabalho no container
WORKDIR /go/src/app

# Copiar go.mod e go.sum para aproveitar o cache do Docker
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o restante do código fonte
COPY . .

# Instalar dependências extras (mockery, swag)
RUN go install github.com/vektra/mockery/v2@v2.43.0
RUN go install github.com/swaggo/swag/cmd/swag@v1.7.8