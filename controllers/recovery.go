package controllers

import (
	"github.com/gin-gonic/gin"
	"idp/service"
)

type RecoveryController struct{}

func (ctrl RecoveryController) Get(c *gin.Context) {
	c.HTML(200, "recovery.html", nil)

}

func (ctrl RecoveryController) Post(c *gin.Context) {

	err := service.CheckUserAndSendResetMessage(c.PostForm("email"))
	if err != nil {
		c.HTML(200, "recovery.html", gin.H{"emailNotExists": true})
		return
	}
	c.HTML(200, "message.html", gin.H{"message": "Recovery link has been sent to your email"})

}
