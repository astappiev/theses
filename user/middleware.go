package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Authorized(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if sess.Get("uid") != nil {
			c.Locals("uid", sess.Get("uid"))
			c.Locals("user", sess.Get("user"))
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).Redirect("/login")
	}
}
