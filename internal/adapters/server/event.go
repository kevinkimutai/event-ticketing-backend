package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) EventRouter(api fiber.Router) {
	api.Post("/", s.event.CreateEvent)
	api.Get("/", s.event.GetEvents)
}
