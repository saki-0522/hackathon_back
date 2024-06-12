FROM golang:1.21 as build
WORKDIR /app
COPY . .
EXPOSE 8000
RUN go build -o main main.go
CMD ["./main"]