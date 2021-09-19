package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTopFrequentWords(t *testing.T) {
	require.Equal(
		t,
		[10]string{"op", "ui", "as", "ty", "qw", "lp", "kl", "jk", "gh", "er"},
		TopFrequentWords("as df gh jk kl lp qw er ty ui op op op op ui ui as"),
	)
	require.Equal(
		t,
		[10]string{"op", "ui", "as"},
		TopFrequentWords("as ui op op op op ui ui as"),
	)
}
