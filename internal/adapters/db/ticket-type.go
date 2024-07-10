package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
)

func (db *DBAdapter) CreateTicketType(ttype *domain.TicketType) (domain.TicketType, error) {
	ctx := context.Background()
	tickettype, err := db.queries.CreateTicketType(ctx, queries.CreateTicketTypeParams{
		Name:             utils.ConvertStringToText(ttype.Name),
		Price:            utils.ConvertFloat64ToNumeric(ttype.Price),
		TotalTickets:     ttype.TotalTickets,
		RemainingTickets: ttype.RemainingTickets,
		EventID:          ttype.EventID,
	})

	return domain.TicketType{
		TicketTypeID:     tickettype.TicketTypeID,
		Name:             tickettype.Name.String,
		Price:            utils.ConvertNumericToFloat64(tickettype.Price),
		TotalTickets:     tickettype.TotalTickets,
		RemainingTickets: tickettype.RemainingTickets,
		EventID:          tickettype.EventID,
	}, err

}

func (db *DBAdapter) GetTicketTypesByEvent(eventID int64) ([]domain.TicketType, error) {
	ttypes, err := db.queries.GetTicketTypesByEvent(context.Background(), eventID)

	var tt []domain.TicketType

	for _, v := range ttypes {
		ticketType := domain.TicketType{
			EventID:          v.EventID,
			TicketTypeID:     v.TicketTypeID,
			Name:             v.Name.String,
			Price:            utils.ConvertNumericToFloat64(v.Price),
			TotalTickets:     v.TotalTickets,
			RemainingTickets: v.RemainingTickets,
		}

		tt = append(tt, ticketType)
	}

	return tt, err
}
