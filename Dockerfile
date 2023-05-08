FROM golang:alpine

WORKDIR /app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]