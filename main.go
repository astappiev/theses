package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"

	"theses/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// https://github.com/gofiber/template/tree/master/html
	// engine := html.NewFileSystem(http.Dir("./views", ".html"))
	engine := html.New("./views", ".html")
	engine.Reload(true) // Reload the templates on each render

	// This stores all of your app's sessions
	store := session.New()

	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
	})

	app.Use(recover.New())
	app.Use(csrf.New(csrf.Config{
		KeyLookup:  "form:csrf",
		ContextKey: "csrf",
	}))

	app.Get("/", user.RequiredMiddleware(store), func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/login", user.Login())
	app.Post("/login", user.LoginPost(store))

	app.Get("/logout", user.Logout(store))

	log.Fatal(app.Listen(":3000"))
}
