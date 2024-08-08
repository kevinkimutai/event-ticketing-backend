package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type UserApiPort interface {
	GetUser(userID int64) (domain.User, error)
}

type UserService struct {
	api UserApiPort
}

func NewUserService(api UserApiPort) *UserService {
	return &UserService{
		api: api,
	}
}

func (s *UserService) GetUser(c *fiber.Ctx) error {

	// //Telemetry Tracer
	// ctx := c.UserContext()
	// _, span := s.t.Tracer().Start(ctx, "CreateEvent")
	// defer span.End()

	userID := c.Params("userID")

	//convert To int64
	userIDInt64, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	user, err := s.api.GetUser(userIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.UserResponse{
		StatusCode: 200,
		Message:    "success",
		Data:       user,
	})
}
