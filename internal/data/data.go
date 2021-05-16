package data

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"w4work/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	rd *redis.Client
}

func NewData(conf *conf.RedisConf) (*Data, error) {
	rd := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.DB,
		DialTimeout:  conf.DialTimeout,
		WriteTimeout: conf.WriteTimeout,
		ReadTimeout:  conf.ReadTimeout,
	})
	return &Data{rd: rd}, nil
}
