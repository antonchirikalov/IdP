package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"html/template"
	. "idp/controllers"
	"idp/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	httpPort := utils.GetConfig().GetString("httpport")
	router := gin.Default()
	router.HTMLRender = loadTemplates()

	loginController := new(LoginController)
	consentController := new(ConsentController)
	callbackController := new(CallbackController)
	tokenController := new(TokenController)
	registerController := new(RegisterController)
	recoveryController := new(RecoveryController)
	resetController := new(ResetController)
	confirmationController := new(ConfirmationController)
	messageController := new(MessageController)
	apiController := new(ApiController)


	router.GET("/login", loginController.Get)
	router.POST("/login", loginController.Post)
	router.GET("/consent", consentController.Get)
	router.POST("/consent", consentController.Post)
	router.GET("/callback", callbackController.Get)
	router.POST("/callback", callbackController.Post)
	router.GET("/token", tokenController.Get)
	router.POST("/token", tokenController.Post)
	router.GET("/register", registerController.Get)
	router.POST("/register", registerController.Post)

	router.GET("/recover", recoveryController.Get)
	router.POST("/recover", recoveryController.Post)

	router.GET("/reset", resetController.Get)
	router.POST("/reset", resetController.Post)
	router.GET("/confirmation", confirmationController.Get)
	router.GET("/message", messageController.Get)

	//api
	router.GET("api/users/email/:email", apiController.GetUserByEmail)


	router.Static("/assets", "./views/assets")

	errs := make(chan error, 2)
	go func() {
		log.WithFields(log.Fields{"transport:": "http",
			"port": httpPort}).Info("HIVEON.API has started")

		errs <- router.Run(":" + httpPort)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Info("terminated", <-errs)

}

func loadTemplates() multitemplate.Render {
	box := packr.New("views", "./views")
	r := multitemplate.New()
	for _, templateName := range box.List(){
		rendered, _ := box.FindString(templateName)
		template, err := template.New(templateName).Parse(rendered)
		if err != nil{
			log.Fatal(err)
		}
		r.Add(templateName, template)
	}
	
	return r
}
