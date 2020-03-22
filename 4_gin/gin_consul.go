package main

import (

	"github.com/micro/go-micro/web"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := web.NewService(
		web.Name("gin_registry_consul"),
		web.Address(":8080"),
		web.Registry(reg),
	)

	//router.Run(":8080")
	service.Run()

}
