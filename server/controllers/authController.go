package controllers

import (
	"fmt"
	"net/mail"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/onfirebyte/simple-jwt-login/db"
	"github.com/onfirebyte/simple-jwt-login/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var regData models.RegisterData
	if err := c.BodyParser(&regData); err != nil {
		return c.Status(400).SendString("Error parsing body")
	}

	if _,err := mail.ParseAddress(regData.Email); err != nil {
		errRes := models.ErrorResponse{
			Status:      models.ErrorInvalidEmail,
			Description: "Invalid Email",
		}
		return c.Status(401).JSON(errRes)
	}

	if regData.Password != regData.ConfirmPassword {
		errRes := models.ErrorResponse{
			Status:      models.ErrorPasswordNotMatch,
			Description: "Password and Confirm Password not match",
		}
		return c.Status(401).JSON(errRes)
	}

	
	new_uuid := uuid.New().String()
	password, _ := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)

	user := models.User{
		Uuid:       new_uuid,
		Username: regData.Username,
		Email:    regData.Email,
		Password: password,
	}
	// if err := db.Create(&registerData).Error; err != nil {
	// 	return c.Status(400).SendString("Error creating user")
	// }

	if err := db.DB.Create(&user).Error; err != nil {
		errText := err.Error()
		switch errText{
		case "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)":
			errRes := models.ErrorResponse{
				Status:models.ErrorUsernameExists,
				Description:"Username already exists",
			}
			return c.Status(400).JSON(errRes)
		case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
			errRes := models.ErrorResponse{
				Status:models.ErrorEmailExists,
				Description:"Username already exists",
			}
			return c.Status(400).JSON(errRes)
		}
		return c.Status(400).SendString("Unknown error")
	}

	return c.Status(200).JSON(user)
}

func Login(c *fiber.Ctx) error{

	var loginData models.LoginData
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).SendString("Error parsing body")
	}

	var user models.User
	if err := db.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		errRes := models.ErrorResponse{
			Status:      models.ErrorUsernameNotFound,
			Description: "Username not found",
		}
		return c.Status(404).JSON(errRes)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		errRes := models.ErrorResponse{
			Status:      models.ErrorPasswordNotMatch,
			Description: "Password not match",
		}
		return c.Status(400).JSON(errRes)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: user.Uuid,
		ExpiresAt: time.Now().Add( 30*24* time.Hour ).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		fmt.Println("Error generating token for", user.Uuid,)
		return c.Status(500).SendString("Error generating token")
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add( 30*24* time.Hour ),
		SameSite: "None",
		Secure: true,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"status" : "login_success",
		"message" : "Successfully login",

	})
}

func User(c *fiber.Ctx) error{

	user, err := AuthCheck(c)
	if err != nil {
		switch err.Error() {
		case "No JWT cookie":
			return c.JSON(fiber.Map{
						"status" : "no_jwt_cookie",
						"message": "No JWT cookie, please login",
					})
		case "signature is invalid":
			return c.JSON(fiber.Map{
						"status" : "login_unauthenticated",
						"message": "signature is invalid",
					})

		}
	}

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error{
	// Set jwt cookie to blank and set expires to past 1 hour

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status" : "logout_success",
		"message": "Successfully logout",
	})
}