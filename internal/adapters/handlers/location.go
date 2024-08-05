package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type LocationApiPort interface {
	GetLocationByID(locationID int64) (domain.Location, error)
	GetLocations() ([]domain.Location, error)
}

type LocationService struct {
	api LocationApiPort
}

func NewLocationService(api LocationApiPort) *LocationService {
	return &LocationService{
		api: api,
	}
}

func (s *LocationService) GetLocationByID(c *fiber.Ctx) error {
	eventID := c.Params("locationID")

	//convert To int64
	locationIDInt64, err := strconv.ParseInt(eventID, 10, 32)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Get Product API
	loc, err := s.api.GetLocationByID(locationIDInt64)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	return c.Status(200).JSON(domain.LocationResponse{
		StatusCode: 200,
		Message:    "success",
		Data:       loc,
	})

}

func (s *LocationService) GetLocations(c *fiber.Ctx) error {

	//Get All Categories API
	categories, err := s.api.GetLocations()
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}
	return c.Status(200).JSON(
		domain.LocationsResponse{
			StatusCode: 200,
			Message:    "Successfully retrieved locations",
			Data:       categories,
		})
}
