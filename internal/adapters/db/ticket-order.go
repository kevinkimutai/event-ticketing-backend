package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
)

func (db *DBAdapter) CreateTicketOrder(t *domain.TicketOrderRequest, userID int64) (domain.TicketOrder, error) {

	ctx := context.Background()

	//Start TX
	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return domain.TicketOrder{}, errors.New("failed to start tx")
	}

	qtx := db.queries.WithTx(tx)

	//	Create Attendee
	attendee, err := qtx.CreateAttendee(ctx, userID)
	if err != nil {
		tx.Rollback(ctx)
		fmt.Println("error occured 1")
		return domain.TicketOrder{}, err
	}

	var pgtype8 pgtype.Int8
	pgtype8.Scan(attendee.AttendeeID)

	//Create Ticket Order
	ticketOrder, err := qtx.CreateTicketOrder(ctx, pgtype8)
	if err != nil {
		tx.Rollback(ctx)
		fmt.Println("error occured 2")

		return domain.TicketOrder{}, err
	}

	var orderItemsReq []domain.TicketOrderItem
	var total float64

	//Calculate Total Price
	for _, v := range t.OrderItems {
		ttype, err := qtx.GetTicketType(ctx, v.TicketTypeID)
		if err != nil {
			tx.Rollback(ctx)
			fmt.Println(v.TicketTypeID)
			return domain.TicketOrder{}, err
		}

		//Check If Tickets are enough
		if ttype.RemainingTickets < int32(v.Quantity) {
			tx.Rollback(ctx)
			err = fmt.Errorf("cannot complete order, only %v tickets available", ttype.RemainingTickets)
			return domain.TicketOrder{}, err
		}

		orderItemsReq = append(orderItemsReq, domain.TicketOrderItem{
			OrderID:      ticketOrder.OrderID,
			TicketTypeID: ttype.TicketTypeID,
			Quantity:     v.Quantity,
			TotalPrice:   (float64(v.Quantity) * utils.ConvertNumericToFloat64(ttype.Price)),
		})
		total += (float64(v.Quantity) * utils.ConvertNumericToFloat64(ttype.Price))

		//Update Remaining Tickets
		err = qtx.UpdateRemainingTicketType(ctx, queries.UpdateRemainingTicketTypeParams{
			TicketTypeID:     v.TicketTypeID,
			RemainingTickets: ttype.RemainingTickets - int32(v.Quantity),
		})
		if err != nil {
			tx.Rollback(ctx)
			fmt.Println("error occured 4")
			return domain.TicketOrder{}, err
		}
	}

	var orderItems []domain.OrderItem

	//Create Ticket
	for _, v := range orderItemsReq {
		ticket, err := qtx.CreateTicket(ctx, v.TicketTypeID)
		if err != nil {
			tx.Rollback(ctx)
			fmt.Println("error occured 5")
			return domain.TicketOrder{}, err
		}

		item := domain.OrderItem{
			OrderID:    ticketOrder.OrderID,
			TicketID:   ticket.TicketID,
			Quantity:   v.Quantity,
			TotalPrice: v.TotalPrice,
		}
		orderItems = append(orderItems, item)

	}

	//Create TicketOrder Items With Total Price

	for _, v := range orderItems {
		qtx.CreateTicketOrderItem(ctx, queries.CreateTicketOrderItemParams{
			OrderID:    v.OrderID,
			TicketID:   v.TicketID,
			Quantity:   v.Quantity,
			TotalPrice: utils.ConvertFloat64ToNumeric(v.TotalPrice),
		})
	}

	//Update Total Amount Order
	err = qtx.UpdateTotalAmountOrder(ctx, queries.UpdateTotalAmountOrderParams{
		OrderID:     ticketOrder.OrderID,
		TotalAmount: utils.ConvertFloat64ToNumeric(total),
	})

	if err != nil {
		tx.Rollback(ctx)
		fmt.Println("error occured 6")
		return domain.TicketOrder{}, err
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return domain.TicketOrder{}, err
	}

	return domain.TicketOrder{
		OrderID:     ticketOrder.OrderID,
		AttendeeID:  ticketOrder.AttendeeID.Int64,
		TotalAmount: total,
		CreatedAt:   ticketOrder.CreatedAt.Time,
	}, nil

}

func (db *DBAdapter) GetTicketOrders(params domain.Params) ([]domain.TicketOrder, error) {
	torders, err := db.queries.GetTicketOrders(context.Background(), queries.GetTicketOrdersParams{
		Limit:  params.Limit,
		Offset: params.Page,
	})

	var ticketOrders []domain.TicketOrder
	for _, v := range torders {
		order := domain.TicketOrder{
			OrderID:     v.OrderID,
			PaymentID:   v.PaymentID,
			AttendeeID:  v.PaymentID.Int64,
			TotalAmount: utils.ConvertNumericToFloat64(v.TotalAmount),
			CreatedAt:   v.CreatedAt.Time,
		}

		ticketOrders = append(ticketOrders, order)

	}

	return ticketOrders, err
}
