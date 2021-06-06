package wasm

import (
	"context"
	"fmt"
	"time"

	"github.com/chutified/smart-passwd/pkg/data"
)

// Runner represents a runner which is in control of the WordList connection.
type Runner struct {
	WordList *data.MongoWordList
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewRunner is a contructor of the Runner controller.
func NewRunner(wl *data.MongoWordList) *Runner {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	return &Runner{
		WordList: wl,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Gen generates a random word with length of l.
func (r *Runner) Gen(l int16) (string, error) {
	w, err := r.WordList.Word(r.ctx, l)
	if err != nil {
		return "", fmt.Errorf("generate word from wordlist: %w", err)
	}

	return w, nil
}

// Stop stops all idle connections.
func (r *Runner) Stop() error {
	if err := r.WordList.Close(r.ctx); err != nil {
		return fmt.Errorf("closing wordlist: %w", err)
	}

	r.cancel()

	return nil
}
