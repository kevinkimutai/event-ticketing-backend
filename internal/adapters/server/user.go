package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) UserRouter(api fiber.Router) {
	api.Get("/:userID", s.user.GetUser)
}
