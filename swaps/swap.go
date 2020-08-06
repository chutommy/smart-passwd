package swaps

// SwapEngine is the swapping controller.
type SwapEngine struct{}

// New is a contructor for the SwapEngine.
func New() *SwapEngine {
	return &SwapEngine{}
}

// GetAlpha returns the i-th alphabet character.
func (se *SwapEngine) GetAlpha(i int) byte {
	i %= 26
	return alpha[i]
}

// GetAlphaCap returns the i-th caputalized alphabet character.
func (se *SwapEngine) GetAlphaCap(i int) byte {
	i %= 26
	return alphaCap[i]
}

// GetNum returns the i-th integer.
func (se *SwapEngine) GetNum(i int) byte {
	i %= 10
	return num[i]
}

// GetSpecial returns the i-th special character.
func (se *SwapEngine) GetSpecial(i int) byte {
	i %= len(special)
	return special[i]
}

// GetSubst return the posible substitition for the given char.
// If no substitution found, returns the same char and the false.
func (se *SwapEngine) GetSubst(char byte, i int) (byte, bool) {
	if subst, ok := swap[char]; ok {
		i %= len(swap)
		return subst[i], true
	}
	return char, false
}
