package repository

import (
	"database/sql"
	"errors"
	"nice-meet/src/models"
	"time"
)

type MemberRepository struct {
	db *sql.DB
}

func NewMember(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (m *MemberRepository) GetByID(id int64) (*models.Member, error) {
	sqlStr := `SELECT id, username, firstName, lastName, createdAt, updatedAt, status FROM members WHERE id = ?`
	row := m.db.QueryRow(sqlStr, id)

	var member models.Member
	err := row.Scan(&member.ID, &member.Username, &member.FirstName, &member.LastName, &member.CreatedAt, &member.UpdatedAt, &member.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (m *MemberRepository) AddMember(id int64, username, firstName, lastName string, status string, createdAt, updatedAt time.Time) error {
	sqlStr := `INSERT INTO members (id, username, firstName, lastName, createdAt, updatedAt, status) VALUES (?, ?, ?, ?, ?, ?, ?)`
	stmt, err := m.db.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, username, firstName, lastName, createdAt, updatedAt, status)

	if err != nil {
		return err
	}

	return nil
}
