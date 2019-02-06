package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	. "idp/models"
	. "idp/utils"
	"strconv"
)

func SendConfirmationMessage(user User) {

	message := gomail.NewMessage()
	message.SetHeader("To", user.Email)
	message.SetHeader("Subject", "Confirmation email")
	confirmationURL := "http://" + GetConfig().GetString("httpaddr") + ":" + GetConfig().GetString("httpport") +
		fmt.Sprintf("/confirmation?token=%s", user.Token)
	messageBody := fmt.Sprintf("Hello <b>%s</b>!<h1>Please click this link to confirm your account</h1>"+
		"<p>%s</p>", user.Username, confirmationURL)
	message.SetBody("text/html", messageBody)

	sendMessage(message)

}


func SendResetMessage(user User) {

	message := gomail.NewMessage()
	message.SetHeader("To", user.Email)
	message.SetHeader("Subject", "Reset password recovery email")
	confirmationURL := "http://" + GetConfig().GetString("httpaddr") + ":" + GetConfig().GetString("httpport") +
		fmt.Sprintf("/reset?token=%s", user.Token)
	messageBody := fmt.Sprintf("Hello <b>%s</b>!<h1>Please click this link to recover your password</h1>"+
		"<p>%s</p>", user.Username, confirmationURL)
	message.SetBody("text/html", messageBody)

	sendMessage(message)

}



func sendMessage(message *gomail.Message){
	active := GetConfig().GetBool("notification.active")
	message.SetHeader("From", GetConfig().GetString("notification.from"))
	if active {
		smtp := GetConfig().GetString("notification.smtp")
		port, _ := strconv.Atoi(GetConfig().GetString("notification.port"))
		username := GetConfig().GetString("notification.user")
		password := GetConfig().GetString("notification.password")

		d := gomail.NewDialer(smtp, port, username, password)

		if err := d.DialAndSend(message); err != nil {
			log.Error(err)
		}
	}
	log.Info("Message sent: ", message)
}
