package app

import (
	"index/suffixarray"
)

type (
	// Searcher is interface for search
	Searcher interface {
		Search(query string) []string
	}
	// SuffixArraySearcher substring search in logarithmic time using an in-memory suffix array
	SuffixArraySearcher struct {
		FullText     string
		SuffixArray  *suffixarray.Index
		PreviewLimit int
	}
)

//
// SuffixArraySearcher
//

var _ Searcher = (*SuffixArraySearcher)(nil)

// NewSuffixArraySearcher return new instance of Substring searcher
func NewSuffixArraySearcher(data []byte, previewLimit int) Searcher {
	return &SuffixArraySearcher{
		FullText:     string(data),
		SuffixArray:  suffixarray.New(data),
		PreviewLimit: previewLimit,
	}
}

// Search ...
func (s *SuffixArraySearcher) Search(q string) []string {
	idxs := s.SuffixArray.Lookup([]byte(q), -1)
	results := []string{}
	for _, idx := range idxs {
		results = append(results, s.Retrieve(idx))
	}
	return results
}

// Retrieve search result in specific idx
func (s *SuffixArraySearcher) Retrieve(idx int) string {
	begin := idx - s.PreviewLimit/2
	if begin < 0 {
		begin = 0
	}
	end := begin + s.PreviewLimit
	length := len(s.FullText)
	if end > length {
		end = length
	}
	return s.FullText[begin:end]
}
