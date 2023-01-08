package strmatch

import (
	"fmt"
	"regexp"
)

// Matcher
type Matcher struct {
	r []regexp.Regexp
	s []string
}

// compileRE is an internal function that compiles a string into a Regular Expression
func compileRE(s string) (*regexp.Regexp, error) {
	cre, err := regexp.Compile(s)
	if err != nil {
		return nil, err
	}
	return cre, nil
}

// containsStr returns True if a given string is a member of a given slice
func containsStr(str string, strs []string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

// NewMatcher contructs a new instance of Matcher
func NewMatcher() *Matcher {
	return &Matcher{}
}

// RAdd adds a single Regular Expression to an existing Matcher instance
func (m *Matcher) RAdd(s string) error {
	cre, err := compileRE(s)
	if err != nil {
		return fmt.Errorf("%s: Adding Regular Expression failed with: %v", s, err)
	}
	m.r = append(m.r, *cre)
	return nil
}

// SAdd adds a string to a Matcher instance
func (m *Matcher) SAdd(s string) error {
	if containsStr(s, m.s) {
		return fmt.Errorf("Matcher instance already contains string: %s", s)
	}
	m.s = append(m.s, s)
	return nil
}

// Match returns true if a given string matches either a string or regex in Matcher
func (m *Matcher) Match(s string) bool {
	if containsStr(s, m.s) {
		return true
	}
	for _, r := range m.r {
		if r.Match([]byte(s)) {
			return true
		}
	}
	return false
}

// Report returns the number of string and regex entries in the Matcher
func (m *Matcher) Report() (int, int) {
	return len(m.s), len(m.r)
}
