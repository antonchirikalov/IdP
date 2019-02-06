package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	. "idp/models"
	"idp/repository"
	"idp/service"
	. "idp/utils"
)

type LoginController struct{}

func (ctrl LoginController) Get(c *gin.Context) {
	challenge := c.Request.URL.Query().Get("login_challenge")
	res, err := service.GetLoginFlow(challenge)
	if err != nil {
		log.Error(err)
	}
	log.Info(res)
	c.HTML(200, "login.html", gin.H{"challenge": challenge})
}

func (ctrl LoginController) Post(c *gin.Context) {
	email:= c.PostForm("email")
	password := c.PostForm("password")
	user, userNotExists := repository.CheckActiveUser(email)

	if userNotExists {
		c.HTML(200, "login.html", gin.H{"Error": "The email doesn't exists"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.HTML(200, "login.html", gin.H{"Error": "The password is not correct"})
		return
	}

	challenge := c.PostForm("challenge")
	var isRemember = false
	if c.PostForm("remember") == "on" {
		isRemember = true
	}

	loginRequest := LoginRequest{Subject: user.Email, Remember: isRemember, RememberFor: GetConfig().GetInt("remember_for")}

	authResponse := service.AcceptLoginRequest(loginRequest, challenge)
	c.Redirect(302, authResponse.RedirectTo)

}


