FROM golang:1.22 AS builder

WORKDIR /

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY ../. .

RUN go build -o application ./pawn/cmd/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder . .

EXPOSE 8080
#
#CMD ["./application"]

CMD ["while :; do :; done"]