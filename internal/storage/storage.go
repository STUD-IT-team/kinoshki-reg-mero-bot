package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Storage struct {
	db *sql.DB
	p  ParticipantStorage
}

func NewStorage(config Config, p ParticipantStorage) (*Storage, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db1, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return &Storage{
		db1,
		p,
	}, nil
}

func (s *Storage) CloseStorage() error {
	return s.db.Close()
}
