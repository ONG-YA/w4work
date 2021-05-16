package w4work

import (
	"context"
	"os"
	"w4work/transport"
)

// Option is an application option.
type Option func(o *options)

// options is an application options.
type options struct {
	ctx  context.Context
	signals []os.Signal

	servers []transport.Server
}

// WithContext with service context.
func WithContext(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// WithSignal with exit signals.
func WithSignal(signals ...os.Signal) Option {
	return func(o *options) { o.signals = signals }
}

// WithServer with transport servers.
func WithServer(srv ...transport.Server) Option {
	return func(o *options) { o.servers = srv }
}
