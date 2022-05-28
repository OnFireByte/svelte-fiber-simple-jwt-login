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

func UpdateNote(c *fiber.Ctx) error {
	User, err := AuthCheck(c)
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

	var data map[string]string
	dataErr := c.BodyParser(&data)
	if dataErr != nil {
		return c.JSON(fiber.Map{
			"status" : "data_error",
			"message": dataErr.Error(),
		})
	}

	note_uuid,ok := data["uuid"]
	if !ok{
		return c.Status(401).JSON(fiber.Map{
			"status" : "no_uuid",
			"message": "note uuid field is empty",
		})
		
	}
	note := models.Note{
		Uuid: note_uuid,
	}

	if content, ok := data["content"]; ok {
		note.Content = content
	}

	if status, ok := data["status"]; ok {
		note.Status = models.NoteStatus(status)
	}

	dbData := db.DB.Model(&note).Where("owner_id = ?", User.ID).Update("content","status")
	if dbData.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status" : "update_error",
			"message": dbData.Error.Error(),
		})
	}
	
	if dbData.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status" : "note_not_found",
			"message": "You don't have this note",
		})
	}

	return c.JSON(fiber.Map{
		"status" : "success",
		"message": "note updated",
	})
}

func DeleteNote(c *fiber.Ctx) error {
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

	var data map[string]string
	dataErr := c.BodyParser(&data)
	if dataErr != nil {
		return c.JSON(fiber.Map{
			"status" : "data_error",
			"message": dataErr.Error(),
		})
	}

	note_uuid,ok := data["uuid"]
	if !ok{
		return c.Status(401).JSON(fiber.Map{
			"status" : "no_uuid",
			"message": "note uuid field is empty",
		})
		
	}

	dbData := db.DB.Where("uuid = ?", note_uuid).Where("owner_id = ?",user.ID).Delete(&models.Note{})
	if dbData.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status" : "delete_error",
			"message": dbData.Error.Error(),
		})
	}
	if dbData.RowsAffected == 0{
		return c.Status(404).JSON(fiber.Map{
			"status" : "note_not_found",
			"message": "You don't have this note",
		})
	}
	return c.JSON(fiber.Map{
		"status" : "success",
		"message": "note deleted",
	})

}