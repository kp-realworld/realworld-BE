FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o realworld -ldflags="-s -w"
# CMD ["./realworld","--config=./config/dev-env.toml"]

FROM alpine

COPY --from=builder /app ./realworld

EXPOSE 8080

CMD ["./realworld","--config=./config/dev-env.toml"]