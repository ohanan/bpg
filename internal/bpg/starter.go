package bpg

import (
	"net"

	"github.com/ohanan/bpg/internal/proto"
	"google.golang.org/grpc"
)

func StartServer() (ipPort string, err error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	s := grpc.NewServer()
	proto.RegisterBpgServer(s, &Bpg{})
	go func() {
		if err := s.Serve(l); err != nil {
			panic(err)
		}
	}()
	return l.Addr().String(), nil
}
