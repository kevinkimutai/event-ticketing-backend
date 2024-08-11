package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type OrganiserApiPort interface {
	GetOrganisersByUserID(userID int64) (domain.OrganisersFetch, error)
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
