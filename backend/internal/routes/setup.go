package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/ultimateshadsform/pi-dashboard/internal/database"
	"github.com/ultimateshadsform/pi-dashboard/internal/helpers"
	"github.com/ultimateshadsform/pi-dashboard/internal/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateAdminUserHandler(c fiber.Ctx) error {
	// Decode body json
	var u structs.RequestSetupUser
	if err := json.Unmarshal(c.Body(), &u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Check if values are empty
	if u.Username == "" || u.Password == "" || u.ConfirmPassword == "" || u.SetupCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username, password, confirm password and setup code cannot be empty",
		})
	}

	// Check if password and confirm password match
	if u.Password != u.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Password and confirm password do not match",
		})
	}

	// Check if setup is already complete
	db := database.GetDB()

	// Check if setup code is correct and matches the one in the database
	var dbsettings structs.DBSettings

	if err := db.Collection("settings").FindOne(c.Context(), bson.D{}).Decode(&dbsettings); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cannot get settings",
			})
		}
	}

	if dbsettings.IsSetupComplete {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Setup is already complete",
		})
	}

	if u.SetupCode != dbsettings.SetupCode {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Setup code is incorrect",
		})
	}

	// Create admin user and hash password
	hashedPassword := helpers.HashPassword(u.Password)

	// Insert user into database
	user := structs.DBUser{
		Username: u.Username,
		Password: hashedPassword,
	}

	if _, err := db.Collection("users").InsertOne(c.Context(), user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create user",
		})
	}

	// Update settings
	dbsettings.IsSetupComplete = true
	if _, err := db.Collection("settings").UpdateOne(c.Context(), bson.D{}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "is_setup_complete", Value: true},
		}},
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update settings",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
