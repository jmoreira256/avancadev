FROM golang:1.15

WORKDIR /avancaDev/Desafio3

COPY . .

# Instalar dependencias
RUN go get github.com/joho/godotenv && \
  go get github.com/wesleywillians/go-rabbitmq/queue

RUN GOOS=linux go build

EXPOSE 9090

ENTRYPOINT ["./Desafio3"]