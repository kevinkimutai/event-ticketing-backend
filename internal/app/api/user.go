package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type UserRepoPort interface {
	GetUserByUserId(userID int64) (domain.User, error)
}

type UserRepo struct {
	db UserRepoPort
}

func NewUserRepo(db UserRepoPort) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUser(userID int64) (domain.User, error) {
	user, err := r.db.GetUserByUserId(userID)

	return user, err
}
