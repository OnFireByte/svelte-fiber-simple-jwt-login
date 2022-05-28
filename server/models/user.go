package models

type User struct {
	OwnModel
	Uuid     string `json:"uuid" gorm:"unique;not null;index"`
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password []byte `json:"-" gorm:"not null"`
}