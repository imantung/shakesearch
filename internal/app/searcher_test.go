package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func TestSearch(t *testing.T) {
	testcases := []struct {
		Name       string
		Book       *app.Book
		SearcherFn func(*app.Book) app.Searcher
		Query      string
		Expected   []app.Result
	}{
		{
			Name: "preview length more than data",
			Book: &app.Book{
				Text:         `abcdefghijklmn`,
				PreviewLimit: 500,
			},
			SearcherFn: app.NewSuffixArraySearcher,
			Query:      "defgh",
			Expected: []app.Result{
				{Preview: "abcdefghijklmn", Chapter: "unknown", LineNumber: -1},
			},
		},
		{
			Name: "query in first text",
			Book: &app.Book{
				Text:         `abcdefghijklmn`,
				PreviewLimit: 500,
			},
			SearcherFn: app.NewSuffixArraySearcher,
			Query:      "abc",
			Expected: []app.Result{
				{Preview: "abcdefghijklmn", Chapter: "unknown", LineNumber: -1},
			},
		},
		{
			Name: "query in last text",
			Book: &app.Book{
				Text:         `abcdefghijklmn`,
				PreviewLimit: 500,
			},
			SearcherFn: app.NewSuffixArraySearcher,
			Query:      "klmn",
			Expected: []app.Result{
				{Preview: "abcdefghijklmn", Chapter: "unknown", LineNumber: -1},
			},
		},
		{
			Name: "multiple occurance",
			Book: &app.Book{
				Text:         `abcdefghijklmn defghijklm defghijklm`,
				PreviewLimit: 5,
			},
			SearcherFn: app.NewSuffixArraySearcher,
			Query:      "de",
			Expected: []app.Result{
				{Preview: "m def", Chapter: "unknown", LineNumber: -1},
				{Preview: "n def", Chapter: "unknown", LineNumber: -1},
				{Preview: "bcdef", Chapter: "unknown", LineNumber: -1},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.Name, func(t *testing.T) {
			searcher := tt.SearcherFn(tt.Book)
			require.Equal(t, tt.Expected, searcher.Search(tt.Query))
		})
	}
}
