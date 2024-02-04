package aho

import (
	"github.com/cloudflare/ahocorasick"
	"strings"
)

// AHO https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
type AHO struct {
	ac *ahocorasick.Matcher
}

func NewAHO(dict []string) *AHO {
	aho := &AHO{
		ac: ahocorasick.NewStringMatcher(dict),
	}
	return aho
}

func (a AHO) Hits(text string) []int {
	return a.ac.MatchThreadSafe([]byte(text))
}

func (a AHO) Contains(text string) bool {
	return a.ac.Contains([]byte(text))
}

type Normal struct {
	dict []string
}

func NewNormal(dict []string) *Normal {
	return &Normal{dict: dict}
}

func (n Normal) Contains(text string) bool {
	for _, d := range n.dict {
		if strings.Contains(text, d) {
			return true
		}
	}
	return false
}
