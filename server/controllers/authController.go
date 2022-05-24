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

	
	id := uuid.New().String()
	password, _ := bcrypt.GenerateFromPassword([]byte(regData.Password), bcrypt.DefaultCost)

	user := models.User{
		Id:       id,
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
		c.Status(fiber.StatusNotFound)
		return c.Status(401).JSON(errRes)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		errRes := models.ErrorResponse{
			Status:      models.ErrorPasswordNotMatch,
			Description: "Password not match",
		}
		c.Status(fiber.StatusBadRequest)
		return c.Status(401).JSON(errRes)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: user.Id,
		ExpiresAt: time.Now().Add( 30*24* time.Hour ).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		fmt.Println("Error generating token for", user.Id,)
		return c.Status(400).SendString("Error generating token")
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add( 30*24* time.Hour ),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"status" : "login_success",
		"message" : "Successfully login",

	})
}

func User(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")
	if cookie == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status" : "no_jwt_cookie",
			"message": "No JWT cookie, please login",
		})
	}
	token, err := jwt.ParseWithClaims(cookie,&jwt.StandardClaims{},func(token *jwt.Token) (interface{},error){
		return []byte(os.Getenv("JWT_SIGNING_KEY")),nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status" : "login_unauthenticated",
			"message": "signature is invalid",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User

	db.DB.Where("id = ?", claims.Issuer).First(&user)


	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error{
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