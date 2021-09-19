package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPack(t *testing.T) {
	require.Equal(t, "aaaabccddddde", Unpack("a4bc2d5e"))
	require.Equal(t, "abcd", Unpack("abcd"))
	require.Equal(t, "", Unpack("45"))
	require.Equal(t, "qwe45", Unpack(`qwe\4\5`))
	require.Equal(t, "qwe44444", Unpack(`qwe\45`))
	require.Equal(t, "qwe\\\\\\\\\\", Unpack(`qwe\\5`))
}
