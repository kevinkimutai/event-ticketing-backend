package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type TicketOrderApiPort interface {
	CreateTicketOrder(order *domain.TicketOrderRequest, userID int64) (domain.TicketOrder, error)
}

type TicketOrderService struct {
	api TicketOrderApiPort
}

func NewTicketOrderService(api TicketOrderApiPort) *TicketOrderService {
	return &TicketOrderService{
		api: api,
	}
}

func (s *TicketOrderService) CreateTicketOrder(c *fiber.Ctx) error {

	//Get UserID from locals
	cus := c.Locals("customer")

	user, ok := cus.(queries.User)
	if !ok {
		fmt.Println("Type assertion failed, cus is not of type queries.User")

	}

	ticketOrder := &domain.TicketOrderRequest{}

	//Bind To struct
	if err := c.BodyParser(&ticketOrder); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//API
	order, err := s.api.CreateTicketOrder(ticketOrder, user.UserID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}

	//JSON Response
	return c.Status(201).JSON(
		domain.TicketOrderResponse{
			StatusCode: 201,
			Message:    "Event created successfully",
			Data:       order,
		})

}
