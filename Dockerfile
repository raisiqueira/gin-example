FROM golang:1.12-alpine3.9

LABEL maintainer="me@raisiqueira.dev"
LABEL version="1.0.1"

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR $GOPATH/src/github.com/raisiqueira/gin-example

COPY . .

RUN go get -d -v ./... \
    && go install -v ./...

RUN CGO_ENABLED=0 go build -a -installsuffix .

EXPOSE 8080

ENTRYPOINT ["./gin-example"]