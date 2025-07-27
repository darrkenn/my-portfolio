FROM golang:alpine
LABEL authors="darragh"

WORKDIR /app
COPY . .

#App build
RUN go build -o main .
#WebUI port
EXPOSE 1375
CMD ["./main"]