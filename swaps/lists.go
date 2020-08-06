package swaps

var (
	// alpha defines the default alphabet.
	alpha = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	// alphaCap defines the capitalized alphabet.
	alphaCap = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	// num holds the list of the numer characters.
	num = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// special holds valid specials characters.
	special = []byte{'!', '@', '#', '$', '%', '&', '*', '?'}

	// swap is the hash table of the swappable characters (look or sound similarly).
	swap = map[byte][]byte{
		'a': []byte{'@', '&', '4'},
		'b': []byte{'6', '&'},
		'i': []byte{'j', 'l', '1', '!'},
		'j': []byte{'i', '!'},
		'l': []byte{'I', '!'},
		'o': []byte{'O', '0', '@'},
		'p': []byte{'q'},
		'q': []byte{'p'},
		's': []byte{'$'},
		'v': []byte{'w'},
		'w': []byte{'v'},
		'B': []byte{'8', '%'},
		'I': []byte{'l', 'l'},
		'O': []byte{'o', '0', '@'},
		'V': []byte{'W'},
		'W': []byte{'V'},
	}
)
