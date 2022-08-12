FROM golang:alpine as gobuilder

RUN apk add git bash
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum

RUN go mod download

ADD . /app

COPY .env.example .env
COPY ./wait-for-it.sh ./wait-for-it.sh

RUN CGO_ENABLED=0 go build -o server .

FROM alpine:latest

RUN apk add bash

WORKDIR /root

COPY --from=gobuilder /app/.env /root/
COPY --from=gobuilder /app/wait-for-it.sh /root/
COPY --from=gobuilder /app/server /root/

ENTRYPOINT ["./wait-for-it.sh", "db:5432", "--", "./server"]
