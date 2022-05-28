package controllers

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/onfirebyte/simple-jwt-login/db"
	"github.com/onfirebyte/simple-jwt-login/models"
)

func AuthCheck(c *fiber.Ctx) (models.User,error) {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		c.Status(fiber.StatusUnauthorized)
		return models.User{},errors.New("No JWT cookie")
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return models.User{},errors.New("signature is invalid")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User

	db.DB.Where("uuid = ?", claims.Issuer).First(&user)

	return user, nil
}