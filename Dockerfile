FROM golang:1.24.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main .
CMD ["/app/main"]
