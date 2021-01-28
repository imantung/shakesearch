package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func TestSearch(t *testing.T) {
	testcases := []struct {
		Name     string
		Searcher app.Searcher
		Query    string
		Expected []app.Result
	}{
		{
			Name:     "preview length more than data",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn`), 500),
			Query:    "defgh",
			Expected: []app.Result{
				{Preview: "abcdefghijklmn", Chapter: "unknown", LineNumber: -1},
			},
		},
		{
			Name:     "query in first text",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn`), 500),
			Query:    "abc",
			Expected: []app.Result{
				{Preview: "abcdefghijklmn", Chapter: "unknown", LineNumber: -1},
			},
		},
		{
			Name:     "query in last text",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn`), 500),
			Query:    "klmn",
			Expected: []app.Result{
				{Preview: "abcdefghijklmn", Chapter: "unknown", LineNumber: -1},
			},
		},
		{
			Name:     "multiple occurance",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn defghijklm defghijklm`), 5),
			Query:    "de",
			Expected: []app.Result{
				{Preview: "m def", Chapter: "unknown", LineNumber: -1},
				{Preview: "n def", Chapter: "unknown", LineNumber: -1},
				{Preview: "bcdef", Chapter: "unknown", LineNumber: -1},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.Name, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Searcher.Search(tt.Query))
		})
	}
}
