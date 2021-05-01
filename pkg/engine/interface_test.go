package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	t.Parallel()

	length := int16(10)
	extraSec := int16(5)
	helper := "foobar"

	req1 := NewRequest(length, extraSec, "")
	req2 := NewRequest(0, extraSec, helper)

	require.Equal(t, length, req1.Length())
	require.Equal(t, extraSec, req1.ExtraSec())
	require.Equal(t, "", req1.Helper())

	require.Equal(t, int16(0), req2.Length())
	require.Equal(t, extraSec, req2.ExtraSec())
	require.Equal(t, helper, req2.Helper())
}

func TestResponse(t *testing.T) {
	t.Parallel()

	password := "f0o6Ar123@" //nolint:gosec
	helper := "foobar"

	resp := NewResponse(password, helper)

	require.Equal(t, password, resp.Password())
	require.Equal(t, helper, resp.Helper())
}
