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
			Text: "12345\n123456789\n1234",
			Expected: &app.Book{
				Source: "filename",
				Text:   "12345\n123456789\n1234\n",
				Eols:   []int{5, 14, 18},
			},
		},
		{
			Text:         "title1\n12345\n123456789\n1234\ntitle2\n3456023834\n0122349\ntitle3\n12434\n94523",
			Chapters:     []string{"title1", "title2", "title3"},
			PreviewLimit: 120,
			Expected: &app.Book{
				Source: "filename",
				Text:   "title1\n12345\n123456789\n1234\ntitle2\n3456023834\n0122349\ntitle3\n12434\n94523\n",
				Eols:   []int{6, 11, 20, 24, 30, 40, 47, 53, 58, 63},
				Chapters: []app.Chapter{
					{Name: "title1", Idx: 6},
					{Name: "title2", Idx: 30},
					{Name: "title3", Idx: 53},
				},
				ChapterIdxs:  []int{6, 30, 53},
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
