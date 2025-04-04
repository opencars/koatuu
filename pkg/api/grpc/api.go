package grpc

import (
	"context"
	"net"

	"github.com/opencars/grpc/pkg/koatuu"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/opencars/koatuu/pkg/domain"
)

// API represents gRPC API server.
type API struct {
	addr string
	s    *grpc.Server
	svc  domain.InternalService
}

func New(addr string, svc domain.InternalService) *API {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			RequestLoggingInterceptor,
		),
	}

	return &API{
		addr: addr,
		svc:  svc,
		s:    grpc.NewServer(opts...),
	}
}

func (a *API) Run(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	koatuu.RegisterServiceServer(a.s, &handler{api: a})
	reflection.Register(a.s)

	errors := make(chan error)
	go func() {
		errors <- a.s.Serve(listener)
	}()

	select {
	case <-ctx.Done():
		a.s.GracefulStop()
		return <-errors
	case err := <-errors:
		return err
	}
}
