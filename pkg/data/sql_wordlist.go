package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/chutommy/smart-passwd/pkg/utils"
)

// SQLiteWordList represents a SQL-based implementation for WordList interface.
type SQLiteWordList struct {
	db *sql.DB
}

// ConnectSQLite connects to the given SQLite3 database and
// constructs a new SQLiteWordList linked with the DB.
func ConnectSQLite(file *utils.File) (*SQLiteWordList, error) {
	if file == nil {
		return nil, utils.ErrNilValue
	}

	if _, err := os.Stat(file.FilePath()); err != nil && os.IsNotExist(err) {
		return nil, fmt.Errorf("read database file: %w", err)
	}

	db, err := sql.Open("sqlite3", file.FilePath())
	if err != nil {
		return nil, fmt.Errorf("connect to sqlite database: %w", err)
	}

	return &SQLiteWordList{db}, nil
}

// Word returns a random word with length of l.
func (wl *SQLiteWordList) Word(_ context.Context, length int16) (string, error) {
	w, err := wl.randomWord(length)
	if err != nil {
		return "", fmt.Errorf("querying for a random word: %w", err)
	}

	return w, nil
}

// Close properly close the database connection.
func (wl *SQLiteWordList) Close(_ context.Context) error {
	if err := wl.db.Close(); err != nil {
		return fmt.Errorf("close wordlist: %w", err)
	}

	return nil
}
