package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type EventApiPort interface {
	CreateEvent(event *domain.Event) (domain.Event, error)
}

type EventService struct {
	api EventApiPort
}

func NewEventService(api EventApiPort) *EventService {
	return &EventService{
		api: api,
	}
}

func (s *EventService) CreateEvent(c *fiber.Ctx) error {
	event := &domain.Event{}

	//Bind To struct
	if err := c.BodyParser(&event); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check Missing Values
	err := domain.NewEventDomain(event)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//API
	newEvent, err := s.api.CreateEvent(event)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}

	//JSON Response
	return c.Status(201).JSON(
		domain.EventResponse{
			StatusCode: 201,
			Message:    "Event created successfully",
			Data:       newEvent,
		})

}
