package controls

import "unicode"

// transform smartly transforms the phrase into another string, but
// keeps the look or the sound of the phrase.
func (c *Controller) transform(s string) string {

	// prepare the result
	result := make([]rune, len(s))

	// range over each character
	for i, ch := range s {

		// capitalize
		if c.rng.Intn(3) == 0 {
			ch = unicode.ToUpper(ch)
		}

		//swap
		if c.rng.Intn(2) == 0 {
			ch = c.se.GetSubst(ch, c.rng.Int())
		}

		// add to the result
		result[i] = ch
	}

	return string(result)
}
