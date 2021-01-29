package app_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func TestCreateBook(t *testing.T) {
	testcases := []struct {
		Name         string
		Text         string
		Chapters     []string
		PreviewLimit int
		Expected     *app.Book
		ExpectedErr  string
	}{
		{
			Name: "no chapters",
			Text: "12345\n123456789\n1234",
			Expected: &app.Book{
				Source:   "filename",
				Text:     "12345\n123456789\n1234\n",
				LineIdxs: []int{0, 6, 16, 21},
			},
		},
		{
			Name:         "with chapters",
			Text:         "title1\nabcdefghij\nklmnopqrstuv\nwxyz\ntitle2\n12345\n67890\ntitle3\nabcdef\nghijklm\n",
			Chapters:     []string{"title1", "title2", "title3"},
			PreviewLimit: 120,
			Expected: &app.Book{
				Source:       "filename",
				Text:         "title1\nabcdefghij\nklmnopqrstuv\nwxyz\ntitle2\n12345\n67890\ntitle3\nabcdef\nghijklm\n",
				LineIdxs:     []int{0, 7, 13, 23, 28, 35, 46, 54, 61, 67, 73},
				Chapters:     []string{"title1", "title2", "title3"},
				ChapterIdxs:  []int{0, 28, 54},
				PreviewLimit: 120,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.Name, func(t *testing.T) {
			require.NoError(t, ioutil.WriteFile("filename", []byte(tt.Text), 0777))
			defer os.Remove("filename")

			book, err := app.CreateBook("filename", tt.Chapters, tt.PreviewLimit)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.Expected, book)
			}
		})
	}
}

func TestBook_Retrieve(t *testing.T) {
	book := &app.Book{
		Source:       "filename",
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
