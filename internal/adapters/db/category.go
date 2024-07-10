package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

func (db *DBAdapter) CreateCategory(category *domain.Category) (domain.Category, error) {
	ctx := context.Background()

	cat, err := db.queries.CreateCategory(ctx, category.Name)

	return domain.Category{
		CategoryID: cat.CategoryID,
		Name:       cat.Name,
	}, err
}

func (db *DBAdapter) GetCategories() ([]domain.Category, error) {
	ctx := context.Background()

	cats, err := db.queries.ListCategories(ctx)

	var categories []domain.Category

	for _, v := range cats {
		c := domain.Category{
			CategoryID: v.CategoryID,
			Name:       v.Name,
		}

		categories = append(categories, c)
	}

	return categories, err

}
