package http

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"time"
)

// Server is a HTTP server wrapper.
type Server struct {
	*http.Server
	lis     net.Listener
	network string
	address string
	timeout time.Duration
}

// ServerOption is HTTP server option.
type ServerOption func(*Server)

func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// Address with server address.
func Address(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
	}
	for _, o := range opts {
		o(srv)
	}
	srv.Server = &http.Server{Handler: srv}
	return srv
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), s.timeout)
	defer cancel()
	s.ServeHTTP(res, req.WithContext(ctx))
}


// Start start the HTTP server.
func (s *Server) Start() error {
	fmt.Println("start")
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	if err := s.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop the HTTP server.
func (s *Server) Stop() error {
	fmt.Println("stop")
	return s.Shutdown(context.Background())
}