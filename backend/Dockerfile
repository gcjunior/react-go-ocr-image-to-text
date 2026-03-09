FROM golang:1.25

RUN apt-get update && \
    apt-get install -y tesseract-ocr libtesseract-dev libleptonica-dev pkg-config

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Create uploads folder
RUN mkdir -p /app/uploads && chmod 777 /app/uploads

CMD ["go", "run", "main.go"]