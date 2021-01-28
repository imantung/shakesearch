package app

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type (
	// Book is pragmatic data for the complete-book
	Book struct {
		Source      string
		Text        string
		Eols        []int // End of line indexes
		Chapters    []Chapter
		ChapterIdxs []int
	}
	// Chapter information
	Chapter struct {
		Name string
		Idx  int
	}
)

// CreateBook create pragramatic data
func CreateBook(source string, chapterTitles []string) (*Book, error) {
	file, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		buf         bytes.Buffer
		eols        []int
		curr        int
		chapterIdxs []int
		chapters    []Chapter
	)
	chapterMap := stringMap(chapterTitles)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(&buf, text)
		curr += len(text)
		eols = append(eols, curr)
		if _, ok := chapterMap[text]; ok {
			chapters = append(chapters, Chapter{Name: text, Idx: curr})
			chapterIdxs = append(chapterIdxs, curr)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Book{
		Source:      source,
		Text:        buf.String(),
		Eols:        eols,
		Chapters:    chapters,
		ChapterIdxs: chapterIdxs,
	}, nil
}

func stringMap(slice []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, s := range slice {
		m[s] = struct{}{}
	}
	return m
}
