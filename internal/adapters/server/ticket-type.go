package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) TicketTypeRouter(api fiber.Router) {
	api.Post("/", s.auth.IsAuthenticated, s.tickettype.CreateTicketType)
}
