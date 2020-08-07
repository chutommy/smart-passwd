package controls

// extraSecurityLvl return extra number of numbers and special characters.
func extraSecurityLvl(lvl int) (int, int) {

	// narrow
	if lvl < 0 {
		lvl = 0
	} else if lvl > 10 {
		lvl = 10
	}

	// select
	switch lvl {
	case 0:
		return 0, 0
	case 1:
		return 1, 0
	case 2:
		return 2, 0
	case 3:
		return 2, 1
	case 4:
		return 3, 1
	case 5:
		return 3, 2
	case 6:
		return 4, 2
	case 7:
		return 4, 3
	case 8:
		return 5, 3
	case 9:
		return 6, 3
	default:
		return 6, 4
	}
}
