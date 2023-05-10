FROM golang:alpine

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .
COPY ./cmd/server/start.sh .

RUN chmod +x ./start.sh

RUN go build -o main ./cmd/server
EXPOSE 8080

ENTRYPOINT [ "./start.sh" ] 