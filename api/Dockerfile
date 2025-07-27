FROM golang:alpine
LABEL authors="darragh"

WORKDIR /app
COPY . .

RUN go build -o main .
EXPOSE 5731
CMD ["./main"]