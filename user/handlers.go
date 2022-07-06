package user

import (
	"encoding/gob"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"theses/ldap"
	"theses/page"
)

func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{
			"Title": "Login",
		}, "layouts/main")
	}
}

func LoginPost(store *session.Store) fiber.Handler {
	form := new(LoginForm)

	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(form); err != nil {
			c.Status(fiber.StatusBadRequest)
			page.AddError(c, "Bad request")
			return Login()(c)
		}

		user, err := ldap.FindUser(form.Email, form.Password)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			page.AddError(c, "Username or password is invalid")
			return Login()(c)
		}

		// Get session from storage
		sess, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		gob.Register(ldap.User{})
		sess.Set("uid", user.Id)
		sess.Set("user", user)

		if err := sess.Save(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Redirect("/")
	}
}

func Logout(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get session from storage
		sess, err := store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if err := sess.Destroy(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Redirect("/login")
	}
}
