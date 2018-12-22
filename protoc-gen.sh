protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:pkg/api/v1 emoji.proto
protoc --proto_path=api/proto/v1 --js_out=import_style=commonjs:client/web-client/api/v1 \
       --grpc-web_out=import_style=commonjs,mode=grpcwebtext:client/web-client/api/v1 emoji.proto