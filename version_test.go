package go_util

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}

	// VERSION 파일에 기록된 값(0.1.0)과 일치하는지 확인
	expected := "0.1.2"
	if Version != expected {
		t.Errorf("Expected version %s, but got %s", expected, Version)
	}
}
