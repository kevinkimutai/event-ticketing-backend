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
