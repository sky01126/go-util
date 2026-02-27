package strings

import (
	"strings"
	"unicode"
)

type CasePolicy struct {
	caseSensitive bool
}

var (
	// CS 대소문자를 구분한다 (Case Sensitive).
	CS = &CasePolicy{caseSensitive: true}
	// CI 대소문자를 구분하지 않는다 (Case Insensitive).
	CI = &CasePolicy{caseSensitive: false}
)

// Equals 문자열이 동일한지 비교한다.
func (cp *CasePolicy) Equals(str1, str2 string) bool {
	if cp.caseSensitive {
		return str1 == str2
	}
	return strings.EqualFold(str1, str2)
}

// EqualsAny 문자열이 여러 대상 중 하나와 동일한지 확인한다.
func (cp *CasePolicy) EqualsAny(str string, searchStrings ...string) bool {
	for _, s := range searchStrings {
		if cp.Equals(str, s) {
			return true
		}
	}
	return false
}

// IsEmpty 문자열이 비어 있는지("") 확인한다.
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty 문자열이 비어 있지 않은지 확인한다.
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsBlank 문자열이 공백, 빈 문자열("") 또는 널인지 확인한다.
func IsBlank(str string) bool {
	if len(str) == 0 {
		return true
	}
	for _, r := range str {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// IsNotBlank 문자열이 비어 있지 않고 공백만으로 구성되지 않았는지 확인한다.
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// Trim 문자열 양 끝에서 제어 문자를 제거한다.
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// TrimToEmpty 문자열을 트리밍하며 빈 문자열이면 빈 문자열을 반환한다.
func TrimToEmpty(str string) string {
	return strings.TrimSpace(str)
}

// TrimToNull 문자열을 트리밍하며 빈 문자열이면 nil을 반환한다.
func TrimToNull(str string) *string {
	ts := strings.TrimSpace(str)
	if ts == "" {
		return nil
	}
	return &ts
}

// Capitalize 문자열의 첫 글자를 대문자로 변경한다.
func Capitalize(str string) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Uncapitalize 문자열의 첫 글자를 소문자로 변경한다.
func Uncapitalize(str string) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// DefaultIfEmpty 문자열이 빈 문자열인 경우 기본값을 반환한다.
func DefaultIfEmpty(str, defaultStr string) string {
	if str == "" {
		return defaultStr
	}
	return str
}

// DefaultIfBlank 문자열이 공백 또는 빈 문자열인 경우 기본값을 반환한다.
func DefaultIfBlank(str, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	}
	return str
}

// Abbreviate 말줄임표를 사용하여 문자열을 축약한다.
func Abbreviate(str string, maxWidth int) string {
	if str == "" {
		return str
	}
	if maxWidth < 4 {
		return str
	}
	runes := []rune(str)
	if len(runes) <= maxWidth {
		return str
	}
	return string(runes[0:maxWidth-3]) + "..."
}
