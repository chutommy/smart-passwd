package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {
	t.Parallel()

	f := NewFile("dir", "test", "txt")
	require.Equal(t, "dir/test.txt", f.FilePath())
	require.Equal(t, "test.txt", f.FileName())
}
