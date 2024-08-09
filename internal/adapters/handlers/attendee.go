package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type AttendeeApiPort interface {
	GetAttendeeByID(attendeeID int64) (domain.Attendee, error)
}

type AttendeeService struct {
	api AttendeeApiPort
}

func NewAttendeeService(api AttendeeApiPort) *AttendeeService {
	return &AttendeeService{
		api: api,
	}
}

func (s *AttendeeService) GetAttendee(c *fiber.Ctx) error {
	attendeeID := c.Params("attendeeID")

	//convert To int64
	attendeeIDInt64, err := strconv.ParseInt(attendeeID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	attendee, err := s.api.GetAttendeeByID(attendeeIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.AttendeeResponse{
		StatusCode: 200,
		Message:    "success",
		Data:       attendee,
	})
}
