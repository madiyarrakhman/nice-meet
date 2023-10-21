package models

import (
	"errors"
	"time"
)

var ErrMemberExistAlready = errors.New("member already exists")

type Member struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}

func AddMember(id int64, username, firstName, lastName string) *Member {
	return &Member{
		ID:        id,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "active",
	}
}
