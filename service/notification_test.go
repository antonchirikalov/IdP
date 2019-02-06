package service

import (
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
	. "idp/models"
	"testing"
)

func TestSendNotificationEmail(t *testing.T) {
	Convey("Test users repository connection", t, func() {
		user := User{Username: "TestUser", Email: "achirikalov@gmail.com", Password: "password", Token: uuid.New().String()}

		Convey("Trying to get user", func() {
			SendConfirmationMessage(user)
		})
	})
}
