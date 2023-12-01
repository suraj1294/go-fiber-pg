package handler

import (
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	AppStore
}

func (uh *UsersHandler) MockUsers(c *fiber.Ctx) error {

	movies := uh.UsersStore.MockUsers()

	c.Accepts("application/json")

	c.JSON(movies)

	return nil
}

func (uh *UsersHandler) GetUserByEmail(c *fiber.Ctx) error {

	id := c.Params("email")

	user, err := uh.UsersStore.GetUserByEmail(id)

	if err != nil {
		return err
	}

	c.Accepts("application/json")

	c.JSON(user)

	return nil
}
