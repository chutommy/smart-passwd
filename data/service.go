package data

import (
	"database/sql"
	"fmt"

	config "github.com/chutified/smart-passwd/config"
	"github.com/pkg/errors"
)

// Service provides all actions with the database.
type Service struct {
	db *sql.DB
}

// New is a Service's constructor.
func New() *Service {
	return &Service{}
}

// Init initialize the database connection.
func (s *Service) Init(cfg config.DBConfig) error {

	// define database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)

	// connect to db
	var err error
	s.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return errors.Wrap(err, "connecting to the database")
	}

	// check the db connecton
	err = s.db.Ping()
	if err != nil {
		return errors.Wrap(err, "ping to DB conn")
	}

	return nil
}

// Stop closes the database connection.
func (s *Service) Stop() error {

	// closure
	if err := s.db.Close(); err != nil {
		return errors.Wrap(err, "failed to close the database connection")
	}
	return nil
}
