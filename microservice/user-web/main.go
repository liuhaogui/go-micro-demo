package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"microservice/user-web/base"
	"microservice/user-web/base/config"
	"microservice/user-web/handler"
)

func main() {
	// config init / db init
	base.Init()

	// etcd option
	etcdReg := etcd.NewRegistry(registryOptions)

	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.user"),
		web.Registry(etcdReg),
		web.Version("latest"),
		web.Address(":8092"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			// 初始化handler
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	// registe handler
	service.HandleFunc("/user/login", handler.Login)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
