package game

import (
	"context"

	"github.com/ohanan/bpg/internal/common"
	"github.com/ohanan/bpg/internal/proto"
	"google.golang.org/grpc"
)

type Game interface {
	Start()
}

func Run(game Game) error {
	return common.StartServer(func(server *grpc.Server) {
		proto.RegisterGameServer(server, &gameServer{})
	})
}

type gameServer struct {
	proto.UnimplementedGameServer
}

func (g gameServer) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}
