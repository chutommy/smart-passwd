package data

import (
	"context"
)

// WordList represents a list of words.
type WordList interface {
	Word(ctx context.Context, length int16) (string, error)
	Close(ctx context.Context) error
}
