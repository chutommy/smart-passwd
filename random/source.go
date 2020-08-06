package random

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

// Source is custom source for the random generator.
type Source struct{}

// Seed has no operations, need to satisfy rand.Source interface.
func (*Source) Seed(int64) { /* noop */ }

// Uint64 returns random uint64.
func (s *Source) Uint64() uint64 {
	var value uint64
	binary.Read(crand.Reader, binary.BigEndian, &value)
	return value
}

// Int63 returns random int64.
func (s *Source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

// GetRNG returns truly random math.Rand.
func GetRNG() *mrand.Rand {
	return mrand.New(&Source{})
}
