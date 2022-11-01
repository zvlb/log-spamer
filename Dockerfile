##
## Build
##
FROM golang:1.18-alpine3.16 AS builder

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o bin/spamer ./cmd/app/main.go

##
## Deploy
##
FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/bin/spamer .

CMD ["/app/spamer"]
