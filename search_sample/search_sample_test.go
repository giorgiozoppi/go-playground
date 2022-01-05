package searchsample

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLoopedList(t *testing.T) {
	item := SearchSample(8)
	item2 := SearchSample(19)
	require.Equal(t, false, item)
	require.Equal(t, true, item2)
}
