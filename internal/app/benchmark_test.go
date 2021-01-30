package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func BenchmarkSuffixArraySearcher(b *testing.B) {
	searcher, err := app.CreateSearcher(app.Config{
		PreviewLimit: 500,
		TextSource:   "../../data/completeworks.txt",
		MetaSource:   "../../data/completeworks-meta.json",
	})
	require.NoError(b, err)
	for i := 0; i < b.N; i++ {
		searcher.Search("Hamlet")
	}
}
