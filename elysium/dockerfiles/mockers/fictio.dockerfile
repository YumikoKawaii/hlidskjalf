FROM golang:1.22 AS builder

WORKDIR /

COPY ./elysium .

RUN go mod download

RUN go build -o application ./applications/mockers/fictio/cmd/main.go

FROM ubuntu:latest

WORKDIR /

COPY --from=builder ./application ./bin/application

EXPOSE 8080
EXPOSE 8081

CMD ["./bin/application", "serve"]