package main

import (
	"flag"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/example/graceful/etcd/api/config"
	"github.com/tal-tech/go-zero/example/graceful/etcd/api/handler"
	"github.com/tal-tech/go-zero/example/graceful/etcd/api/svc"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/graceful-api.json", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	client := zrpc.MustNewClient(c.Rpc)
	ctx := &svc.ServiceContext{
		Client: client,
	}

	engine := rest.MustNewServer(c.RestConf)
	defer engine.Stop()

	handler.RegisterHandlers(engine, ctx)
	engine.Start()
}
