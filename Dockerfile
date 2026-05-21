FROM golang:1.26-alpine AS builder

WORKDIR /src

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/backend ./apps/backend && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/daemon ./apps/daemon

FROM alpine:3.21 AS base-runtime

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /src/configs ./configs

FROM base-runtime AS backend-runtime

COPY --from=builder /out/backend /app/backend

EXPOSE 8888

ENTRYPOINT ["/app/backend"]

FROM base-runtime AS daemon-runtime

COPY --from=builder /out/daemon /app/daemon

ENTRYPOINT ["/app/daemon"]
