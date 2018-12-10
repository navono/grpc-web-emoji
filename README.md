# gRPC-Web-Emoji

# dev

- 从 `https://github.com/protocolbuffers/protobuf/releases` 下载相应平台的 `proto`
- 安装 Golang 版的插件
  ```js
    go get -u github.com/golang/protobuf/proto
    ```
  
- 从 `https://github.com/grpc/grpc-web/releases` 下载安装 `protoc-gen-grpc-web`

# 问题
目前安装的 `protobuf` 和 `protc-gen-go` 产生了点问题.在运行 `main.go` 时,会提示
> emoji/emoji.pb.go:23:11: undefined: proto.ProtoPackageIsVersion3

错误.解决方案:
- 在 `Gopkg.toml` 中加入
    > required = ["github.com/golang/protobuf/protoc-gen-go",]
- 安装 `vendor` 中的 `protoc-gen-go`
    > go install ./vendor/github.com/golang/protobuf/protoc-gen-go