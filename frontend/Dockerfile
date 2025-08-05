FROM golang:alpine
LABEL authors="darragh"

#Initial setup
WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
RUN apk add --no-cache build-base

RUN mkdir db
RUN mkdir md
RUN mkdir info
RUN mkdir message


RUN go mod tidy
RUN go mod vendor
RUN go build -o main .
EXPOSE 1375
CMD ["./main"]
