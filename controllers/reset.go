package controllers

import (
	"github.com/gin-gonic/gin"
	"idp/service"
)

type ResetController struct{}

func (ctrl ResetController) Get(c *gin.Context) {
    token := c.Request.URL.Query().Get("token")
	c.HTML(200, "reset.html", gin.H{"token":token})
}

func (ctrl ResetController) Post(c *gin.Context) {

	token := c.PostForm("token")
	newPassword := c.PostForm("password")
	service.UpdateUserPassword(token, newPassword)
	c.HTML(200, "message.html", gin.H{"message": "Password is updated"})
}
