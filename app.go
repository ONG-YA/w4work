package w4work

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	opts     options
	ctx     context.Context
	cancel  func()
}

func New(opts ...Option) *App {
	options := options{
		ctx:    context.Background(),
		signals:   []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &App{ctx: ctx, cancel: cancel, opts:options}
}

func (app *App) Run() error {
	erp, ctx := errgroup.WithContext(app.ctx)
	for _, srv := range app.opts.servers {
		srv := srv
		erp.Go(func() error {
			<-ctx.Done() // wait for stop signal
			return srv.Stop()
		})
		erp.Go(func() error {
			return srv.Start()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, app.opts.signals...)
	erp.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				app.Stop()
			}
		}
	})
	if err := erp.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (app *App) Stop() error {
	if app.cancel != nil {
		app.cancel()
	}
	return nil
}

