package cmd

import (
	"context"
	"github.com/navono/gRPC-Web-emoji/pkg/protocol/grpc"
	"github.com/navono/gRPC-Web-emoji/pkg/service/v1"
)

func RunServer() error {
	ctx := context.Background()

	v1API := v1.NewEmojiServiceServer()

	return grpc.RunServer(ctx, v1API)
}