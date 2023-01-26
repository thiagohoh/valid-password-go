ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /passwordvalidation
WORKDIR /passwordvalidation

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./pswvalidation ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /passwordvalidation
WORKDIR /passwordvalidation
COPY --from=builder /passwordvalidation/pswvalidation .

ENV GIN_MODE release
EXPOSE 8080

ENTRYPOINT ["./pswvalidation"]
