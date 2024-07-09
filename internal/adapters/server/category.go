package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) CategoryRouter(api fiber.Router) {
	api.Post("/", s.category.CreateCategory)
}
