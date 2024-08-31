package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type UsherApiPort interface {
	GetTicketOrderDetails(orderID int64) (domain.TicketOrderDetails, error)
	AdmitTicketOrder(orderId int64) error
}

type UsherService struct {
	api UsherApiPort
}

func NewUsherService(api UsherApiPort) *UsherService {
	return &UsherService{
		api: api,
	}
}

func (s *UsherService) GetTicketOrderDetails(c *fiber.Ctx) error {

	orderID := c.Params("orderID")

	//convert To int64
	orderIDInt64, err := strconv.ParseInt(orderID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	tOrder, err := s.api.GetTicketOrderDetails(orderIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.TicketOrderDetailsResponse{
		StatusCode: 200,
		Message:    "success",
		Data:       tOrder,
	})
}
func (s *UsherService) AdmitTicketOrder(c *fiber.Ctx) error {
	orderID := c.Params("orderID")

	//convert To int64
	orderIDInt64, err := strconv.ParseInt(orderID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	err = s.api.AdmitTicketOrder(orderIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.TicketOrderResponse{
		StatusCode: 200,
		Message:    "success",
		// Data:       tOrder,
	})
}
