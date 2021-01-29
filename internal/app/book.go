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
		Source       string
		Text         string
		LineIdxs     []int // Beginning of line indexes
		Chapters     []string
		ChapterIdxs  []int
		PreviewLimit int
	}
)

// CreateBook create pragramatic data
func CreateBook(source string, chapterTitles []string, previewLimit int) (*Book, error) {
	file, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		buf         bytes.Buffer
		lineIdxs    []int
		curr        int
		chapterIdxs []int
		chapters    []string
	)
	chapterMap := stringMap(chapterTitles)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(&buf, text)
		lineIdxs = append(lineIdxs, curr)
		if _, ok := chapterMap[text]; ok {
			chapters = append(chapters, text)
			chapterIdxs = append(chapterIdxs, curr)
		}
		curr += len(text) + 1
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Book{
		Source:       source,
		Text:         buf.String(),
		LineIdxs:     lineIdxs,
		Chapters:     chapters,
		ChapterIdxs:  chapterIdxs,
		PreviewLimit: previewLimit,
	}, nil
}

func stringMap(slice []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, s := range slice {
		m[s] = struct{}{}
	}
	return m
}

// Retrieve search result in specific idx
func (b *Book) Retrieve(idx int) *Result {
	return &Result{
		Preview:    b.preview(idx),
		LineNumber: b.lineNumber(idx),
		Chapter:    b.chapter(idx),
	}
}

func (b *Book) preview(idx int) string {
	begin := idx - b.PreviewLimit/2
	if begin < 0 {
		begin = 0
	}
	end := begin + b.PreviewLimit
	length := len(b.Text)
	if end > length {
		end = length
	}
	return b.Text[begin:end]
}

func (b *Book) lineNumber(idx int) int {
	for i := len(b.LineIdxs) - 1; i >= 0; i-- {
		if idx >= b.LineIdxs[i] {
			return i + 1
		}
	}
	return -1
}

func (b *Book) chapter(idx int) string {
	for i := len(b.ChapterIdxs) - 1; i >= 0; i-- {
		if idx > b.ChapterIdxs[i] {
			return b.Chapters[i]
		}
	}
	return "unknown"
}
