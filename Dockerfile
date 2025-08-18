# Go base image
FROM golang:1.24 AS builder

WORKDIR /app

# Go mod dosyalarını kopyala ve bağımlılıkları indir
COPY go.mod go.sum ./
RUN go mod download

# Kaynak kodu kopyala
COPY . .

# Binary build et
RUN go build -o main .

# Final stage - daha küçük bir image
FROM debian:bookworm-slim

WORKDIR /app

# gerekli CA sertifikaları ve PostgreSQL client
RUN apt-get update && \
    apt-get install -y ca-certificates postgresql-client && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .
COPY wait-for-db.sh ./
COPY .env ./
RUN chmod +x wait-for-db.sh

EXPOSE 8080

CMD ["./wait-for-db.sh", "./main"]
