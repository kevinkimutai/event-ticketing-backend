package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
)

func (db *DBAdapter) GetOrderItemByTicketID(ticketID int64) (domain.TicketOrderItem, error) {
	ctx := context.Background()

	orderItem, err := db.queries.GetTicketOrderItemByTicketID(ctx, ticketID)

	return domain.TicketOrderItem{
		ItemID:     orderItem.ItemID,
		OrderID:    orderItem.OrderID,
		Quantity:   orderItem.Quantity,
		TotalPrice: utils.ConvertNumericToFloat64(orderItem.TotalPrice),
	}, err

}
