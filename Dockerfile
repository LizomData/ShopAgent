FROM golang:1.24-bullseye AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -trimpath -o /app/shopAgent

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8000

CMD ["/app/shopAgent"]