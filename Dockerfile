# builder image
FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache bash git

COPY . /go/src/github.com/flameous/pandahack-2018
WORKDIR /go/src/github.com/flameous/pandahack-2018

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o=/app main.go

# resulting image
FROM alpine:latest

WORKDIR /
COPY --from=builder /app /usr/bin/
RUN chmod +x /usr/bin/app

CMD app
