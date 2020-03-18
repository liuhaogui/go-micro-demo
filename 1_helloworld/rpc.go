package main

import (
	pb "github.com/liuhaogui/go-micro-demo/1_helloworld/proto"
	"context"
	"github.com/micro/go-micro/v2"
	"log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("hello world "),
	)
	service.Init()
	pb.RegisterGreeterHandler(service.Server(),new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}