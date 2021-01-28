package app

import (
	"fmt"
	"index/suffixarray"
	"io/ioutil"
)

type (
	// Searcher is interface for search
	Searcher interface {
		Search(query string) []string
	}
	// SuffixArraySearcher is search engine with substring search
	// NOTE: original implementation
	SuffixArraySearcher struct {
		CompleteWorks string
		SuffixArray   *suffixarray.Index
	}
)

//
// SuffixArraySearcher
//

var _ Searcher = (*SuffixArraySearcher)(nil)

// NewSuffixArraySearcher return new instance of Substring searcher
func NewSuffixArraySearcher(filename string) (Searcher, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Substring-Searcher: %w", err)
	}
	return &SuffixArraySearcher{
		CompleteWorks: string(dat),
		SuffixArray:   suffixarray.New(dat),
	}, nil
}

// Search ...
func (s *SuffixArraySearcher) Search(query string) []string {
	idxs := s.SuffixArray.Lookup([]byte(query), -1)
	results := []string{}
	for _, idx := range idxs {
		results = append(results, s.CompleteWorks[idx-250:idx+250])
	}
	return results
}
