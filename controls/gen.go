package controls

import "fmt"

func (c *Controller) Gen(l int) {

	fmt.Println(c.newPhrase(l))
}
