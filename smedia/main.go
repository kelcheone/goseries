package main

import (
	"fmt"
	"log"
	"net/http"
	"smedia/database"
	"smedia/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Repo struct {
	DB *gorm.DB
}

func main() {
	app := fiber.New()
	db, err := database.Init()
	if err != nil {
		log.Fatal("could not load the database")
	}

	r := Repo{
		DB: db,
	}
	r.UserRoutes(app)
	fmt.Println("Listening in port 8080")
	app.Listen(":8000")
}

func (R *Repo) UserRoutes(app *fiber.App) {
	api := app.Group("/users")
	api.Post("/", R.CreateUser)
	api.Get("/", R.GetUsers)
	api.Get("/:id", R.GetUserByID)
	api.Put("/:id", UpdateUser)
	api.Delete("/:id", DeleteUser)
}

func (R *Repo) CreateUser(c *fiber.Ctx) error {
	userModel := User{}
	err := c.BodyParser(&userModel)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "could not parse user",
			},
		)
	}

	err = R.DB.Create(&userModel).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "could not create user ",
			},
		)
	}

	return c.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "User created",
		},
	)

}
func (R *Repo) GetUsers(c *fiber.Ctx) error {
	userModel := &[]models.User{}

	err := R.DB.Find(&userModel).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "could not get users",
			},
		)
	}
	return c.Status(http.StatusOK).JSON(&userModel)

}
func (R *Repo) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "id can't be null ",
			},
		)
	}

	user := &models.User{}
	err := R.DB.First(&user, "id=?", id)
	if err.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "user not found",
			},
		)
	}
	return c.Status(http.StatusOK).JSON(&user)
}
func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Updating user")
}
func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Deleting user")
}
