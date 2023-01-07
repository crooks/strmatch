package strmatch

import (
	"testing"
)

func TestMatcher(t *testing.T) {
	var err error
	m := NewMatcher()
	err = m.SAdd("baz")
	if err != nil {
		t.Error(err)
	}
	if m.Match("bazz") {
		t.Error("\"bazz\" should not match")
	}
	err = m.SAdd("bazz")
	if err != nil {
		t.Error(err)
	}
	if !m.Match("bazz") {
		t.Error("\"bazz\" should match")
	}
	err = m.RAdd("^foo")
	if err != nil {
		t.Error(err)
	}
	if m.Match("barfoo") {
		t.Error("string \"barfoo\" shouldn't match the test Regex")
	}
	err = m.RAdd("^bar")
	if err != nil {
		t.Error(err)
	}
	if !m.Match("barfoo") {
		t.Error("string \"barfoo\" should match the test Regex")
	}
}
