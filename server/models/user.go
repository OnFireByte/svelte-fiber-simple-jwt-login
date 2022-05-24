package models

type User struct {
	Id       string `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password []byte `json:"-" gorm:"not null"`
}