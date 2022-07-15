package bot

import (
	"context"

	"github.com/ohanan/bpg/internal/common"
	"github.com/ohanan/bpg/internal/proto"
	"google.golang.org/grpc"
)

type Bot interface {
	Init(ctx Context, config *Config) error
}

func Run(bot Bot) error {
	return common.StartServer(func(server *grpc.Server) {
		proto.RegisterBotServer(server, &botServer{})
	})
}

type botServer struct {
	proto.UnimplementedBotServer
}

func (b botServer) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}
