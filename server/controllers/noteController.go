package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/onfirebyte/simple-jwt-login/db"
	"github.com/onfirebyte/simple-jwt-login/models"
	"github.com/samber/lo"
)

func AddNote(c *fiber.Ctx) error {

	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(400).SendString("Error parsing body")
	}

	user, err := AuthCheck(c)
	if err != nil {
		switch err.Error() {
		case "No JWT cookie":
			return c.Status(403).JSON(fiber.Map{
						"status" : "no_jwt_cookie",
						"message": "No JWT cookie, please login",
					})
		case "signature is invalid":
			return c.Status(403).JSON(fiber.Map{
						"status" : "login_unauthenticated",
						"message": "signature is invalid",
					})

		}
	}

	note_uuid := uuid.New().String()
	note.Uuid = note_uuid
	note.Owner_id = user.ID
	note.Status = models.NoteStatusActive

	dbErr := db.DB.Create(&note).Error
	if dbErr != nil {
		return c.Status(405).JSON(fiber.Map{
			"status" : "db_error",
			"message": "Something went wrong",
		})
	}

	return c.Status(200).JSON(note)
	

}

func SeeNote(c *fiber.Ctx) error{
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
	notes := []models.Note{}
	dbErr := db.DB.Where("owner_id = ?", user.ID).Find(&notes).Order("create_at").Error
	if dbErr != nil {
		return c.JSON(fiber.Map{
			"status" : "note_error",
			"message": err.Error(),
		})
	}
	
	isGroup := bool(c.Query("group") == "true")
	if !isGroup {
		return c.JSON(notes)
	}
	
	groups := lo.GroupBy(notes,func(note models.Note) string {
		return string(note.Status)
	})
	return c.JSON(groups)
}

