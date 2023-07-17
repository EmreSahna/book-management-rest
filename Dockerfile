FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]