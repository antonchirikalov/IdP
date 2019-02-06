package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	. "idp/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	Convey("Test users repository connection", t, func() {
		user := User{Username: "TestUser", Email: "testuser@gmail.com", Password: "password"}
		err := RegisterUser(user)
		Convey("Trying to get user", func() {
			So(err, ShouldBeNil)
			user, userNotExist := CheckActiveUser("testuser@gmail.com")
			So(userNotExist, ShouldEqual, false)
			So(user.Username, ShouldEqual, "TestUser")
			So(user.Password, ShouldEqual, "password")
			DeleteUser("testuser@gmail.com")
		})
	})
}
