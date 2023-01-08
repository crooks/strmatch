# Overview
strmatch is a [Go](https://go.dev) package that aims to provide a simple means to match a string against a list of strings and/or [Regular Expressions](https://en.wikipedia.org/wiki/Regular_expression).

# Example
```
import (
    "github.com/crooks/strmatch"
)

// Create a new matcher instance
m := NewMatcher()

// Add a string to match against
err := m.SAdd("foo")

// Test if "bar" generates a hit (will return False)
match := m.Match("bar")

// Add a Regular Expression
err := m.RAdd("^ba[a-z]")

// Test if "bar" generates a hit (will return True)
match = m.Match("bar")
```