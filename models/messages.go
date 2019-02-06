package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Token string
	Active bool `gorm:"not null"`
}

type LoginRequest struct {
	Subject string `json:"subject"`
	Remember bool `json:"remember"`
	RememberFor int `json:"remember_for"`
}

type RejectConsentRequest struct {
	Error string `json:"error"`
	ErrorDescription string `json:"error_description"`
}