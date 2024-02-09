package main

import(
	"github.com/gofiber/fiber/v3" // fast framework 
	"gorm.io/gorm" // relationship databases
	"https://github.com/joho/godotenv" 
) // required packages

type Repository struct{
	DB *gorm.DB
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBooksById)
	api.Get("/get_books", r.GetBooks)
}

func main(){ // Load ambient variables 
	err := godotenv.Load(".env") //  Load  .env file
	if err != nil {
		log.Fatal(err)
	}

	r := Repository {
		DB: db,	
	}

	app := fiber.New()

	r.SetupRoutes(app)
	app.Listen(":8080") // defining the route
}