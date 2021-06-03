package data

// WordList represents a list of words.
type WordList interface {
	Word(length int16) (string, error)
	Close() error
}
