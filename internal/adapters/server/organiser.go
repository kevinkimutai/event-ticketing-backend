package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) OrganiserRouter(api fiber.Router) {

	api.Get("/user", s.organiser.GetOrganiserByUserID)
	api.Get("/event/:eventID", s.organiser.GetOrganiserEvent)

}
