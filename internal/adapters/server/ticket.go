package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) TicketRouter(api fiber.Router) {
	api.Post("/", s.auth.IsAuthenticated, s.ticket.CreateTicketOrder)
}
