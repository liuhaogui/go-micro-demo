package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"microservice/user-srv/handler"
	"microservice/user-srv/model"

	//"github.com/micro/examples/kubernetes/go/micro"
	//"github.com/micro/go-micro"
	log "github.com/micro/go-micro/v2/logger"
	"microservice/user-srv/base/config"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	//user "microservice/user-srv/proto/user"
	"microservice/user-srv/base"
	proto "microservice/user-srv/proto/user"
)

func main() {
	// config init / db init
	base.Init()

	// etcd option
	etcdReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Version("latest"),
		micro.Registry(etcdReg),
		micro.Address(":8090"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) error {
			model.Init()
			handler.Init()
			return nil
		}),
	)

	// resgister handler
	proto.RegisterUserHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
