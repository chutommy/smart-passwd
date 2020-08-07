package controls

// randomAdds adds extra numbers and special symbols to the string.
func (c *Controller) randomAdds(s string, nums int, specials int) string {

	// prepare the result
	result := []rune(s)

	// add nums
	for i := 0; i < nums; i++ {

		// generate position and character
		pos := c.rng.Intn(len(result))
		ch := c.se.GetNum(c.rng.Int())

		// insert and shift
		temp := append(result[:pos+1], result[pos:]...)
		temp[pos] = ch
		result = temp
	}

	// add specials
	for i := 0; i < specials; i++ {

		// generate position and character
		pos := c.rng.Intn(len(result))
		ch := c.se.GetSpecial(c.rng.Int())

		// insert and shift
		temp := append(result[:pos+1], result[pos:]...)
		temp[pos] = ch
		result = temp
	}

	return string(result)
}
