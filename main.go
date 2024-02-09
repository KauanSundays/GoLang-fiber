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