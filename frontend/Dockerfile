FROM golang:alpine
LABEL authors="darragh"

#Initial setup
WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
RUN mkdir db
RUN mkdir md
RUN apk add build-base
RUN touch .env
RUN echo "DATABASE_LOCATION=db/portfolio.db" > .env
RUN echo "MARKDOWN_LOCATION=md/"
#Go building
RUN go mod tidy
RUN go build -o main .
EXPOSE 1375
CMD ["./main"]