package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
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
