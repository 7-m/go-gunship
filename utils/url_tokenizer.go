package utils

import "strings"

type UrlTokenizer struct {
	pattern    string
	pathAtoms  []string
	queryAtoms []string
}

func NewUrlTokenzier(pattern string) *UrlTokenizer {

	pathAtoms := strings.Split(pattern[1:], "/")
	size := len(pathAtoms)

	// query handling
	last := pathAtoms[size-1]

	var queryAtoms []string
	if strings.ContainsRune(last, '?') {
		split := strings.Split(last, "?")
		// return tha path back
		pathAtoms[size-1] = split[0]
		queryAtoms = strings.Split(split[1], "&")
	}
	return &UrlTokenizer{
		pattern:    pattern,
		pathAtoms:  pathAtoms,
		queryAtoms: queryAtoms,
	}

}
func (s *UrlTokenizer) Pattern() string {
	return s.pattern
}

func (s *UrlTokenizer) PathAtoms() []string {
	return s.pathAtoms
}

func (s *UrlTokenizer) QueryAtoms() []string {
	return s.queryAtoms
}



