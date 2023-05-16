FROM golang:alpine

# Install air
RUN go install github.com/cosmtrek/air@latest

# Intall go migrate tool
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install tesseract-ocr
RUN apk --update --no-cache add tesseract-ocr=5.3.1-r2

RUN mkdir main

WORKDIR /app

COPY go.* ./
RUN go mod download && go mod verify

COPY . .

#May need to run chmod +x ./start.sh locally if mounting host to container
RUN chmod +x ./start.sh 

RUN go build -o ../main/main ./cmd/server

WORKDIR ../app
EXPOSE 8080

ENTRYPOINT [ "./start.sh" ] 