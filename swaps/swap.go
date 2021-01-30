package swaps

// SwapEngine is the swapping controller.
type SwapEngine struct{}

// New is a constructor for the SwapEngine.
func New() *SwapEngine {
	return &SwapEngine{}
}

// GetAlpha returns the i-th alphabet character.
func (se *SwapEngine) GetAlpha(i int) rune {
	i %= 26
	return alpha[i]
}

// GetAlphaCap returns the i-th capitalized alphabet character.
func (se *SwapEngine) GetAlphaCap(i int) rune {
	i %= 26

	return alphaCap[i]
}

// GetNum returns the i-th integer.
func (se *SwapEngine) GetNum(i int) rune {
	i %= 10

	return num[i]
}

// GetSpecial returns the i-th special character.
func (se *SwapEngine) GetSpecial(i int) rune {
	i %= len(special)

	return special[i]
}

// GetSubst return the possible substitution for the given char.
// If no substitution found, returns the same char.
func (se *SwapEngine) GetSubst(char rune, i int) rune {
	if subst, ok := swap[char]; ok {
		i %= len(subst)

		return subst[i]
	}

	return char
}
