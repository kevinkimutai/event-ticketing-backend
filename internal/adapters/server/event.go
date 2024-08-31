package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) EventRouter(api fiber.Router) {

	api.Post("/", s.event.CreateEvent)
	api.Get("/", s.event.GetEvents)
	api.Get("/:eventID", s.event.GetEvent)
	api.Patch("/:eventID", s.event.UpdateEvent)
	api.Get("/:eventID/ticket-types", s.tickettype.GetTicketTypesByEvent)
	api.Post("/:eventID/ticket-types", s.tickettype.CreateTicketType)

}
