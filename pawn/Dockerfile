FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
