package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/ory/hydra/sdk/go/hydra/swagger"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"gopkg.in/resty.v1"
	. "idp/utils"
)

type TokenController struct{}

var oauthClient *oauth2.Config

func (ctrl TokenController) Get(c *gin.Context) {
	c.HTML(200, "token.html", gin.H{"clients": GetClients()})
}

func (ctrl TokenController) Post(c *gin.Context) {
	if c.PostForm("submit") == "Authorize" {
		clientId := c.PostForm("clientName")
		secret := c.PostForm("secret")
		oauthClient = InitClient(clientId, secret, c.Request.Host)
		redirectUrl := oauthClient.AuthCodeURL("state123")
		log.Info("Redirect to ", redirectUrl)
		c.Redirect(302, redirectUrl)
	} else {

		newClient := OAuth2Client{ClientId: c.PostForm("newClientId"), ClientSecret: c.PostForm("newSecret"),
			RedirectUris: []string{c.PostForm("callbackURL")}, Scope: "openid offline", GrantTypes: []string{"authorization_code", "refresh_token"},
			ResponseTypes: []string{"code", "id_token"}}
		hydraAPI := GetConfig().GetString("hydraAdmin") + "/clients"
		jsonValue, _ := json.Marshal(newClient)
		res, err := resty.R().SetHeader("Content-Type", "application/json").SetHeader("Accept", "application/json").
			SetBody(jsonValue).Post(hydraAPI)
		if err != nil {
			log.Error(err)
		}
		log.Info(res)
		c.Redirect(302, "/token")

	}

}

func GetClients() []string {
	hydraAPI := GetConfig().GetString("hydraAdmin")
	res, err := resty.R().Get(hydraAPI + "/clients")
	if err != nil {
		log.Error(err)
	}

	var clients []OAuth2Client
	var ret []string
	json.Unmarshal(res.Body(), &clients)
	for _, c := range clients {
		ret = append(ret, c.ClientId)
	}

	return ret
}

func GetClient(clientId string) OAuth2Client {
	clientUrl := GetConfig().GetString("hydraAdmin") + "/clients/" + clientId
	res, err := resty.R().Get(clientUrl)
	if err != nil {
		log.Info(err)
	}
	var client OAuth2Client
	json.Unmarshal(res.Body(), &client)
	return client

}

func InitClient(clientId string, secret string, host string) *oauth2.Config {

	hydraAPI := GetConfig().GetString("hydraAPI")
	client := GetClient(clientId)
	oauthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  hydraAPI + "/oauth2/auth",
			TokenURL: hydraAPI + "/oauth2/token",
		},
		RedirectURL: client.RedirectUris[0],
		Scopes:      getScopes(),
	}

	return oauthConfig
}

func getScopes() []string {
	return []string{"openid", "offline"}
}
