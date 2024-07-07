package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type TicketTypeApiPort interface {
	CreateTicketType(t *domain.TicketType) (domain.TicketType, error)
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
	tickettype := &domain.TicketType{}

	//Bind To struct
	if err := c.BodyParser(&tickettype); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check Missing Values
	err := domain.NewTicketTypeDomain(tickettype)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//API
	newttype, err := s.api.CreateTicketType(tickettype)
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
