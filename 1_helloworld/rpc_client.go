package main

import (
	"context"
	"fmt"
	pb "go-micro-demo/1_helloworld/proto"
	"github.com/micro/go-micro/v2"
)
func main()  {

	service := micro.NewService(
		micro.Name("greeter.clent"),
	)
	service.Init()

	greeter := pb.NewGreeterService("rpc.server", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &pb.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)

}
