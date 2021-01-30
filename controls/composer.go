package controls

import (
	"strings"

	"github.com/pkg/errors"
)

// composeWords composes the slice of words into the one string and returns it + the spaced phrase.
func (c *Controller) composeWords(ws []word) (string, string, error) {
	// define the phrase
	phrase := make([]string, 0)

	// range over the structures and generates the words for them
	for _, w := range ws {
		// get the random index
		capacity, err := c.ds.Len(w.length, w.category)
		if err != nil {
			return "", "", errors.Wrap(err, "count the max length")
		}

		randI := c.rng.Intn(capacity)

		// generate the word
		newW, err := c.ds.Gen(w.length, w.category, randI)
		if err != nil {
			return "", "", errors.Wrap(err, "generating random word")
		}

		// add to the result
		phrase = append(phrase, newW)
	}

	// join the phrase
	smashed := strings.Join(phrase, "")
	spaced := strings.Join(phrase, " ")

	return smashed, spaced, nil
}
