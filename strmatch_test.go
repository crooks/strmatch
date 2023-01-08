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

func TestEmpty(t *testing.T) {
	m := NewMatcher()
	if m.Match("EmptyTest") {
		t.Error("Empty matcher should always return False")
	}
}

func TestReport(t *testing.T) {
	m := NewMatcher()
	sn1, rn1 := m.Report()
	if sn1 != 0 {
		t.Errorf("Empty Matcher string should report zero.  Got=%d", sn1)
	}
	if rn1 != 0 {
		t.Errorf("Empty Matcher Regular Expression should report zero.  Got=%d", rn1)
	}
	err := m.SAdd("bazz")
	if err != nil {
		t.Error(err)
	}
	sn2, rn2 := m.Report()
	if sn2 != 1 {
		t.Errorf("String matcher should report one entry.  Got=%d", sn2)
	}
	if rn2 != 0 {
		t.Errorf("Empty Matcher Regular Expression should report zero.  Got=%d", rn2)
	}
}
