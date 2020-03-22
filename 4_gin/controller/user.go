package controller

import (
	"github.com/gin-gonic/gin"
)

type User struct {
}

// Login godoc
// @Summary Log in to the service
// @Description Log in to the service
// @Router /api/v1/login/:username [get]
func (m *User) Login(ctx *gin.Context) {
	ctx.JSON(200, "hello world")
}
