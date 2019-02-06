package service

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	. "idp/models"
	"idp/repository"
)

func RegisterUser(user User) error {
	user.Password = EncryptPassword(user.Password)
	if err := repository.RegisterUser(user); err != nil {
		return err
	}
	SendConfirmationMessage(user)
	return nil
}

func CheckUserAndSendResetMessage(email string) error {
	user, err := repository.UpdateUserToken(email)
	if err != nil {
		return err
	}
	SendResetMessage(user)
	return nil
}

func UpdateUserPassword(token string, password string) error {
	hashedPass := EncryptPassword(password)
	_, err := repository.UpdateUserPasswordByToken(token, hashedPass)
	return err

}

func EncryptPassword(password string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
	}
	return string(hashedPass)
}
