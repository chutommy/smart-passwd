package engine

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
	if l < 3 {
		return nil, fmt.Errorf("%w: %d", ErrInvalidLength, l)
	}

	var ll []int16

	// sokut l into ll
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

			// check ramainder
			if d := l - i; d < c.min && d != 0 {
				continue
			}

			break
		}

		// append
		l -= i
		ll = append(ll, i)
	}

	return ll, nil
}
