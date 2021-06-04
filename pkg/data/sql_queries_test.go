package data

import (
	"errors"
	"fmt"
	"testing"

	"github.com/chutified/smart-passwd/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestRandomWord(t *testing.T) {
	t.Parallel()

	t.Run("nil db", func(t *testing.T) {
		t.Parallel()

		_, err := (&SQLiteWordList{db: nil}).randomWord(8)
		require.True(t, errors.Is(err, utils.ErrNilValue))
	})

	t.Run("no words", func(t *testing.T) {
		t.Parallel()

		_, err := testSQLWordList.randomWord(0)
		require.Error(t, err)

		_, err = testSQLWordList.randomWord(25)
		require.Error(t, err)
	})

	for i := int16(1); i <= 22; i++ {
		i := i
		t.Run(fmt.Sprintf("len-%d", i), func(t *testing.T) {
			t.Parallel()

			w1, err1 := testSQLWordList.randomWord(i)
			w2, err2 := testSQLWordList.randomWord(i)

			require.NoError(t, err1)
			require.NoError(t, err2)

			require.NotEmpty(t, w1)
			require.NotEmpty(t, w2)

			var ok bool
			for j := 0; j < 800; j++ {
				if w1 != w2 {
					ok = true

					break
				}

				w2, err2 = testSQLWordList.randomWord(i)
				require.NoError(t, err2)
				require.NotEmpty(t, w2)
			}
			require.True(t, ok, "failed to generate two different words")
		})
	}
}
