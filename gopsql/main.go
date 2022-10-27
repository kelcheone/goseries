package main

import (
	"log"

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