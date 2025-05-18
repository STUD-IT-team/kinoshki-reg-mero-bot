package storage

import (
	"fmt"
)

type Participant struct {
	ID         int
	FullName   string
	StudyGroup string
	Phone      string
	Telegram   string
}

type ParticipantStorage interface {
	CreateParticipant(p Participant) (int, error)
	GetParticipant(id int) (Participant, error)
	UpdateParticipant(p Participant) error
	DeleteParticipant(id int) error
}

func (s *Storage) CreateParticipant(p Participant) (int, error) {
	var id int
	err := s.db.QueryRow(
		"INSERT INTO participants (full_name, study_group, phone, telegram) VALUES ($1, $2, $3, $4) RETURNING id",
		p.FullName, p.StudyGroup, p.Phone, p.Telegram,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to  participant  %w", err)
	}

	return id, nil
}

func (s *Storage) GetParticipant(id int) (Participant, error) {
	var p Participant
	err := s.db.QueryRow(
		"SELECT id, full_name, study_group, phone, telegram FROM participants WHERE id = $1",
		id,
	).Scan(&p.ID, &p.FullName, &p.StudyGroup, &p.Phone, &p.Telegram)

	if err != nil {
		return Participant{}, fmt.Errorf("failed to get participant: %w", err)
	}

	return p, nil
}

func (s *Storage) UpdateParticipant(p Participant) error {
	res, err := s.db.Exec(
		"UPDATE participants SET full_name = $1, study_group = $2, phone = $3, telegram = $4 WHERE id = $5",
		p.FullName, p.StudyGroup, p.Phone, p.Telegram, p.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update participant: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows  %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("not found %w", err)
	}

	return nil
}

func (s *Storage) DeleteParticipant(id int) error {
	res, err := s.db.Exec("DELETE FROM participants WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete participant %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("not found %w", err)
	}

	return nil
}
