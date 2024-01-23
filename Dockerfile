# Use a imagem oficial do Go como imagem base
FROM golang:latest

# Configure o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código fonte para o diretório de trabalho
COPY . .

# Compile o código
RUN go build -o main .

# Exponha a porta que a aplicação está utilizando
EXPOSE 8080

# Comando padrão para executar a aplicação
CMD ["./main"]
