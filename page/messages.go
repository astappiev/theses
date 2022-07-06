package page

import (
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Severity string `json:"severity"`
	Message  string `json:"text"`
}

func AddMessage(c *fiber.Ctx, message Message) {
	var msgs []Message
	local := c.Locals("Messages")
	if local != nil {
		msgs = local.([]Message)
	}

	msgs = append(msgs, message)
	c.Locals("Messages", msgs)
}

func AddError(c *fiber.Ctx, message string) {
	AddMessage(c, Message{Severity: "danger", Message: message})
}

func AddSuccess(c *fiber.Ctx, message string) {
	AddMessage(c, Message{Severity: "success", Message: message})
}
