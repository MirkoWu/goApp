FROM golang:latest

ENV GOPROXY https://goproxy.io,direct
WORKDIR $GOPATH/src/github.com/mirkowu/go-gin-demo
COPY . $GOPATH/src/github.com/mirkowu/go-gin-demo
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./go-gin-demo"]