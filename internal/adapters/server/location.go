package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) LocationRouter(api fiber.Router) {

	api.Get("/:locationID", s.location.GetLocationByID)
	api.Get("/", s.location.GetLocations)
}
