package main

import (
	"gopsql/models"
	"gopsql/storage"
	"log"
	"net/http"
	"os"

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
	bookModel := Book{}

	err := context.BodyParser(&bookModel)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"Message": "request failed"},
		)
		return err
	}

	err = r.DB.Create(&bookModel).Error
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

func (r *Repository) deleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"},
		)
		return nil
	}

	err := r.DB.Delete(&bookModel, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Book deleted"},
	)

	return nil
}
func (r *Repository) getBookById(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := models.Books{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id can't be null"},
		)
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "book id fetched",
			"data":    bookModel,
		},
	)
	return nil
}
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

// func (r *Repository) updateBook(context *fiber.Ctx) error { return nil }

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-post", r.createBook)
	api.Delete("/delete-book/:id", r.deleteBook)
	api.Get("/get-books/:id", r.getBookById)
	api.Get("/get-books", r.getBooks)
	// api.Put("/get-boooks/:id", r.updateBook)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		Db_name:  os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// db, err := storage.NewConnection(config)
	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	// err = models.MigrateBooks(db)
	err = models.MigrateBooks(db)

	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
