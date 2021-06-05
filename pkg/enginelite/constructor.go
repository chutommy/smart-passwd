package enginelite

import (
	"fmt"
	"math/rand"

	"github.com/chutified/smart-passwd/pkg/utils"
	"github.com/pkg/errors"
)

// ErrInvalidLength is returned if invalid length is given.
var ErrInvalidLength = errors.New("invalid length")

// Constructor holds a scheme for splitting the length.
type Constructor struct {
	rand *rand.Rand
	min  int16
	max  int16
}

// NewConstructor is a constructor for the Constructor.
func NewConstructor(min, max int16) *Constructor {
	return &Constructor{
		rand: utils.Rand(),
		min:  min,
		max:  max,
	}
}

// Distribute distributes a length l into a slice of lengths
// which is limited with the Constructor c.
func (c *Constructor) Distribute(l int16) ([]int16, error) {
	if l < c.min {
		return nil, fmt.Errorf("%w: %d", ErrInvalidLength, l)
	}

	var ll []int16

	// split l into ll
	for l > 0 {
		var i int16

		// upper limit
		max := c.max
		if l < max {
			max = l
		}

		// evaluate i
		for {
			// generate random number in range [min,max]
			i = int16(c.rand.Intn(int(max-c.min+1))) + c.min

			// check remainder
			if d := l - i; d < c.min && d != 0 {
				continue
			}

			break
		}

		// append
		l -= i
		ll = append(ll, i)
	}

	// shuffle
	c.rand.Shuffle(len(ll), func(i, j int) {
		ll[i], ll[j] = ll[j], ll[i]
	})

	return ll, nil
}
