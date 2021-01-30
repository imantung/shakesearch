package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type (
	// Book is pragmatic data for the complete-book
	Book struct {
		Text         string
		LineIdxs     []int // Beginning of line indexes
		Chapters     []string
		ChapterIdxs  []int
		PreviewLimit int
	}
	// Meta data
	Meta struct {
		Chapters []string `json:"chapters"`
	}
)

// NewBook create pragramatic data
func NewBook(r io.Reader, meta *Meta, previewLimit int) *Book {
	var (
		buf         bytes.Buffer
		lineIdxs    []int
		curr        int
		chapterIdxs []int
		chapters    []string
	)
	chapterMap := chapterMap(meta)
	scanner := bufio.NewScanner(r)
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

	return &Book{
		Text:         buf.String(),
		LineIdxs:     lineIdxs,
		Chapters:     chapters,
		ChapterIdxs:  chapterIdxs,
		PreviewLimit: previewLimit,
	}
}

func chapterMap(meta *Meta) map[string]struct{} {
	m := make(map[string]struct{})
	for _, s := range meta.Chapters {
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
