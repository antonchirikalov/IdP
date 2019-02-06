package controllers

import "github.com/gin-gonic/gin"

type MessageController struct{}

func (ctrl MessageController) Get(c *gin.Context) {
	message := c.Request.URL.Query().Get("msg")
	c.HTML(200, "message.html", gin.H{"message": message})
}
