package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) LocationRouter(api fiber.Router) {

	api.Get("/:locationID", s.location.GetLocationByID)
	api.Post("/", s.ticket.CreateTicketOrder)
}
