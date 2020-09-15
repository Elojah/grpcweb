package grpcweb

import (
	"crypto/tls"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Service embed a connected redis client.
type Service struct {
	Register func()

	*grpc.Server
	*grpcweb.WrappedGrpcServer
}

// Dial connects client to external redis service.
func (s *Service) Dial(cfg Config) error {

	cert, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key)
	if err != nil {
		return err
	}

	s.Server = grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.ConnectionTimeout(time.Duration(cfg.ConnectionTimeout)*time.Second),
		grpc.NumStreamWorkers(uint32(cfg.NumStreamWorkers)),
		grpc.MaxMsgSize(int(cfg.MaxMsgSize)),
	)

	s.WrappedGrpcServer = grpcweb.WrapServer(s.Server,
		grpcweb.WithOriginFunc(func(origin string) bool {
			if _, ok := cfg.Origin[origin]; !ok {
				return false
			}
			return true
		}),
	)

	s.Register()

	return nil
}

func (s *Service) Close() error {
	s.Server.GracefulStop()
	return nil
}
