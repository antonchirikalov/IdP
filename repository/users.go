package repository

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	. "idp/models"
	. "idp/utils"
)

var db *gorm.DB

func init() {

	conn := fmt.Sprintf("host=%s  port=%s sslmode=disable user=%s dbname=%s password=%s",
		GetConfig().GetString("postgres.host"),
		GetConfig().GetString("postgres.port"),
		GetConfig().GetString("postgres.user"),
		GetConfig().GetString("postgres.dbName"),
		GetConfig().GetString("postgres.password"),
	)
	var err error
	db, err = gorm.Open("postgres", conn)
	db.LogMode(true)

	if err != nil {
		log.Fatal(err)
	}
}

func CheckUserAndPassword(email string, password string) (User, bool) {
	var user User
	userNotExists := db.Where("email = ? AND password = ? AND active = ?", email, password, true).
		First(&user).RecordNotFound()
	return user, userNotExists

}

func CheckActiveUser(email string) (User, bool) {
	var user User
	userNotExists := db.Where("email = ? AND active = ?", email, true).
		First(&user).RecordNotFound()
	return user, userNotExists
}

func RegisterUser(user User) error {
	db.NewRecord(user)
	return db.Create(&user).Error
}

func DeleteUser(email string) {
	db.Unscoped().Where("email LIKE ?", email).Delete(User{})
	db.Commit()
}

func UpdateUserToken(email string) (User, error) {
	var user User
	userNotExists := db.Where("email = ? AND active = ?", email, true).First(&user).RecordNotFound()
	if userNotExists {
		return user, errors.New("User doesn't exists!")
	}
	db.Model(&user).Update("token", uuid.New().String())
	return user, nil
}

func UpdateUserPasswordByToken(token string, password string) (User, error) {
	var user User
	userNotExists := db.Where("token = ?", token).First(&user).RecordNotFound()
	if userNotExists {
		return user, errors.New("Token doesn't exists!")
	}
	db.Model(&user).Update("password", password)
	return user, nil
}

func ActivateUser(token string) (User, bool) {
	var user User
	userNotExists := db.Where("token = ?", token).First(&user).RecordNotFound()

	if !userNotExists {
		db.Model(&user).Update("active", true)
	}
	return user, userNotExists
}


func GetUserByEmail(email string) (User, bool){
	var user User
	userNotExists := db.Where("email = ?", email).First(&user).RecordNotFound()
	return user, userNotExists
}
