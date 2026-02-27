package compress

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sky01126/go-util/files"
)

func TestCompressAndUncompress(t *testing.T) {
	// 임시 디렉토리 생성
	tmpDir := t.TempDir()

	sourceDir := filepath.Join(tmpDir, "source")
	err := os.Mkdir(sourceDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	// 테스트 파일 생성
	testFile := filepath.Join(sourceDir, "test.txt")
	testContent := "Hello, Archive!"
	err = files.WriteString(testFile, testContent)
	if err != nil {
		t.Fatal(err)
	}

	// 압축 파일 경로
	zipFile := filepath.Join(tmpDir, "test.zip")

	// 1. 압축 테스트
	err = Compress([]string{testFile}, zipFile)
	if err != nil {
		t.Fatalf("Compress failed: %v", err)
	}

	// 압축 파일 존재 확인
	if !files.Exists(zipFile) {
		t.Fatal("Zip file was not created")
	}

	// 2. 해제 테스트
	extractDir := filepath.Join(tmpDir, "extracted")
	err = Uncompress(zipFile, extractDir)
	if err != nil {
		t.Fatalf("Uncompress failed: %v", err)
	}

	// 해제된 파일 내용 확인
	// archiver/v3 zip은 기본적으로 파일명을 유지함
	extractedFile := filepath.Join(extractDir, "test.txt")
	if !files.Exists(extractedFile) {
		t.Fatalf("Extracted file not found at %s", extractedFile)
	}

	content, err := files.ReadString(extractedFile)
	if err != nil {
		t.Fatal(err)
	}

	if content != testContent {
		t.Errorf("Content mismatch. Expected: %s, Got: %s", testContent, content)
	}
}

func TestIsSupported(t *testing.T) {
	tests := []struct {
		filename string
		want     bool
	}{
		{"test.zip", true},
		{"test.tar.gz", true},
		{"test.txt", false},
		{"test.tar.bz2", true},
	}

	for _, tt := range tests {
		if got := IsSupported(tt.filename); got != tt.want {
			t.Errorf("IsSupported(%s) = %v, want %v", tt.filename, got, tt.want)
		}
	}
}
