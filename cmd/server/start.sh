#!/bin/sh

migrate -path ./migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@db:${DB_PORT}/${DB_NAME}?sslmode=disable" up
./main