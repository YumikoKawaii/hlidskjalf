FROM golang:1.22 AS builder

WORKDIR /

COPY ./elysium .

RUN go mod download

RUN go build -o application ./applications/authenticator/cmd/main.go

FROM ubuntu:latest

WORKDIR /root/

COPY --from=builder ./application /bin/application
COPY --from=builder ./applications/authenticator/api_keys_cfg.yaml /media/api_keys_cfg.yaml

EXPOSE 8080
EXPOSE 8081

CMD ["./bin/application", "serve"]