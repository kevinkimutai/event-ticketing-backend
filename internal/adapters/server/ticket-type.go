package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) TicketTypeRouter(api fiber.Router) {
	// Admin Get All TickeTypes
	api.Post("/", s.tickettype.CreateTicketType)
}
