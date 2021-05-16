package main

import (
	"w4work"
	"w4work/internal/biz"
	"w4work/internal/conf"
	"w4work/internal/data"
	"w4work/internal/server"
	"w4work/internal/service"
	"w4work/transport/http"
)

func newApp(server *http.Server) *w4work.App{
	return w4work.New(w4work.WithServer(server))
}

func main() {
	redisConf := &conf.RedisConf{Addr:"127.0.0.1:6379",DB:0}
	serverConf := &conf.HttpConf{Addr:":8080",Network:"tcp",}
	app,err := initApp(serverConf,redisConf)
	if err != nil {
		panic(err)
	}
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}