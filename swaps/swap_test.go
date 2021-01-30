package swaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapEngine(t *testing.T) {
	se := New()

	r1 := se.GetAlpha(2)
	r2 := se.GetAlphaCap(2)
	r3 := se.GetNum(2)
	r4 := se.GetSpecial(1000)
	r5 := se.GetSubst('a', 1000)

	assert.NotEqual(t, r1, 'a')
	assert.NotEqual(t, r2, 'A')
	assert.NotEqual(t, r3, '1')
	assert.NotEqual(t, r4, 'a')
	assert.NotEqual(t, r5, 'a')
}
