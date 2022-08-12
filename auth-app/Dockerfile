FROM golang:alpine as gobuilder

RUN apk add git bash
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum

RUN go mod download

ADD . /app

COPY .env.example .env
COPY ./wait-for-it.sh .
RUN go build -o server .
ENTRYPOINT ["./wait-for-it.sh", "db:5432", "--", "./server"]
