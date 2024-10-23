FROM golang:1.22 AS builder

WORKDIR /

COPY ./elysium .

RUN go mod download

RUN go build -o application ./applications/users/cmd/main.go

FROM ubuntu:latest

WORKDIR /

COPY --from=builder ./application ./bin/application

EXPOSE 8080

CMD ["./bin/application", "serve"]