package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func TestSuffixArraySearcher_Retrieve(t *testing.T) {
	testcases := []struct {
		Name     string
		Searcher app.Searcher
		Query    string
		Expected []string
	}{
		{
			Name:     "preview length more than data",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn`), 500),
			Query:    "defgh",
			Expected: []string{"abcdefghijklmn"},
		},
		{
			Name:     "query in first text",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn`), 500),
			Query:    "abc",
			Expected: []string{"abcdefghijklmn"},
		},
		{
			Name:     "query in last text",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn`), 500),
			Query:    "klmn",
			Expected: []string{"abcdefghijklmn"},
		},
		{
			Name:     "multiple occurance",
			Searcher: app.NewSuffixArraySearcher([]byte(`abcdefghijklmn defghijklm defghijklm`), 5),
			Query:    "de",
			Expected: []string{"m def", "n def", "bcdef"},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.Name, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Searcher.Search(tt.Query))
		})
	}
}
