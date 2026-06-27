FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o out ./src/main/
EXPOSE 8080
CMD ["./out"]
