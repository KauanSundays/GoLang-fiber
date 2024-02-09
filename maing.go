package main

import(
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"https://github.com/joho/godotenv"
)

type Reposutory struct{
	DB
}

func main(){
	err := godotenv.Load(".env") //  Load  .env file
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	r.SetupRoutes(app)
	app.Listen(":8080") // defining the route
}