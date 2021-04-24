package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chutified/smart-passwd/pkg/utils"
)

// WordList represents a list of words.
type WordList struct {
	db *sql.DB
}

// Connect connects to the given SQLite3 database and
// constructs a new WordList linked with the DB.
func Connect(f *utils.File) (*WordList, error) {
	if _, err := os.Stat(f.FilePath()); err != nil && os.IsNotExist(err) {
		return nil, fmt.Errorf("read database file: %w", err)
	}

	db, err := sql.Open("sqlite3", f.FilePath())
	if err != nil {
		return nil, fmt.Errorf("connect to sqlite3 database: %w", err)
	}

	return &WordList{db}, nil
}

// Word returns a random word with length of l.
func (wl *WordList) Word(l int16) (string, error) {
	w, err := randomWord(wl.db, l)
	if err != nil {
		return "", fmt.Errorf("querying for a random word: %w", err)
	}

	return w, nil
}

// Close properly close the database connection.
func (wl *WordList) Close() error {
	if err := wl.db.Close(); err != nil {
		return fmt.Errorf("close wordlist: %w", err)
	}

	return nil
}
