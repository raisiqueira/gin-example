FROM golang:1.12-alpine3.9

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR $GOPATH/src/github.com/raisiqueira/gin-example

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["go", "run", "src/main.go"]