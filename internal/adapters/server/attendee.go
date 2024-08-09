package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) AttendeeRouter(api fiber.Router) {

	api.Get("/:attendeeID", s.attendee.GetAttendee)
}
