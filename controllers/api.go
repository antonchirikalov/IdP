package controllers

import (
	"github.com/gin-gonic/gin"
	"idp/service"
)

type ApiController struct {}

func (ctrl ApiController) GetUserByEmail(c *gin.Context){
	email := c.Param("email")
	user, err := service.GetUserByEmail(email)
	if err != nil{
		c.JSON(204, err.Error())
		return
	}
	c.JSON(200, user)
}
