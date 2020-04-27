FROM golang:1.14
ENV GOPROXY=https://goproxy.io GOPATH="$HOME/privatespace/go" PATH="${PATH}:$GOPATH/bin"

# 定义使用的Golang 版本
ARG GO_VERSION=1.14

WORKDIR $GOPATH/crawler
COPY . $GOPATH/go/crawler

EXPOSE 8080

CMD ["./main"]