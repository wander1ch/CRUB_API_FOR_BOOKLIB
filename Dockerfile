FROM 1.26.3-alpine3.23
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build /main ./cmd/api
EXPOSE 8000
CMD ["./main"]