FROM golang:1.18.1-alpine3.15 AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o ext-authz cmd/main.go

FROM alpine:3.15.4
WORKDIR /app
COPY --from=builder /app/ext-authz ./
CMD ["sh", "-c", "/app/ext-authz"]
