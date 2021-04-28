package engine

import (
	"math/rand"
	"unicode"

	"github.com/chutified/smart-passwd/pkg/utils"
)

// Swapper represents a random swapper engine which can also generate
// a truly random letter, number or special character.
type Swapper struct {
	rand     *rand.Rand
	alpha    []rune
	special  []rune
	swapList map[rune][]rune
}

// NewSwapper is a constructor for the Swapper. It populates the alphabet,
// special symbol list and the map swapList with the given values.
func NewSwapper(alpha []rune, special []rune, swap map[rune][]rune) *Swapper {
	return &Swapper{
		rand:     utils.Rand(),
		alpha:    alpha,
		special:  special,
		swapList: swap,
	}
}

// Alpha returns a random lower-case letter.
func (s *Swapper) Alpha() rune {
	i := s.rand.Intn(len(s.alpha))

	return s.alpha[i]
}

// AlphaCap returns a random upper-case letter.
func (s *Swapper) AlphaCap() rune {
	i := s.rand.Intn(len(s.alpha))

	return unicode.ToUpper(s.alpha[i])
}

// Special returns a random special symbol.
func (s *Swapper) Special() rune {
	i := s.rand.Intn(len(s.special))

	return s.special[i]
}

// Num returns a random number.
func (s *Swapper) Num() int16 {
	return int16(s.rand.Intn(10))
}

// Swap returns an alternative character to the given one if possible.
// If there are no alternatives, it returns the same character.
func (s *Swapper) Swap(char rune) rune {
	l := len(s.swapList[char])
	if l == 0 {
		return char
	}

	i := s.rand.Intn(l)

	return s.swapList[char][i]
}