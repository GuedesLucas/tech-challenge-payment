# Use a imagem oficial do Go como a imagem base
FROM golang:1.21

# Crie e defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código-fonte para o contêiner
COPY . .

# Construa o executável
RUN go build

# Exponha a porta em que a aplicação Go estará ouvindo
EXPOSE 7575

# Comando para iniciar a aplicação quando o contêiner for executado
CMD ["./tech-challenge-payment"]