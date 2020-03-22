package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"go-micro-demo/4_gin/controller"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-micro-demo/4_gin/docs"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	u := controller.User{}
	//router.Handle("GET", "/user ", u.Login)
	v1 := router.Group("/api/v1")
	{
		// http://localhost:8080/api/v1/login
		v1.GET("/login",u.Login)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	service := web.NewService(
		web.Name("gin"),
		web.Address(":8080"),
		web.Handler(router),
	)


	service.Run()
}
