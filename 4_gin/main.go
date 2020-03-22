package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Handle("GET", "/user ", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")

	})

	//router.Handle("GET", "/hello/:name", func(c *gin.Context) {
	//	c.String(200, fmt.Sprintf("hello %s ", c.Param("name")))
	//})

	//// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	//router.GET("/user/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	//

	service := web.NewService(
		web.Name("gin"),
		web.Address(":8080"),
		web.Handler(router),
	)

	//router.Run(":8080")
	service.Run()
}
