package main

import (
	"fibrecode/student"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/students", student.GetStudents)
	app.Get("/student/:id", student.GetStudent)
	app.Post("/student", student.SaveStudent)
	app.Delete("/student/:id", student.DeleteStudent)
	app.Put("/student/:id", student.UpdateStudent)
}

func main() {
	student.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)

	Routers(app)

	log.Fatal(app.Listen(":3000"))
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Go API")
}
