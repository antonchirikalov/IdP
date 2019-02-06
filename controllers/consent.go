package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/ory/hydra/consent"
	log "github.com/sirupsen/logrus"
	"idp/service"
	. "idp/utils"
)

type ConsentController struct {
}

func (ctrl ConsentController) Get(c *gin.Context) {
	challenge := c.Request.URL.Query().Get("consent_challenge")
	consentResponse, err := service.GetConsentFlow(challenge)
	grantedScope := []string{"openid", "offline"}

	if err != nil {
		log.Error(err)
		return
	}
	log.Info(consentResponse)

	consentResponse, err = service.GetConsentFlow(challenge)
	if err != nil {
		log.Error(err)
		return
	}

	acceptRequest := HandledConsentRequest{GrantedScope: grantedScope, GrantedAudience: consentResponse.RequestedAudience,
		Remember: false, RememberFor: GetConfig().GetInt("remember_for")}

	consentRedirect, err := service.AcceptConsentRequest(acceptRequest, challenge)
	if err != nil {
		log.Error(err)
		return
	}

	c.Redirect(302, consentRedirect.RedirectTo)
}

func (ctrl ConsentController) Post(c *gin.Context) {
	challenge := c.PostForm("challenge")
	if c.PostForm("submit") == "Deny access" {
		consentReject, _ := service.PutConsentReject(challenge)
		c.Redirect(302, consentReject.RedirectTo)
	}

	grantedScope := c.PostFormArray("grant_scope")
	var isRemember = false
	if c.PostForm("remember") == "on" {
		isRemember = true
	}
	consentResponse, err := service.GetConsentFlow(challenge)
	if err != nil {
		log.Error(err)
		return
	}

	acceptRequest := HandledConsentRequest{GrantedScope: grantedScope, GrantedAudience: consentResponse.RequestedAudience,
		Remember: isRemember, RememberFor: GetConfig().GetInt("remember_for")}

	consentRedirect, err := service.AcceptConsentRequest(acceptRequest, challenge)
	if err != nil {
		log.Error(err)
		return
	}

	c.Redirect(302, consentRedirect.RedirectTo)
}
