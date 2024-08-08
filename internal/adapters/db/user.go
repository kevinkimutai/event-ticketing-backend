package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

func (db *DBAdapter) CreateUser(customer queries.CreateUserParams) (queries.User, error) {
	ctx := context.Background()
	cus, err := db.queries.CreateUser(ctx, customer)
	if err != nil {
		return cus, err
	}

	return cus, nil
}

func (db *DBAdapter) GetUserByEmail(email string) (queries.User, error) {
	ctx := context.Background()
	customer, err := db.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (db *DBAdapter) GetUserByUserId(userID int64) (domain.User, error) {
	ctx := context.Background()
	user, err := db.queries.GetUserByUserID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		UserID:   user.UserID,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}
