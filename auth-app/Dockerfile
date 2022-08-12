FROM golang:alpine as gobuilder

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./server .

FROM alpine:latest
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY .env.example .env
COPY --from=gobuilder /app/server /bin/server
ENTRYPOINT ["/bin/server"]
