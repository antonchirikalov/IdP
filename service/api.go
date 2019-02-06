package service

import (
	"github.com/go-errors/errors"
	. "idp/models"
	"idp/repository"
)

func GetUserByEmail(email string) (user User, err error) {
	user, userNotExists := repository.GetUserByEmail(email)
	if userNotExists {
		return user, errors.New("User does not exist")
	}
	return user, nil
}
