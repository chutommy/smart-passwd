package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chutified/smart-passwd/pkg/utils"
)

// SQLWordList represents a SQL-based implementation for WordList interface.
type SQLWordList struct {
	db *sql.DB
}

// ConnectSQLite connects to the given SQLite3 database and
// constructs a new SQLWordList linked with the DB.
func ConnectSQLite(file *utils.File) (*SQLWordList, error) {
	if file == nil {
		return nil, utils.ErrNilValue
	}

	if _, err := os.Stat(file.FilePath()); err != nil && os.IsNotExist(err) {
		return nil, fmt.Errorf("read database file: %w", err)
	}

	db, err := sql.Open("sqlite", file.FilePath())
	if err != nil {
		return nil, fmt.Errorf("connect to sqlite database: %w", err)
	}

	return &SQLWordList{db}, nil
}

// Word returns a random word with length of l.
func (wl *SQLWordList) Word(length int16) (string, error) {
	w, err := randomWord(wl.db, length)
	if err != nil {
		return "", fmt.Errorf("querying for a random word: %w", err)
	}

	return w, nil
}

// Close properly close the database connection.
func (wl *SQLWordList) Close() error {
	if err := wl.db.Close(); err != nil {
		return fmt.Errorf("close wordlist: %w", err)
	}

	return nil
}
