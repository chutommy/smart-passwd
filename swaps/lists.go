package swaps

var (
	// alpha defines the default alphabet.
	alpha = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	// alphaCap defines the capitalized alphabet.
	alphaCap = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	// num holds the list of the numer characters.
	num = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// special holds valid specials characters.
	special = []rune{'!', '@', '#', '$', '%', '&', '*', '?', '+'}

	// swap is the hash table of the swappable characters (look or sound similarly).
	swap = map[rune][]rune{
		'a': []rune{'@', '&', '4'},
		'b': []rune{'6', '&', 'G', 'd'},
		'c': []rune{'e'},
		'd': []rune{'b'},
		'e': []rune{'c'},
		'g': []rune{'9'},
		'i': []rune{'j', 'l', '1', '!', 'y'},
		'j': []rune{'i', '!', '1'},
		'l': []rune{'I', '!'},
		'o': []rune{'O', '0', '@'},
		'p': []rune{'q'},
		'q': []rune{'p'},
		's': []rune{'$'},
		'v': []rune{'w'},
		'w': []rune{'v'},
		'x': []rune{'+'},
		'y': []rune{'i'},

		'A': []rune{'4', '@', '&'},
		'B': []rune{'8', '%'},
		'D': []rune{'0', 'O'},
		'G': []rune{'6'},
		'I': []rune{'l', '1', '!'},
		'O': []rune{'o', '0', '@', 'D'},
		'T': []rune{'7'},
		'V': []rune{'W'},
		'W': []rune{'V'},
		'X': []rune{'+'},
		'Y': []rune{'I'},
	}
)
