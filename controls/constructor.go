package controls

import (
	"errors"

	data "github.com/chutified/smart-passwd/data"
)

// word is the basic structure of the generating phrase.
type word struct {
	category int
	length   int
}

// newWord generates a new word.
func newWord(c int, l int) word {
	return word{
		category: c,
		length:   l,
	}
}

// newPhrase generates the meaningful english phrase.
func (c *Controller) newPhrase(l int) ([]word, error) {

	// check the length
	if l < 5 || l > 32 {
		return nil, errors.New("invalid lengh (5-36)")
	}

	// create a random influencer
	ranN1 := c.rng.Intn(2) == 1
	ranN2 := c.rng.Intn(2) == 1
	ranN3 := c.rng.Intn(2) == 1

	// prepare the placeholder for the result
	var structure []word
	switch {

	// (5-8)
	case l < 9:
		w1 := newWord(data.Adjective, l)
		structure = append(structure, w1)

	// (9-12)
	case l < 13:

		// random factor
		if ranN1 && ranN2 {
			w1 := newWord(data.Adjective, 4)
			w2 := newWord(data.Noun, l-4)
			structure = append(structure, w1, w2)
		} else if ranN1 && !ranN2 {
			w1 := newWord(data.Adjective, 5)
			w2 := newWord(data.Noun, l-5)
			structure = append(structure, w1, w2)
		} else if !ranN1 && ranN2 {
			w1 := newWord(data.Adjective, 6)
			w2 := newWord(data.Noun, l-6)
			structure = append(structure, w1, w2)
		} else {
			w1 := newWord(data.Adverb, 5)
			w2 := newWord(data.Verb, l-5)
			structure = append(structure, w1, w2)
		}

	// (13-18)
	case l < 19:

		// random factor
		if ranN1 && ranN2 && ranN3 {
			w1 := newWord(data.Noun, 5)
			w2 := newWord(data.Adverb, 5)
			w3 := newWord(data.Verb, l-10)
			structure = append(structure, w1, w2, w3)
		} else if !ranN1 && ranN2 && ranN3 {
			w1 := newWord(data.Noun, 4)
			w2 := newWord(data.Adverb, 6)
			w3 := newWord(data.Verb, l-10)
			structure = append(structure, w1, w2, w3)
		} else if ranN1 && !ranN2 && ranN3 {
			w1 := newWord(data.Adjective, 5)
			w2 := newWord(data.Noun, 5)
			w3 := newWord(data.Verb, l-10)
			structure = append(structure, w1, w2, w3)
		} else if ranN1 && ranN2 && !ranN3 {
			w1 := newWord(data.Adjective, 6)
			w2 := newWord(data.Noun, 4)
			w3 := newWord(data.Verb, l-10)
			structure = append(structure, w1, w2, w3)
		} else if !ranN1 && !ranN2 && ranN3 {
			w1 := newWord(data.Adjective, 5)
			w2 := newWord(data.Adjective, 5)
			w3 := newWord(data.Noun, l-10)
			structure = append(structure, w1, w2, w3)
		} else if !ranN1 && ranN2 && !ranN3 {
			w1 := newWord(data.Adjective, 6)
			w2 := newWord(data.Adjective, 4)
			w3 := newWord(data.Noun, l-10)
			structure = append(structure, w1, w2, w3)
		} else if ranN1 && !ranN2 && !ranN3 {
			w1 := newWord(data.Noun, 5)
			w2 := newWord(data.Verb, 5)
			w3 := newWord(data.Noun, l-10)
			structure = append(structure, w1, w2, w3)
		} else {
			w1 := newWord(data.Noun, 4)
			w2 := newWord(data.Verb, 6)
			w3 := newWord(data.Noun, l-10)
			structure = append(structure, w1, w2, w3)
		}

	// (20-25)
	case l < 26:

		// random factor
		if ranN1 && ranN2 && ranN3 {
			w1 := newWord(data.Noun, 5)
			w2 := newWord(data.Adverb, 6)
			w3 := newWord(data.Verb, 6)
			w4 := newWord(data.Noun, l-17)
			structure = append(structure, w1, w2, w3, w4)
		} else if !ranN1 && ranN2 && ranN3 {
			w1 := newWord(data.Noun, 6)
			w2 := newWord(data.Adverb, 5)
			w3 := newWord(data.Verb, 6)
			w4 := newWord(data.Noun, l-17)
			structure = append(structure, w1, w2, w3, w4)
		} else if ranN1 && !ranN2 && ranN3 {
			w1 := newWord(data.Adjective, 6)
			w2 := newWord(data.Noun, 6)
			w3 := newWord(data.Verb, 5)
			w4 := newWord(data.Noun, l-17)
			structure = append(structure, w1, w2, w3, w4)
		} else if ranN1 && ranN2 && !ranN3 {
			w1 := newWord(data.Adjective, 5)
			w2 := newWord(data.Noun, 5)
			w3 := newWord(data.Verb, 6)
			w4 := newWord(data.Noun, l-16)
			structure = append(structure, w1, w2, w3, w4)
		} else if !ranN1 && !ranN2 && ranN3 {
			w1 := newWord(data.Adjective, 5)
			w2 := newWord(data.Noun, 6)
			w3 := newWord(data.Verb, 5)
			w4 := newWord(data.Noun, l-16)
			structure = append(structure, w1, w2, w3, w4)
		} else if !ranN1 && ranN2 && !ranN3 {
			w1 := newWord(data.Noun, 6)
			w2 := newWord(data.Verb, 6)
			w3 := newWord(data.Adjective, 5)
			w4 := newWord(data.Noun, l-17)
			structure = append(structure, w1, w2, w3, w4)
		} else if ranN1 && !ranN2 && !ranN3 {
			w1 := newWord(data.Noun, 5)
			w2 := newWord(data.Verb, 6)
			w3 := newWord(data.Adjective, 6)
			w4 := newWord(data.Noun, l-17)
			structure = append(structure, w1, w2, w3, w4)
		} else {
			w1 := newWord(data.Adjective, 6)
			w2 := newWord(data.Noun, 6)
			w3 := newWord(data.Adverb, 5)
			w4 := newWord(data.Verb, l-17)
			structure = append(structure, w1, w2, w3, w4)
		}

	// (26-32)
	default:

		// random factor
		if ranN1 && ranN2 && ranN3 {
			w1 := newWord(data.Adjective, 6)
			w2 := newWord(data.Noun, 7)
			w3 := newWord(data.Adverb, 6)
			w4 := newWord(data.Verb, 4)
			w5 := newWord(data.Noun, l-23)
			structure = append(structure, w1, w2, w3, w4, w5)
		} else if !ranN1 && ranN2 && ranN3 {
			w1 := newWord(data.Adjective, 7)
			w2 := newWord(data.Noun, 6)
			w3 := newWord(data.Adverb, 6)
			w4 := newWord(data.Verb, 4)
			w5 := newWord(data.Noun, l-23)
			structure = append(structure, w1, w2, w3, w4, w5)
		} else if ranN1 && !ranN2 && ranN3 {
			w1 := newWord(data.Adjective, 5)
			w2 := newWord(data.Noun, 6)
			w3 := newWord(data.Adverb, 6)
			w4 := newWord(data.Verb, 6)
			w5 := newWord(data.Noun, l-23)
			structure = append(structure, w1, w2, w3, w4, w5)
		} else if ranN1 && ranN2 && !ranN3 {
			w1 := newWord(data.Adjective, 5)
			w2 := newWord(data.Noun, 6)
			w3 := newWord(data.Adverb, 6)
			w4 := newWord(data.Verb, 5)
			w5 := newWord(data.Noun, l-22)
			structure = append(structure, w1, w2, w3, w4, w5)
		} else if !ranN1 && !ranN2 && ranN3 {
			w1 := newWord(data.Noun, 7)
			w2 := newWord(data.Adverb, 6)
			w3 := newWord(data.Verb, 4)
			w4 := newWord(data.Adjective, 6)
			w5 := newWord(data.Noun, l-23)
			structure = append(structure, w1, w2, w3, w4, w5)
		} else if !ranN1 && ranN2 && !ranN3 {
			w1 := newWord(data.Noun, 7)
			w2 := newWord(data.Adverb, 6)
			w3 := newWord(data.Verb, 4)
			w4 := newWord(data.Adjective, 6)
			w5 := newWord(data.Noun, l-23)
			structure = append(structure, w1, w2, w3, w4, w5)
		} else if ranN1 && !ranN2 && !ranN3 {
			w1 := newWord(data.Noun, 4)
			w2 := newWord(data.Adverb, 5)
			w3 := newWord(data.Verb, 6)
			w4 := newWord(data.Adjective, 6)
			w5 := newWord(data.Noun, l-21)
			structure = append(structure, w1, w2, w3, w4, w5)
		} else {
			w1 := newWord(data.Noun, 7)
			w2 := newWord(data.Adverb, 5)
			w3 := newWord(data.Verb, 4)
			w4 := newWord(data.Adjective, 6)
			w5 := newWord(data.Noun, l-22)
			structure = append(structure, w1, w2, w3, w4, w5)
		}
	}

	return structure, nil
}
