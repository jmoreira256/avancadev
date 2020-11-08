FROM golang:1.15

WORKDIR /avancaDev/Desafio3

COPY . .

# Instalar dependencias
#RUN go get github.com/joho/godotenv && \
#  go get github.com/wesleywillians/go-rabbitmq/queue && \
#  go get github.com/satori/go.uuid

RUN GOOS=linux go build b.go

EXPOSE 9091

ENTRYPOINT ["./Desafio3"]