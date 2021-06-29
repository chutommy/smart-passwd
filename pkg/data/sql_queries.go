package data

import (
	"fmt"

	"github.com/chutommy/smart-passwd/pkg/utils"
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

// randomWords queries the table 'word' with the given sql
// database connection. It returns a random string value under
// the column 'word'.
func (wl *SQLiteWordList) randomWord(l int16) (string, error) {
	if wl.db == nil {
		return "", utils.ErrNilValue
	}

	var word string

	err := wl.db.QueryRow(wordRandomLen, l).Scan(&word)
	if err != nil {
		return "", fmt.Errorf("failed to query for a random word of length: %d: %w", l, err)
	}

	return word, nil
}
