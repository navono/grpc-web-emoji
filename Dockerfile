FROM golang:1.11 as builder
WORKDIR /home/zen/sourcecode/GO/src/github.com/navono/gRPC-Web-emoji/
COPY ./ .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -a -installsuffix cgo -v -o emoji-service main.go

FROM scratch
WORKDIR /bin/
COPY --from=builder /home/zen/sourcecode/GO/src/github.com/navono/gRPC-Web-emoji/emoji-service .

ENTRYPOINT [ "/bin/emoji-service" ]
CMD [ "9000" ]
EXPOSE 9000