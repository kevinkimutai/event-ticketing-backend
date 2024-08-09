package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type TicketOrderApiPort interface {
	CreateTicketOrder(order *domain.TicketOrderRequest, user queries.User) (domain.TicketOrder, error)
	GetTicketOrders(*domain.Params) ([]domain.TicketOrder, error)
	GetOrderItemByTicketID(ticketID int64) (domain.TicketOrderItem, error)
	GetOrder(orderItem int64) (domain.TicketOrder, error)
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
	order, err := s.api.CreateTicketOrder(ticketOrder, user)
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

func (s *TicketOrderService) GetTicketOrders(c *fiber.Ctx) error {
	//Get Query Params
	m := c.Queries()

	//Bind Params
	params := domain.CheckEventParams(m)

	//Get All TicketOrders API
	data, err := s.api.GetTicketOrders(params)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.TicketOrdersResponse{
			StatusCode: 200,
			Message:    "Successfully retrieved orders",
			Data:       data,
		})
}
func (s *TicketOrderService) GetOrder(c *fiber.Ctx) error {
	orderItemID := c.Params("orderID")

	//convert To int64
	orderItemIDInt64, err := strconv.ParseInt(orderItemID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	ticketOrder, err := s.api.GetOrder(orderItemIDInt64)
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
		Data:       ticketOrder,
	})
}

func (s *TicketOrderService) GetTicketOrderItem(c *fiber.Ctx) error {
	ticketID := c.Params("ticketID")

	//convert To int64
	ticketIDInt64, err := strconv.ParseInt(ticketID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	ticketOrderItem, err := s.api.GetOrderItemByTicketID(ticketIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.TicketOrderItemResponse{
		StatusCode: 200,
		Message:    "success",
		Data:       ticketOrderItem,
	})
}
