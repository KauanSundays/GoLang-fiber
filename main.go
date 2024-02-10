package main

import(
	"github.com/gofiber/fiber/v3" // fast framework 
	"gorm.io/gorm" // relationship databases
	"https://github.com/joho/godotenv" 
) // required packages

type Book struct{
	Author  string `json:"author"`
	Title    string `json:"title"`
	Publisher  string `json:"publisher"`
}

type Repository struct{
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {

	book := Book{}  //stores the data

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(Http.StatusUnprocessableEntity).JSON(
			$fiber.Map{("message": "request failed"})
		return err
	}

	if err!= nil {
		context.Status(http.StatusBadRequest).JSON( //erro 400
			&fiber.Map{"error":"could not create book"})
			return err
		)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{ // status 200
		"message": "Top, it's added"
	}) 
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx)  error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Couldn't get books"})
			return err
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

	db, err := storage.NewConnection(config)

	if err != mil {
		log.Fatal('could not load the database')
	}

	r := Repository {
		DB: db,	
	}

	app := fiber.New()

	r.SetupRoutes(app)
	app.Listen(":8080") // defining the route
}