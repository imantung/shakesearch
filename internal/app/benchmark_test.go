package app_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func BenchmarkSuffixArraySearcher(b *testing.B) {
	data, err := ioutil.ReadFile("../../completeworks.txt")
	require.NoError(b, err)

	searcher := app.NewSuffixArraySearcher(data, 500)
	for i := 0; i < b.N; i++ {
		searcher.Search("Hamlet")
	}
}
