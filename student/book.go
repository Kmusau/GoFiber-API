package student

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func AddBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Create(&book)
	return c.JSON(&book)
}

func GetAllBooks(c *fiber.Ctx) error {
	var books []Book
	DB.Find(&books)
	return c.JSON(&books)
}

func GetABook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	DB.Find(&book, id)
	return c.JSON(&book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	DB.First(&book, id)

	if book.Title == "" {
		return c.Status(500).SendString("The book is not available")
	}
	DB.Delete(&book)
	return c.SendString("Book deleted successfully")
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(Book)
	DB.First(&book, id)

	if book.Title == "" {
		return c.Status(500).SendString("Book is not available")
	}
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&book)
	return c.JSON(&book)
}
