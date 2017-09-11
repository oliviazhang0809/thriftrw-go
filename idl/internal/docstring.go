package internal

import (
	"strings"
	"unicode"
)

const (
	_docstringOpen  = "/**"
	_docstringClose = "*/"
)

// ParseDocstring takes a docstring in the form,
//
//  /**
//   * foo bar
//   */
//
// And returns,
//
//  foo bar
func ParseDocstring(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return s
	}

	unindent(lines, true /* skipFirstIfUnindented */)

	lastIdx := len(lines) - 1

	// Strip comment markers from start and end.
	lines[0] = strings.TrimPrefix(lines[0], _docstringOpen)
	lines[lastIdx] = strings.TrimSuffix(lines[lastIdx], _docstringClose)

	if len(lines) == 1 {
		// Single-line doc block like, /** foo */
		return strings.TrimSpace(lines[0])
	}

	lines = dropLeadingEmptyLines(lines)
	lines = dropTrailingEmptyLines(lines)

	// At this point, we need to strip the leading "*" and " *" from every
	// line and unindent again.
	for i, l := range lines {
		switch {
		case len(l) > 0 && l[0] == '*':
			lines[i] = l[1:]
		case len(l) > 1 && l[:2] == " *":
			lines[i] = l[2:]
		}
	}

	unindent(lines, false /* skipFirstIfUnindented */)
	lines = dropLeadingEmptyLines(lines)
	lines = dropTrailingEmptyLines(lines)

	return strings.Join(lines, "\n")
}

func dropLeadingEmptyLines(lines []string) []string {
	for len(lines) > 0 {
		if len(lines[0]) > 0 {
			break
		}
		lines = lines[1:]
	}
	return lines
}

func dropTrailingEmptyLines(lines []string) []string {
	for i := len(lines) - 1; i >= 0; i-- {
		if len(lines[i]) > 0 {
			break
		}
		lines = lines[:i]
	}
	return lines
}

// Takes a series of lines that have been indented and removes the indentation
// of the first line from all lines.
//
// So,
//
//  "    foo"
//  "      bar"
//
// Becomes,
//
//  "foo"
//  "  bar"
//
// If the first line was not indentend, skipFirstIfUnindented controls whether
// we strip the indentation of the second line from all lines instead. This
// flag is needed because our docstringcs will start at "/**" without any
// leading whitespace.
func unindent(lines []string, skipFirstIfUnindented bool) {
	if len(lines) == 0 {
		return
	}

	var (
		havePrefix bool
		prefix     string
	)

	// Don't consider the first line for the prefix if it hasn't been
	// indented.
	if nonSpace := strings.IndexFunc(lines[0], isNotSpace); nonSpace >= 0 {
		if nonSpace == 0 && skipFirstIfUnindented {
			// The first line starts with a non-space character. Skip this line.
			lines = lines[1:]
		} else {
			havePrefix = true
			prefix = lines[0][:nonSpace]
		}
	}

	for i, s := range lines {
		nonSpace := strings.IndexFunc(s, isNotSpace)
		if nonSpace < 0 {
			// Whitespace-only. Use an empty string.
			lines[i] = ""
			continue
		}

		if !havePrefix {
			prefix = s[:nonSpace]
			havePrefix = true
		}

		// unindent only if the first non-space character appears at or after
		// the prefix.
		if nonSpace >= len(prefix) {
			lines[i] = s[len(prefix):]
		}
	}
}

func isNotSpace(r rune) bool {
	return !unicode.IsSpace(r)
}