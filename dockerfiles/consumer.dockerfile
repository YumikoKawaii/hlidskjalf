FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o /app/consumer-server ./applications/consumer/cmd

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/consumer-server .

RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 10080

ENTRYPOINT ["./consumer-server"]
CMD ["serve"]
