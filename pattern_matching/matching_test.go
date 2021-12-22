package pattern

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatching(t *testing.T) {
	require.Equal(t, false, IsMatch("aa", "a"))
	require.Equal(t, false, IsMatch("a", "*"))
	require.Equal(t, true, IsMatch("a", "."))
	require.Equal(t, false, IsMatch("babbo", "ba*z"))
	require.Equal(t, true, IsMatch("baaaaaaaaaaaaaaaaao", "ba*o"))

}
