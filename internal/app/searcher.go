package app

import (
	"index/suffixarray"
)

type (
	// Searcher is interface for search
	Searcher interface {
		Search(query string) []*Result
	}
	// Result of search
	Result struct {
		Preview    string `json:"preview"`
		Chapter    string `json:"chapter"`
		LineNumber int    `json:"line_number"`
	}
	// SuffixArraySearcher substring search in logarithmic time using an in-memory suffix array
	SuffixArraySearcher struct {
		Book        *Book
		SuffixArray *suffixarray.Index
	}
)

//
// SuffixArraySearcher
//

var _ Searcher = (*SuffixArraySearcher)(nil)

// NewSuffixArraySearcher return new instance of Substring searcher
func NewSuffixArraySearcher(book *Book) Searcher {
	return &SuffixArraySearcher{
		Book:        book,
		SuffixArray: suffixarray.New([]byte(book.Text)),
	}
}

// Search ...
func (s *SuffixArraySearcher) Search(q string) []*Result {
	idxs := s.SuffixArray.Lookup([]byte(q), -1)
	var results []*Result
	for _, idx := range idxs {
		results = append(results, s.Book.Retrieve(idx))
	}
	return results
}
