protoc -I emoji/ emoji/emoji.proto --js_out=import_style=commonjs:emoji \
       --grpc-web_out=import_style=commonjs,mode=grpcwebtext:emoji