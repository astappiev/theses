package topic

import (
	"github.com/gofiber/fiber/v2"
	"theses/page"

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

func CreateTopic() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("create", fiber.Map{
			"Title": "Create topic",
		}, "layouts/main")
	}
}

func CreateTopicPost() fiber.Handler {
	form := new(Topic)

	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(form); err != nil {
			c.Status(fiber.StatusBadRequest)
			page.AddError(c, "Bad request")
			return CreateTopic()(c)
		}

		db.Con().Save(&form)

		return c.Redirect("/")
	}
}
