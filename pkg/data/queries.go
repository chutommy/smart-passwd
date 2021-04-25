package data

import (
	"database/sql"
	"errors"

	"github.com/chutified/smart-passwd/pkg/utils"
)

const (
	/*
		wordCount = `
		SELECT max(ROWID)
		FROM words;
		`

		wordRandom = `
		SELECT word
		FROM words
		WHERE id = (
			SELECT ABS(random()) % max(ROWID)
			FROM words
		);
		`
	*/

	wordRandomLen = `
	-- noinspection SqlResolve
	SELECT word FROM words
		WHERE length(word) = ?
		ORDER BY RANDOM()
		LIMIT 1;
		`
)

// ErrNoWords is returned by randomWord when there are no rows
// that satisfy the length condition.
var ErrNoWords = errors.New("there are no words in database")

// randomWords queries the table 'word' with the given sql
// database connection. It returns a random string value under
// the column 'word'.
func randomWord(db *sql.DB, l int16) (string, error) {
	if db == nil {
		return "", utils.ErrNilValue
	}

	var word string

	err := db.QueryRow(wordRandomLen, l).Scan(&word)
	if err != nil {
		return "", ErrNoWords
	}

	return word, nil
}
