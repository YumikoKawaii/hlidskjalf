FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o /app/skidbladnir-server ./applications/skidbladnir/cmd

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata iptables

RUN adduser -D -u 1337 -g '' skidbladnir

WORKDIR /app

COPY --from=builder /app/skidbladnir-server .

RUN chown -R skidbladnir:skidbladnir /app

USER 1337

EXPOSE 15001

ENTRYPOINT ["./skidbladnir-server"]
CMD ["serve"]
