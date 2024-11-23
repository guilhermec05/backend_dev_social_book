# Use uma imagem base do Golang
FROM golang:1.22

# Instale ferramentas necessárias
RUN apt-get update && apt-get install -y \
    git \
    && apt-get clean

# Instale o Delve (debugger)
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos do projeto
COPY . .

# Baixe as dependências
RUN go mod tidy

# Comando padrão: o `docker-compose.yml` vai substituir o comando por um apropriado para debug
CMD ["go", "run", "main.go"]
