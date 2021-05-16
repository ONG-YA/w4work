package server

import (
	v1 "w4work/api/user/v1"
	"w4work/internal/conf"
	"w4work/internal/service"
	"w4work/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.HttpConf,user *service.UserService) *http.Server {
	var opts []http.ServerOption
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr))
	}
	srv := http.NewServer(opts...)
	srv.Handler = v1.NewUserServiceHandler(user)
	return srv
}

//return &http.Server{Addr:c.Addr,Handler: v1.NewUserServiceHandler(user)}

