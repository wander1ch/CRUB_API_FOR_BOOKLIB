FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./cmd/api
EXPOSE 8080

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /main .
RUN chmod +x /app/main
EXPOSE 8080
CMD ["/app/main"]