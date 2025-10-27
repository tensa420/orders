FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server ./cmd

FROM alpine:latest

WORKDIR /app

RUN addgroup -S -g 1000 appgroup && \
adduser -S -u 1000 -G appgroup appuser

COPY --from=builder /app/server .

COPY --from=builder /app/migrations ./migrations

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./server"]