FROM golang:1.22 as build
WORKDIR /app
COPY . .
EXPOSE 8000
RUN go build -o main main.go
CMD ["./main"]