package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type TicketTypeApiPort interface {
	CreateTicketType(t *domain.TicketType, eventID int64) (domain.TicketType, error)
	GetTicketTypesByEvent(int64) ([]domain.TicketType, error)
}

type TicketTypeService struct {
	api TicketTypeApiPort
}

func NewTicketTypeService(api TicketTypeApiPort) *TicketTypeService {
	return &TicketTypeService{
		api: api,
	}
}

func (s *TicketTypeService) CreateTicketType(c *fiber.Ctx) error {
	eventID := c.Params("eventID")

	tickettype := &domain.TicketType{}

	//Bind To struct
	if err := c.BodyParser(&tickettype); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	intEvent, err := strconv.ParseInt(eventID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check Missing Values
	err = domain.NewTicketTypeDomain(tickettype)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//API
	newttype, err := s.api.CreateTicketType(tickettype, intEvent)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}

	//JSON Response
	return c.Status(201).JSON(
		domain.TicketTypeResponse{
			StatusCode: 201,
			Message:    "Event created successfully",
			Data:       newttype,
		})

}

func (s *TicketTypeService) GetTicketTypesByEvent(c *fiber.Ctx) error {
	eventID := c.Params("eventID")

	intEvent, err := strconv.ParseInt(eventID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//API
	tickettypes, err := s.api.GetTicketTypesByEvent(intEvent)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}

	//JSON Response
	return c.Status(200).JSON(
		domain.TicketTypesResponse{
			StatusCode: 200,
			Message:    "Ticket Types retrieved successfully",
			Data:       tickettypes,
		})

}
