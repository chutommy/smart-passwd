package engine

// Alphabet returns an array of runes of letters.
func Alphabet() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
}

// Specials returns an array of runes of special symbols.
func Specials() []rune {
	return []rune{'!', '@', '#', '$', '%', '^', '&', '*', '?', '+', '_', '-', '(', ')'}
}

// SwapList returns a map of runes linked with similar looking characters.
func SwapList() map[rune][]rune {
	return map[rune][]rune{
		'a': {'@', '&', '4'},
		'b': {'6', '&', 'G', 'd'},
		'c': {'e'},
		'd': {'b'},
		'e': {'c'},
		'f': {},
		'g': {'9'},
		'h': {},
		'i': {'j', 'l', '1', '!', 'y'},
		'j': {'i', '!', '1'},
		'k': {},
		'l': {'I', '!'},
		'm': {},
		'n': {},
		'o': {'O', '0', '@'},
		'p': {'q'},
		'q': {'p'},
		'r': {},
		's': {'$'},
		't': {},
		'u': {},
		'v': {'w'},
		'w': {'v'},
		'x': {'+'},
		'y': {'i'},
		'z': {},

		'A': {'4', '@', '&'},
		'B': {'8', '%'},
		'C': {'('},
		'D': {'0', 'O'},
		'E': {},
		'F': {},
		'G': {'6'},
		'H': {},
		'I': {'l', '1', '!'},
		'J': {},
		'K': {},
		'L': {},
		'M': {},
		'N': {},
		'O': {'o', '0', '@', 'D'},
		'P': {},
		'Q': {},
		'R': {},
		'S': {},
		'T': {'7'},
		'U': {},
		'V': {'W'},
		'W': {'V'},
		'X': {'+'},
		'Y': {'I'},
		'Z': {},
	}
}
