package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
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

func (db *DBAdapter) GetOrganiserEvent(eventID int64) (domain.OrganiserEventFetch, error) {

	ctx := context.Background()

	organisers, err := db.queries.GetOrganisersEventByID(ctx, eventID)
	if err != nil {
		fmt.Println("1")
		return domain.OrganiserEventFetch{}, err
	}

	count, err := db.queries.GetOrganisersEventCount(ctx, eventID)
	if err != nil {
		fmt.Println("2")
		return domain.OrganiserEventFetch{}, err
	}

	tSums, err := db.queries.GetOrganisersEventSums(ctx, eventID)
	if err != nil {
		fmt.Println("3")
		return domain.OrganiserEventFetch{}, err
	}

	var orgs []domain.OrganiserEvent

	for _, v := range organisers {
		org := domain.OrganiserEvent{
			AttendeeID:     v.AttendeeID,
			Fullname:       v.Fullname,
			Email:          v.Email,
			TicketTypeName: v.TicketTypeName.String,
			Quantity:       v.Quantity,
			Total:          utils.ConvertNumericToFloat64(v.Total),
		}

		orgs = append(orgs, org)

	}

	//TODO:will add pagination
	return domain.OrganiserEventFetch{
		Page:          0,
		NumberOfPages: 0,
		Total:         count,
		TicketsSold:   utils.ConvertNumericToFloat64(tSums.TotalSoldTickets.(pgtype.Numeric)),
		TotalAmount:   utils.ConvertNumericToFloat64(tSums.TotalPrice.(pgtype.Numeric)),
		Data:          orgs,
	}, nil
}
