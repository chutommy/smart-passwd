package swaps

var (
	// alpha defines the default alphabet.
	alpha = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	// alphaCap defines the capitalized alphabet.
	alphaCap = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	// num holds the list of the number characters.
	num = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// special holds valid specials characters.
	special = []rune{'!', '@', '#', '$', '%', '&', '*', '?', '+', '_', '-'}

	// swap is the hash table of the swappable characters (look or sound similarly).
	swap = map[rune][]rune{
		'a': {'@', '&', '4'},
		'b': {'6', '&', 'G', 'd'},
		'c': {'e'},
		'd': {'b'},
		'e': {'c'},
		'g': {'9'},
		'i': {'j', 'l', '1', '!', 'y'},
		'j': {'i', '!', '1'},
		'l': {'I', '!'},
		'o': {'O', '0', '@'},
		'p': {'q'},
		'q': {'p'},
		's': {'$'},
		'v': {'w'},
		'w': {'v'},
		'x': {'+'},
		'y': {'i'},

		'A': {'4', '@', '&'},
		'B': {'8', '%'},
		'D': {'0', 'O'},
		'G': {'6'},
		'I': {'l', '1', '!'},
		'O': {'o', '0', '@', 'D'},
		'T': {'7'},
		'V': {'W'},
		'W': {'V'},
		'X': {'+'},
		'Y': {'I'},
	}
)
