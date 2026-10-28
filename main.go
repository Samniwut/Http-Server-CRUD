package main

import (
	"fmt"
	"os"

	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt"
	"github.com/gofiber/template/html/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf("URL = %s, Method = %s , Time = %s\n",
		c.OriginalURL(), c.Method(), start)
	return c.Next()
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Initialize in-memory data
	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})
	books = append(books, Book{ID: 3, Title: "The Forest", Author: "Steam"})

	app.Post("/login", login)

	app.Use(checkMiddleware)

	// JWTMiddleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload/", uploadfile)
	app.Get("/test-html", testHtml)

	app.Listen(":8080")

}

func testHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "World",
	})
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{
	Email:    "user@example.com",
	Password: "password123",
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}
	return c.JSON(fiber.Map{
		"message": "Login sucess",
	})
}
