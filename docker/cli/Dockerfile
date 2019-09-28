FROM golang:1.12 AS builder
RUN mkdir /app
RUN curl https://gist.githubusercontent.com/donvito/23141efc95e22d0b275480b65dde53b4/raw/00027e00a9a14d9ad78bcfd59137c3d68339edf1/main.go  --output /app/main.go
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]
