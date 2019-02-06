package service

import (
	"encoding/json"
	"fmt"
	. "github.com/ory/hydra/consent"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
	. "idp/models"
	. "idp/utils"
)

func init() {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
}

func GetLoginFlow(challenge string) (AuthenticationRequest, error) {
	url := fmt.Sprintf("%s/oauth2/auth/requests/login/%s", GetConfig().GetString("hydraAdmin"), challenge)
	authResult := AuthenticationRequest{}

	res, err := resty.R().Get(url)
	if err != nil {
		log.Error(err)
		return authResult, err
	}

	if res.StatusCode() < 200 || res.StatusCode() > 302 {
		err = errors.New("An error occurred while making a HTTP request")
		log.Error(err)
		return authResult, err
	}

	json.Unmarshal(res.Body(), &authResult)
	return authResult, nil
}

func AcceptLoginRequest(loginRequest LoginRequest, challenge string) RequestHandlerResponse {
	url := fmt.Sprintf("%s/oauth2/auth/requests/login/%s/accept", GetConfig().GetString("hydraAdmin"), challenge)
	res, err := resty.R().SetBody(loginRequest).
		SetHeader("Content-Type", "application/json").Put(url)
	if err != nil {
		log.Error(err)
	}

	authResponse := RequestHandlerResponse{}
	json.Unmarshal(res.Body(), &authResponse)

	return authResponse
}

func GetConsentFlow(challenge string) (ConsentRequest, error) {
	url := fmt.Sprintf("%s/oauth2/auth/requests/consent/%s", GetConfig().GetString("hydraAdmin"), challenge)
	consentResult := ConsentRequest{}
	res, err := resty.R().Get(url)
	if err != nil {
		log.Error(err)
		return consentResult, err
	}

	if res.StatusCode() < 200 || res.StatusCode() > 302 {
		err = errors.New("An error occurred while making a HTTP request")
		log.Error(err)
		return consentResult, err
	}

	json.Unmarshal(res.Body(), &consentResult)
	return consentResult, nil
}

func PutConsentReject(challenge string) (RequestHandlerResponse, error){
	url := fmt.Sprintf("%s/oauth2/auth/requests/consent/%s/reject", GetConfig().GetString("hydraAdmin"), challenge)
	rejectRequest := RejectConsentRequest{Error:"access_denied", ErrorDescription:"The resource owner denied the request"}
	consentResponse := RequestHandlerResponse{}
	res, err := resty.R().SetBody(rejectRequest).
		SetHeader("Content-Type", "application/json").Put(url)
	if err != nil {
		log.Error(err)
		return consentResponse, err
	}
	json.Unmarshal(res.Body(), &consentResponse)
	return consentResponse, nil
}

func AcceptConsentRequest(acceptRequest HandledConsentRequest, challenge string) (RequestHandlerResponse, error){
	url := fmt.Sprintf("%s/oauth2/auth/requests/consent/%s/accept", GetConfig().GetString("hydraAdmin"), challenge)
	consentResponse := RequestHandlerResponse{}
	res, err := resty.R().SetBody(acceptRequest).
		SetHeader("Content-Type", "application/json").Put(url)

	if err != nil {
		log.Error(err)
		return consentResponse, err
	}
	json.Unmarshal(res.Body(), &consentResponse)
	return consentResponse, nil
}
