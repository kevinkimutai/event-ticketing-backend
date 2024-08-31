package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type OrganiserRepoPort interface {
	GetOrganisersByUserID(userID int64, params *domain.OrganiserParams) (domain.OrganisersFetch, error)
	GetOrganiserEvent(eventID int64, params *domain.OrganiserParams) (domain.OrganiserEventFetch, error)
	DownloadOrganiserEvent(eventID int64) ([]domain.OrganiserEvent, error)
	GetEventByID(int64) (domain.Event, error)
	CheckIfUserIsOrganiser(userID int64, eventID int64) (bool, error)
}

type PDFRepoPort interface {
	GenerateAttendeesPDF([]domain.OrganiserEvent, domain.Event) ([]byte, error)
}

type OrganiserRepo struct {
	db  OrganiserRepoPort
	pdf PDFRepoPort
}

func NewOrganiserRepo(db OrganiserRepoPort, pdf PDFRepoPort) *OrganiserRepo {
	return &OrganiserRepo{
		db:  db,
		pdf: pdf,
	}
}

func (r *OrganiserRepo) GetOrganisersByUserID(userID int64, params *domain.OrganiserParams) (domain.OrganisersFetch, error) {
	org, err := r.db.GetOrganisersByUserID(userID, params)

	return org, err
}

func (r *OrganiserRepo) GetOrganiserEvent(eventID int64, params *domain.OrganiserParams) (domain.OrganiserEventFetch, error) {
	event, err := r.db.GetOrganiserEvent(eventID, params)

	return event, err
}

func (r *OrganiserRepo) DownloadOrganiserEvent(eventID int64) ([]byte, error) {
	orgAttendees, err := r.db.DownloadOrganiserEvent(eventID)
	if err != nil {
		return []byte{}, err
	}

	//Get Event
	e, err := r.db.GetEventByID(eventID)
	if err != nil {
		return []byte{}, err
	}

	//Generate PDF
	bytes, err := r.pdf.GenerateAttendeesPDF(orgAttendees, e)
	if err != nil {
		return bytes, err
	}

	return bytes, nil

}

func (r *OrganiserRepo) CheckIfUserIsOrganiser(userID int64, eventID int64) (bool, error) {
	bool, err := r.db.CheckIfUserIsOrganiser(userID, eventID)

	return bool, err
}
