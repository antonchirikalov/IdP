package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"idp/repository"
)

type ConfirmationController struct {
}

func (ctrl ConfirmationController) Get(c *gin.Context) {
	token := c.Request.URL.Query().Get("token")
	user, ok := repository.ActivateUser(token)
	var message string

	if ok {
		message = "Something went wrong"
	} else {
		message = fmt.Sprintf("Congratulate %s! Your account is active!", user.Username)
	}
	c.HTML(200, "confirmation.html", gin.H{"message": message})
}
