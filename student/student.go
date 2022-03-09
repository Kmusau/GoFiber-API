package student

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Musau6565.@tcp(127.0.0.1:3306)/gostudentsdb?charset=utf8mb4&parseTime=True&loc=Local"

type Student struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Course    string `json:"course"`
}

//function to connect to the database
func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the Database")
	}
	DB.AutoMigrate(&Student{})
}

//function to find all the students
func GetStudents(c *fiber.Ctx) error {
	var students []Student
	DB.Find(&students)
	return c.JSON(&students)
}

//function to return a single student given its id
func GetStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student Student
	DB.Find(&student, id)
	return c.JSON(&student)
}

//function to create a new student
func SaveStudent(c *fiber.Ctx) error {
	student := new(Student)
	if err := c.BodyParser(student); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Create(&student)
	return c.JSON(&student)
}

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student Student
	DB.First(&student, id)

	if student.Course == "" {
		return c.Status(500).SendString("Student not available")
	}
	DB.Delete(&student)
	return c.SendString("Student deleted successfully")
}

func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")
	var student Student
	DB.First(&student, id)

	if student.Course == "" {
		return c.Status(500).SendString("Student not available")
	}
	if err := c.BodyParser(student); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&student)
	return c.JSON(&student)
}
