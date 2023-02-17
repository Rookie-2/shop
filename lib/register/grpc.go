package register

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type GrpcConfig struct {
	Name         string
	Addr         string
	Opts         []grpc.ServerOption
	RegisterFunc func(server *grpc.Server)
}

func RunGrpcServer(ctx context.Context, cfg *GrpcConfig) error {
	listen, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer(cfg.Opts...)
	cfg.RegisterFunc(s)
	return s.Serve(listen)
}
