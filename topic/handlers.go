package topic

import (
	"github.com/gofiber/fiber/v2"

	"theses/db"
)

func ListTopics() fiber.Handler {
	var topics []Topic
	db.Con().Preload("Users").Find(&topics)

	return func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":  "Topics",
			"Topics": topics,
		}, "layouts/main")
	}
}
