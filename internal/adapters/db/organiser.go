package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

func (db *DBAdapter) GetOrganisersByUserID(userID int64) (domain.OrganisersFetch, error) {
	ctx := context.Background()

	organisers, err := db.queries.GetOrganisersByUserID(ctx, userID)
	if err != nil {
		return domain.OrganisersFetch{}, err
	}

	count, err := db.queries.GetCountOrganisersByUserID(ctx, userID)
	if err != nil {
		return domain.OrganisersFetch{}, err
	}

	var orgs []domain.Organiser

	for _, v := range organisers {
		org := domain.Organiser{
			OrganiserID: v.OrganiserID,
			UserID:      v.UserID,
			EventID:     v.EventID,
			CreatedAt:   v.CreatedAt.Time,
		}

		orgs = append(orgs, org)

	}

	//TODO:will add pagination
	return domain.OrganisersFetch{
		Page:          0,
		NumberOfPages: 0,
		Total:         count[0],
		Data:          orgs,
	}, nil
}
