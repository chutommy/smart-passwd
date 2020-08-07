package data

import (
	"database/sql"

	"github.com/pkg/errors"
)

const (
	// N is a noun placeholder
	Noun = 1
	// V is a verb placeholder
	Verb = 2
	// A is an adjective placeholder
	Adjective = 3
	// R is an adverb placeholder
	Adverb = 4
)

// Len returns the number of words with the lenght l in category c.
func (s *Service) Len(l int, c int) (int, error) {

	// define query
	query := `
	SELECT COUNT(*)
	FROM words
	WHERE LENGTH(word) = $1
	AND category_id = $2`

	// run query
	row := s.db.QueryRow(query, l, c)

	// get the count
	var count int
	if err := row.Scan(&count); err == sql.ErrNoRows {
		return 0, errors.Wrap(err, "no rows")
	} else if err != nil {
		return 0, errors.Wrap(err, "db internal errror")
	}

	return count, nil
}

// Gen returns an i-th word with the length l in category c.
func (s *Service) Gen(l int, c int, i int) (string, error) {

	// define query
	query := `
	SELECT word
	FROM words
	WHERE LENGTH(word) = $1
	AND category_id = $2
	OFFSET $3
	LIMIT 1`

	// run query
	row := s.db.QueryRow(query, l, c, i)

	// get the word
	var word string
	if err := row.Scan(&word); err == sql.ErrNoRows {
		return "", errors.Wrap(err, "no rows")
	} else if err != nil {
		return "", errors.Wrap(err, "db internal errror")
	}

	return word, nil
}
