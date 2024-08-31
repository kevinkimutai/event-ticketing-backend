package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) TicketRouter(api fiber.Router) {
	api.Post("/", s.ticket.CreateTicketOrder)
	api.Get("/:orderID", s.ticket.GetOrder)
	api.Get("/:orderID/details", s.ticket.GetOrder)

	api.Get("/:ticketID/ticket-order-item", s.ticket.GetTicketOrderItem)

}
