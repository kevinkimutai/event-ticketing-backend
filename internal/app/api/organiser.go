package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type OrganiserRepoPort interface {
	GetOrganisersByUserID(userID int64) (domain.OrganisersFetch, error)
}

type OrganiserRepo struct {
	db OrganiserRepoPort
}

func NewOrganiserRepo(db OrganiserRepoPort) *OrganiserRepo {
	return &OrganiserRepo{
		db: db,
	}
}

func (r *OrganiserRepo) GetOrganisersByUserID(userID int64) (domain.OrganisersFetch, error) {
	org, err := r.db.GetOrganisersByUserID(userID)

	return org, err
}
