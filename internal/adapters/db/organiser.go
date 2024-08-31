package db

import (
	"context"
	"math"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
)

func (db *DBAdapter) GetOrganisersByUserID(userID int64, params *domain.OrganiserParams) (domain.OrganisersFetch, error) {
	ctx := context.Background()

	organisers, err := db.queries.GetOrganisersByUserID(ctx, queries.GetOrganisersByUserIDParams{
		UserID: userID,
		Limit:  params.Limit,
		Offset: params.Page,
	})
	if err != nil {
		return domain.OrganisersFetch{}, err
	}

	count, err := db.queries.GetCountOrganisersByUserID(ctx, userID)
	if err != nil {
		return domain.OrganisersFetch{}, err
	}

	//Get Total amount made
	tAmount, err := db.queries.SumAmountEvents(ctx, userID)
	if err != nil {
		return domain.OrganisersFetch{}, err
	}

	//Get Page
	page := getPage(params.Page, params.Limit)

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
		Page:          page,
		NumberOfPages: uint(math.Ceil(float64(count) / float64((params.Limit)))),
		Total:         count,
		Data: domain.OrganiserData{
			TotalAmountEvents: float64(tAmount),
			Data:              orgs,
		},
	}, nil
}

func (db *DBAdapter) GetOrganiserEvent(eventID int64, params *domain.OrganiserParams) (domain.OrganiserEventFetch, error) {

	ctx := context.Background()

	organisers, err := db.queries.GetOrganisersEventByID(ctx, eventID)
	if err != nil {

		return domain.OrganiserEventFetch{}, err
	}

	count, err := db.queries.GetOrganisersEventCount(ctx, eventID)
	if err != nil {

		return domain.OrganiserEventFetch{}, err
	}

	tSums, err := db.queries.GetOrganisersEventSums(ctx, eventID)
	if err != nil {

		return domain.OrganiserEventFetch{}, err
	}

	//GetAdmitted
	admittedCount, err := db.queries.GetCountAdmittedOrganisersEventByID(ctx, eventID)
	if err != nil {

		return domain.OrganiserEventFetch{}, err
	}

	//notAdmitted
	notAdmittedCount, err := db.queries.GetCountNotAdmittedOrganisersEventByID(ctx, eventID)
	if err != nil {

		return domain.OrganiserEventFetch{}, err
	}

	//Get Page
	page := getPage(params.Page, params.Limit)

	var orgs []domain.OrganiserEvent

	for _, v := range organisers {
		org := domain.OrganiserEvent{
			AttendeeID:     v.AttendeeID,
			Fullname:       v.Fullname,
			Email:          v.Email,
			TicketTypeName: v.TicketTypeName.String,
			Quantity:       v.Quantity,
			Admitted:       v.Admitted.Bool,
			Total:          utils.ConvertNumericToFloat64(v.Total),
		}

		orgs = append(orgs, org)

	}

	return domain.OrganiserEventFetch{
		Page:               page,
		NumberOfPages:      uint(math.Ceil(float64(count) / float64((params.Limit)))),
		Total:              count,
		TicketsSold:        utils.ConvertNumericToFloat64(tSums.TotalSoldTickets.(pgtype.Numeric)),
		TotalAmount:        utils.ConvertNumericToFloat64(tSums.TotalPrice.(pgtype.Numeric)),
		TicketsAdmitted:    admittedCount,
		TicketsNotAdmitted: notAdmittedCount,
		Data:               orgs,
	}, nil
}

func (db *DBAdapter) DownloadOrganiserEvent(eventID int64) ([]domain.OrganiserEvent, error) {

	ctx := context.Background()

	organisers, err := db.queries.GetOrganisersEventByID(ctx, eventID)
	if err != nil {

		return []domain.OrganiserEvent{}, err
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

	return orgs, nil
}

func (db *DBAdapter) CheckIfUserIsOrganiser(userID int64, eventID int64) (bool, error) {
	ctx := context.Background()

	org, err := db.queries.GetOrganiserByEventID(ctx, eventID)
	if err != nil {
		return false, err
	}

	return org.UserID == userID, nil

}
