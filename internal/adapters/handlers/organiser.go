package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type OrganiserApiPort interface {
	GetOrganisersByUserID(userID int64) (domain.OrganisersFetch, error)
	GetOrganiserEvent(eventID int64) (domain.OrganiserEventFetch, error)
}

type OrganiserService struct {
	api OrganiserApiPort
}

func NewOrganiserService(api OrganiserApiPort) *OrganiserService {
	return &OrganiserService{
		api: api,
	}
}

func (s *OrganiserService) GetOrganiserByUserID(c *fiber.Ctx) error {

	//Get UserID from locals
	cus := c.Locals("customer")

	user, ok := cus.(queries.User)
	if !ok {
		fmt.Println("Type assertion failed, cus is not of type queries.User")

	}

	organisers, err := s.api.GetOrganisersByUserID(user.UserID)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.OrganisersResponse{
		StatusCode:    200,
		Message:       "success",
		Page:          organisers.Page,
		NumberOfPages: organisers.NumberOfPages,
		Total:         organisers.Total,
		Data:          organisers.Data,
	})
}

func (s *OrganiserService) GetOrganiserEvent(c *fiber.Ctx) error {
	eventID := c.Params("eventID")

	//convert To int64
	eventIDInt64, err := strconv.ParseInt(eventID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	event, err := s.api.GetOrganiserEvent(eventIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.OrganiserEventResponse{
		StatusCode:    200,
		Message:       "success",
		Page:          event.Page,
		NumberOfPages: event.NumberOfPages,
		Total:         event.Total,
		TicketsSold:   event.TicketsSold,
		TotalAmount:   event.TotalAmount,
		Data:          event.Data,
	})
}
