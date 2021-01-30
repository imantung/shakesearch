package app_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"pulley.com/shakesearch/internal/app"
)

func TestCreateSearcher(t *testing.T) {
	ioutil.WriteFile("some-text-source", []byte(`Lorem ipsum dolor sit integer`), 0777)
	ioutil.WriteFile("some-meta-source", []byte(`{}`), 0777)
	defer os.Remove("some-text-source")
	defer os.Remove("some-meta-source")

	searcher, err := app.CreateSearcher(app.Config{
		PreviewLimit: 500,
		TextSource:   "some-text-source",
		MetaSource:   "some-meta-source",
	})
	require.NoError(t, err)
	require.Equal(t, []*app.Result{
		{Preview: "Lorem ipsum dolor sit integer\n", Chapter: "unknown", LineNumber: 1},
	}, searcher.Search("ipsum"))
}

func TestCreateSearche_MissingTextSource(t *testing.T) {
	ioutil.WriteFile("some-meta-source", []byte(`{}`), 0777)
	defer os.Remove("some-meta-source")

	_, err := app.CreateSearcher(app.Config{
		PreviewLimit: 500,
		TextSource:   "some-text-source",
		MetaSource:   "some-meta-source",
	})
	require.EqualError(t, err, "open some-text-source: no such file or directory")
}

func TestCreateSearcher_MissingMetaSource(t *testing.T) {
	ioutil.WriteFile("some-text-source", []byte(`Lorem ipsum dolor sit integer`), 0777)
	defer os.Remove("some-text-source")

	_, err := app.CreateSearcher(app.Config{
		PreviewLimit: 500,
		TextSource:   "some-text-source",
		MetaSource:   "some-meta-source",
	})
	require.EqualError(t, err, "open some-meta-source: no such file or directory")
}

func TestCreateSearcher_BadMetaSource(t *testing.T) {
	ioutil.WriteFile("some-text-source", []byte(`Lorem ipsum dolor sit integer`), 0777)
	ioutil.WriteFile("some-meta-source", []byte(`{bad-json`), 0777)
	defer os.Remove("some-text-source")
	defer os.Remove("some-meta-source")

	_, err := app.CreateSearcher(app.Config{
		PreviewLimit: 500,
		TextSource:   "some-text-source",
		MetaSource:   "some-meta-source",
	})
	require.EqualError(t, err, "invalid character 'b' looking for beginning of object key string")
}
