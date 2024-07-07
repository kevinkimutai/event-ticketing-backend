package api

import "github.com/kevinkimutai/ticketingapp/internal/app/domain"

type CategoryRepoPort interface {
	CreateCategory(*domain.Category) (domain.Category, error)
}

type CategoryRepo struct {
	db CategoryRepoPort
}

func NewCategoriesRepo(db CategoryRepoPort) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) CreateCategory(category *domain.Category) (domain.Category, error) {
	cat, err := r.db.CreateCategory(category)

	return cat, err
}
