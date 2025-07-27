FROM golang:alpine
LABEL authors="darragh"

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main .
RUN mkdir db

EXPOSE 1375
CMD ["./main"]