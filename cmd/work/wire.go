// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"w4work"
	"w4work/internal/biz"
	"w4work/internal/conf"
	"w4work/internal/data"
	"w4work/internal/server"
	"w4work/internal/service"
)

// initApp init servers.
func initApp(*conf.HttpConf, *conf.RedisConf) (*w4work.App, error) {
	panic(wire.Build(wire.NewSet(w4work.NewApp),server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet,  ))
}
