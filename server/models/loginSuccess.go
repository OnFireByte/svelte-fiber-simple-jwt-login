package models

type LoginSuccess struct {
	Status string `json:"status"`
	Jwt    string `json:"jwt"`
}