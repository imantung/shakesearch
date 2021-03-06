package app_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func TestCreateBook(t *testing.T) {
	testcases := []struct {
		Name         string
		Text         string
		Meta         *app.Meta
		PreviewLimit int
		Expected     *app.Book
		ExpectedErr  string
	}{
		{
			Name: "no chapters",
			Text: "12345\n123456789\n1234",
			Meta: &app.Meta{},
			Expected: &app.Book{
				Text:     "12345\n123456789\n1234\n",
				LineIdxs: []int{0, 6, 16},
			},
		},
		{
			Name: "with chapters",
			Text: "title1\nabcdefghij\nklmnopqrstuv\nwxyz\ntitle2\n12345\n67890\ntitle3\nabcdef\nghijklm\n",
			Meta: &app.Meta{
				Chapters: []string{"title1", "title2", "title3"},
			},
			PreviewLimit: 120,
			Expected: &app.Book{
				Text:         "title1\nabcdefghij\nklmnopqrstuv\nwxyz\ntitle2\n12345\n67890\ntitle3\nabcdef\nghijklm\n",
				LineIdxs:     []int{0, 7, 18, 31, 36, 43, 49, 55, 62, 69},
				Chapters:     []string{"title1", "title2", "title3"},
				ChapterIdxs:  []int{0, 36, 55},
				PreviewLimit: 120,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.Name, func(t *testing.T) {
			require.Equal(t, tt.Expected, app.NewBook(strings.NewReader(tt.Text), tt.Meta, tt.PreviewLimit))
		})
	}
}

func TestBook_Retrieve(t *testing.T) {
	book := &app.Book{
		Text:         "title1\nabcdefghij\nklmnopqrstuv\nwxyz\ntitle2\n12345\n67890\ntitle3\nabcdef\nghijklm\n",
		LineIdxs:     []int{0, 7, 18, 31, 36, 43, 49, 55, 62, 69, 77},
		Chapters:     []string{"title1", "title2", "title3"},
		ChapterIdxs:  []int{0, 36, 55},
		PreviewLimit: 5,
	}

	testcases := []struct {
		Name     string
		Book     *app.Book
		Idx      int
		Expected *app.Result
	}{
		{
			Name:     "first line",
			Book:     book,
			Idx:      5,
			Expected: &app.Result{Preview: "le1\na", Chapter: "title1", LineNumber: 1},
		},
		{
			Name:     "random line",
			Book:     book,
			Idx:      10,
			Expected: &app.Result{Preview: "bcdef", Chapter: "title1", LineNumber: 2},
		},
		{
			Name:     "random line",
			Book:     book,
			Idx:      54,
			Expected: &app.Result{Preview: "90\nti", Chapter: "title2", LineNumber: 7},
		},
		{
			Name:     "last line",
			Book:     book,
			Idx:      75,
			Expected: &app.Result{Preview: "klm\n", Chapter: "title3", LineNumber: 10},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.Name, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Book.Retrieve(tt.Idx), tt.Name)
		})
	}
}
