package api

import (
	"github.com/gin-gonic/gin"
	"go_http/pkg/utils"
	"go_http/service"
)

func LoginHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		utils.SendError(c, "密码错误")
		return
	}
	rep, err := service.NewLoginService(username, password).DoLogin()
	if err != nil {
		utils.SendError(c, err.Error())
		return
	}
	utils.SendOk(c, rep)
}

func RegisterHandler(c *gin.Context) {
	username := c.Query("username")
	raw, _ := c.Get("password")
	password, ok := raw.(string)
	if !ok {
		utils.SendError(c, "密码错误")
		return
	}
	rep, err := service.NewLoginService(username, password).DoRegister()
	if err != nil {
		utils.SendError(c, err.Error())
		return
	}

	utils.SendOk(c, rep)
}
