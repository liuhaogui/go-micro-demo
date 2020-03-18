package main

import (
	"context"
	"log"

	pb "github.com/liuhaogui/go-micro-demo/1_helloworld/proto"
	"github.com/micro/go-micro/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}


func main() {
	service := micro.NewService(
		micro.Name("rpc.server"),
	)

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), &Greeter{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
