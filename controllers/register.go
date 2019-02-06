package controllers

import (
	"fmt"
	"github.com/dmnlk/stringUtils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	. "idp/models"
	"idp/repository"
	"idp/service"
)

type RegisterController struct{}

func (ctrl RegisterController) Get(c *gin.Context) {
	challenge := c.Request.URL.Query().Get("login_challenge")
	c.HTML(200, "register.html", gin.H{"challenge": challenge})
}

func (ctrl RegisterController) Post(c *gin.Context) {

	challenge := c.PostForm("challenge")

	_, userNotExists := repository.CheckActiveUser(c.PostForm("email"))
	if !userNotExists {
		c.HTML(200, "register.html", gin.H{"challenge": challenge, "userExists": true})
		return
	}

	user := User{Username: c.PostForm("username"), Email: c.PostForm("email"),
		Password: c.PostForm("password"), Token: uuid.New().String()}
	if err := service.RegisterUser(user); err != nil {
		c.HTML(200, "message.html", gin.H{"message": err.Error()})
	}

	var redirectUrl string

	if stringUtils.IsNoneEmpty(challenge) {
		redirectUrl = fmt.Sprintf("/login?login_challenge=%s", challenge)
	} else {
		c.HTML(200, "message.html", gin.H{"message": "Registration completed succesfully. " +
			"Activation email will be sent shortly"})
	}

	c.Redirect(302, redirectUrl)
}
