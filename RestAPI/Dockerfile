FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY . .

RUN go build -o main main.go

ENV GIN_MODE=release

EXPOSE 8080

CMD [ "/app/main" ]
