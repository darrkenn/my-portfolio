FROM golang:alpine
LABEL authors="darragh"

WORKDIR /app
COPY . .

ENV CGO_ENABLED=1
RUN apk add --no-cache build-base
RUN go mod tidy
RUN go mod vendor
RUN go build -o main .
RUN mkdir db
EXPOSE 5731
CMD ["./main"]