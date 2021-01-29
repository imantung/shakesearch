package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func BenchmarkSuffixArraySearcher(b *testing.B) {
	book, err := app.CreateBook("../../data/completeworks.txt", []string{}, 120)
	require.NoError(b, err)

	searcher := app.NewSuffixArraySearcher(book)
	for i := 0; i < b.N; i++ {
		searcher.Search("Hamlet")
	}
}
