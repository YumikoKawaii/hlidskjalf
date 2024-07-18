FROM golang:1.22 AS builder

WORKDIR /app

COPY ../. .

RUN go mod download
RUN go build -o application ./applications/pawn/cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder ./app/application .

EXPOSE 8080

CMD ["application"]