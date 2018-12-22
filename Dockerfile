FROM golang:1.11 as builder
WORKDIR  $GOPATH/src/github.com/navono/gRPC-Web-emoji

# 开启 go mod
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux

# 设置代理，因为需要下载 google 的包
ENV http_proxy=http://192.168.0.103:10086
ENV https_proxy=http://192.168.0.103:10086

# 拷贝 mod 文件
COPY ./go.mod ./go.sum ./

# 下载依赖包
RUN go mod download

COPY ./ .
RUN go build -a -installsuffix cgo -v -o /emoji-service ./cmd/server/main.go

FROM scratch
WORKDIR /bin/
COPY --from=builder /emoji-service .

ENTRYPOINT [ "/bin/emoji-service" ]
CMD [ "9000" ]
EXPOSE 9000