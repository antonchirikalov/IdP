package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type CallbackController struct {
}

func (ctrl CallbackController) Get(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	scope := c.Request.URL.Query().Get("scope")
	state := c.Request.URL.Query().Get("state")
	token, err := oauthClient.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Error(err)
		c.HTML(500, "callback.html", nil)
	}
	c.HTML(200, "callback.html", gin.H{"code": code,
		"scope":   scope,
		"state":   state,
		"token":   token.AccessToken,
		"refresh": token.RefreshToken})
}

func (ctrl CallbackController) Post(c *gin.Context) {
	log.Info(c.Request.PostForm)
}
