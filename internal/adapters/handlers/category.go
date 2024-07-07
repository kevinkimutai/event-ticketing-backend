package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type CategoryApiPort interface {
	CreateCategory(*domain.Category) (domain.Category, error)
}

type CategoryService struct {
	api CategoryApiPort
}

func NewCategoryService(api CategoryApiPort) *CategoryService {
	return &CategoryService{
		api: api,
	}
}

func (s *CategoryService) CreateCategory(c *fiber.Ctx) error {
	category := &domain.Category{}

	//Bind To struct
	if err := c.BodyParser(&category); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	if category.Name == "" {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    errors.New("missing name field in category").Error(),
			})
	}

	cat, err := s.api.CreateCategory(category)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})

	}

	//JSON Response
	return c.Status(201).JSON(
		domain.CategoryResponse{
			StatusCode: 201,
			Message:    "Event created successfully",
			Data:       cat,
		})

}
