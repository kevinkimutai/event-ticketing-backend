package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) UsherRouter(api fiber.Router) {
	api.Get("/order/:orderID", s.auth.AllowedRoles("usher"), s.usher.GetTicketOrderDetails)
	api.Patch("/order/:orderID", s.auth.AllowedRoles("usher"), s.usher.AdmitTicketOrder)

}
