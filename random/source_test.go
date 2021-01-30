package random

import "testing"

func TestSource(t *testing.T) {
	r := GetRNG()
	r.Seed(0)
	_ = r.Int()
}
