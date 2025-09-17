// Package dfmt formats values for display purposes.
package dfmt

import (
	"encoding/json"
	"strconv"
	"strings"
	"unicode"
)

// FormatValue formats a value for display as part of a key=value pair.
// It quotes as minimally as possible, to keep things readable.
func FormatValue(v any) string {
	if s, gotstr := v.(string); gotstr {
		if !shouldQuoteString(s) {
			return s
		}
		qs := strconv.Quote(s)
		if strings.ContainsRune(qs, '\\') && strconv.CanBackquote(s) {
			return "`" + s + "`"
		}
		return qs
	}

	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func shouldQuoteString(s string) bool {
	if strings.ContainsAny(s, "'\"`\\ \t\n\r") {
		return true
	}
	for _, r := range s {
		if strings.ContainsRune("-_.,<>()[]{}:;@#?/|+*&^%$~", r) {
			continue
		}
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			continue
		}
		return true
	}
	return false
}
