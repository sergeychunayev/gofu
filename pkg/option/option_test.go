package option

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOption(t *testing.T) {
	testCases := []struct {
		name  string
		input Option[int]
		check func(t *testing.T, res Option[int])
	}{
		{
			"Some",
			Of[int](1),
			func(t *testing.T, res Option[int]) {
				require.True(t, res.IsSome())
				require.False(t, res.IsNone())
				require.Equal(t, 1, res.Unwrap())
				require.Equal(t, 1, res.UnwrapOr(2))
			},
		},
		{
			"None panic",
			No[int](),
			func(t *testing.T, res Option[int]) {
				defer func() {
					if r := recover(); r != nil {
						require.Equal(t, "None", r)
					} else {
						require.Fail(t, "Must panic")
					}
				}()
				require.True(t, res.IsNone())
				require.False(t, res.IsSome())
				require.Equal(t, 1, res.Unwrap())
			},
		},
		{
			"None",
			No[int](),
			func(t *testing.T, res Option[int]) {
				require.True(t, res.IsNone())
				require.Equal(t, 2, res.UnwrapOr(2))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.check(t, tc.input)
		})
	}
}
