package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type UsherRepoPort interface {
	GetTicketOrderDetails(orderID int64) (domain.TicketOrderDetails, error)
	AdmitTicketOrder(orderId int64) error
}

type UsherRepo struct {
	db UsherRepoPort
}

func NewUsherRepo(db UsherRepoPort) *UsherRepo {
	return &UsherRepo{
		db: db,
	}
}

func (r *UsherRepo) GetTicketOrderDetails(orderID int64) (domain.TicketOrderDetails, error) {
	tOrder, err := r.db.GetTicketOrderDetails(orderID)

	return tOrder, err
}
func (r *UsherRepo) AdmitTicketOrder(orderId int64) error {
	err := r.db.AdmitTicketOrder(orderId)

	return err
}
