package services

import (
	"nice-meet/src/models"
	"nice-meet/src/repository"
	"time"
)

type Repository interface {
	GetByID(id int64) (*models.Member, error)
	AddMember(id int64, username, firstName, lastName string, status string, createdAt, updatedAt time.Time) error
}

type Member struct {
	rep Repository
}

func NewMember(rep *repository.MemberRepository) *Member {
	return &Member{rep: rep}
}

func (m *Member) AddMember(id int64, username, firstName, lastName string) error {
	member, errGetByID := m.rep.GetByID(id)
	if errGetByID != nil {
		return errGetByID
	}

	if member != nil {
		return models.ErrMemberExistAlready
	}

	err := m.rep.AddMember(
		id,
		username,
		firstName,
		lastName,
		"active",
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
