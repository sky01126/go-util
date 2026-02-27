package strings

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.Errorf("IsEmpty(\"\") should be true")
	}
	if IsEmpty(" ") {
		t.Errorf("IsEmpty(\" \") should be false")
	}
	if IsEmpty("abc") {
		t.Errorf("IsEmpty(\"abc\") should be false")
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank("") {
		t.Errorf("IsBlank(\"\") should be true")
	}
	if !IsBlank(" ") {
		t.Errorf("IsBlank(\" \") should be true")
	}
	if !IsBlank("  \t\n") {
		t.Errorf("IsBlank whitespace should be true")
	}
	if IsBlank("abc") {
		t.Errorf("IsBlank(\"abc\") should be false")
	}
}

func TestTrimToNull(t *testing.T) {
	if TrimToNull("") != nil {
		t.Errorf("TrimToNull(\"\") should be nil")
	}
	if TrimToNull("  ") != nil {
		t.Errorf("TrimToNull(\"  \") should be nil")
	}
	res := TrimToNull(" abc ")
	if res == nil || *res != "abc" {
		t.Errorf("TrimToNull(\" abc \") should be \"abc\"")
	}
}

func TestCapitalize(t *testing.T) {
	if Capitalize("cat") != "Cat" {
		t.Errorf("Capitalize(\"cat\") failed")
	}
	if Capitalize("Cat") != "Cat" {
		t.Errorf("Capitalize(\"Cat\") failed")
	}
	if Capitalize("") != "" {
		t.Errorf("Capitalize(\"\") failed")
	}
}

func TestAbbreviate(t *testing.T) {
	if Abbreviate("abcdefg", 6) != "abc..." {
		t.Errorf("Abbreviate(\"abcdefg\", 6) failed: %s", Abbreviate("abcdefg", 6))
	}
	if Abbreviate("abcdefg", 7) != "abcdefg" {
		t.Errorf("Abbreviate(\"abcdefg\", 7) failed")
	}
	if Abbreviate("abcdefg", 3) != "abcdefg" {
		t.Errorf("Abbreviate with small maxWidth failed")
	}
}

func TestEquals(t *testing.T) {
	if !CS.Equals("abc", "abc") {
		t.Errorf("CS.Equals(\"abc\", \"abc\") should be true")
	}
	if CS.Equals("abc", "ABC") {
		t.Errorf("CS.Equals(\"abc\", \"ABC\") should be false")
	}
	if !CI.Equals("abc", "ABC") {
		t.Errorf("CI.Equals(\"abc\", \"ABC\") should be true")
	}
}

func TestEqualsAny(t *testing.T) {
	if !CS.EqualsAny("abc", "abc", "def") {
		t.Errorf("CS.EqualsAny(\"abc\", \"abc\", \"def\") should be true")
	}
	if CS.EqualsAny("abc", "ABC", "def") {
		t.Errorf("CS.EqualsAny(\"abc\", \"ABC\", \"def\") should be false")
	}
	if !CI.EqualsAny("abc", "ABC", "def") {
		t.Errorf("CI.EqualsAny(\"abc\", \"ABC\", \"def\") should be true")
	}
	if CS.EqualsAny("abc") {
		t.Errorf("CS.EqualsAny with no search strings should be false")
	}
}
