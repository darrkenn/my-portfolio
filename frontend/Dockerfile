FROM golang:alpine
LABEL authors="darragh"

#Initial setup
WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
RUN apk add build-base

RUN mkdir db
RUN mkdir md
RUN mkdir info


#.env values
ENV DATABASE_LOCATION=db/portfolio.db \
    CW_LOCATION=info/cw.json \
    T_LOCATION=info/tech.json \
    S_LOCATION=info/song.json \
    MD_LOCATION=md/

#Go building
RUN go mod tidy
RUN go build -o main .
EXPOSE 1375
CMD ["./main"]