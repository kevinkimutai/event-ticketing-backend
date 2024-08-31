package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type EventApiPort interface {
	CreateEvent(event *domain.Event, userID int64) (domain.Event, error)
	GetEvents(*domain.Params) (domain.EventsFetch, error)
	GetEventByID(int64) (domain.Event, error)
	CheckIfUserIsOrganiser(userID int64, eventID int64) (bool, error)
	UpdateEvent(event *domain.Event, eventID int64) error
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

	//Get UserID from locals
	cus := c.Locals("customer")

	user, ok := cus.(queries.User)
	if !ok {
		fmt.Println("Type assertion failed, cus is not of type queries.User")

	}

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
	newEvent, err := s.api.CreateEvent(event, user.UserID)
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

func (s *EventService) GetEvents(c *fiber.Ctx) error {

	//Get Query Params
	m := c.Queries()

	//Bind To ProductParams
	params := domain.CheckEventParams(m)

	//Get All Products API
	data, err := s.api.GetEvents(params)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.EventsResponse{
			StatusCode:    200,
			Message:       "Successfully retrieved events",
			Page:          data.Page,
			NumberOfPages: data.NumberOfPages,
			Total:         data.Total,
			Data:          data.Data,
		})
}

func (s *EventService) GetEvent(c *fiber.Ctx) error {

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
	event, err := s.api.GetEventByID(eventIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.EventResponse{
		StatusCode: 200,
		Message:    "success",
		Data:       event,
	})
}
func (s *EventService) UpdateEvent(c *fiber.Ctx) error {
	//Check if user is organiser
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

	//Get UserID from locals
	cus := c.Locals("customer")

	user, ok := cus.(queries.User)
	if !ok {
		fmt.Println("Type assertion failed, cus is not of type queries.User")

	}

	bool, err := s.api.CheckIfUserIsOrganiser(user.UserID, eventIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	if !bool {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 401,
				Message:    "Unauthorised,login with proper rights",
			})
	}

	//Update Event
	event := &domain.Event{}

	//Bind To struct
	if err := c.BodyParser(&event); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	err = s.api.UpdateEvent(event, eventIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(
		domain.EventResponse{
			StatusCode: 200,
			Message:    "Event Updated successfully",
		})

}
