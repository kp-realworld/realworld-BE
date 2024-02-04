FROM golang:1.20-alpine
WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o realworld
CMD ["./realworld","--config=./config/local-env.toml"]
