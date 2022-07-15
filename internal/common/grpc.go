package common

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const hostPortPrefix = "--bpg started server at "

func StartServer(onRegister func(server *grpc.Server)) error {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	onRegister(s)
	fmt.Printf(hostPortPrefix+"%s\n", l.Addr().String())
	return s.Serve(l)
}

func NewClient(cmd string) (grpc.ClientConnInterface, error) {
	c := exec.Command(cmd)
	b := commandOutput{
		ch: make(chan string),
	}
	errCh := make(chan error)
	c.Stdout = &b
	go func() { errCh <- c.Run() }()
	defer func() {
		close(b.ch)
		close(errCh)
	}()
	select {
	case err := <-errCh:
		return nil, err
	case s := <-b.ch:
		return grpc.Dial(s, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}

type commandOutput struct {
	ch chan string
}

func (c *commandOutput) Write(p []byte) (n int, err error) {
	s := string(p)
	if !strings.HasPrefix(s, hostPortPrefix) {
		return
	}
	c.ch <- strings.TrimSpace(s[len(hostPortPrefix):])
	return
}
