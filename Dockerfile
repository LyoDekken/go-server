FROM golang:1.17.5-alpine3.15

WORKDIR /app

COPY ./cmd .

RUN go build -o main .

CMD ["./main"]
