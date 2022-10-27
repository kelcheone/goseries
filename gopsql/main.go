package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) createBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"Message": "request failed"},
		)
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not create book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book has been created"},
	)

	return nil
}

func (r *Repository) deleteBook(context *fiber.Ctx) error  { return nil }
func (r *Repository) getBookById(context *fiber.Ctx) error { return nil }
func (r *Repository) getBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not find books"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "books fetched sucessfully",
			"data":    bookModels,
		},
	)

	return nil
}
func (r *Repository) updateBook(context *fiber.Ctx) error { return nil }

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-post", r.createBook)
	api.Delete("/delete-book/:id", r.deleteBook)
	api.Get("/get-books/:id", r.getBookById)
	api.Get("/get-books", r.getBooks)
	api.Put("/get-boooks/:id", r.updateBook)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()

	r.SetupRoutes(app)
	app.Listen("8080 ")
}
