# Build stage
FROM golang:1.25-bookworm AS builder

ENV CGO_ENABLED=1

RUN apt-get update && \
    apt-get install -y \
    tesseract-ocr \
    libtesseract-dev \
    libleptonica-dev \
    pkg-config

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /app/uploads && chmod 777 /app/uploads

RUN go build -o ocr-api main.go


# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
    tesseract-ocr \
    libtesseract5 \
    libleptonica-dev \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/ocr-api .
COPY --from=builder /app/uploads ./uploads

EXPOSE 8080

CMD ["./ocr-api"]