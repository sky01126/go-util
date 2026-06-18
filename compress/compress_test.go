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
	err = Compress(t.Context(), []string{testFile}, zipFile)
	if err != nil {
		t.Fatalf("Compress failed: %v", err)
	}

	// 압축 파일 존재 확인
	if !files.Exists(zipFile) {
		t.Fatal("Zip file was not created")
	}

	// 2. 해제 테스트
	extractDir := filepath.Join(tmpDir, "extracted")
	err = Uncompress(t.Context(), zipFile, extractDir)
	if err != nil {
		t.Fatalf("Uncompress failed: %v", err)
	}

	// 해제된 파일 내용 확인
	// archives zip은 기본적으로 파일명을 유지함
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

func TestIsSupportedEdgeCases(t *testing.T) {
	tests := []struct {
		filename string
		want     bool
	}{
		{"", false},
		{"noext", false},
		{"archive.xz", true},
		{"archive.lz4", true},
		{"archive.sz", false},
		{"archive.br", false},
		{"archive.rar", false},
	}

	for _, tt := range tests {
		if got := IsSupported(tt.filename); got != tt.want {
			t.Errorf("IsSupported(%q) = %v, want %v", tt.filename, got, tt.want)
		}
	}
}

func TestGetExtension(t *testing.T) {
	tests := []struct {
		filename string
		want     string
	}{
		{"test.zip", ".zip"},
		{"test.tar.gz", ".gz"},
		{"noext", ""},
		{"", ""},
	}

	for _, tt := range tests {
		if got := GetExtension(tt.filename); got != tt.want {
			t.Errorf("GetExtension(%q) = %q, want %q", tt.filename, got, tt.want)
		}
	}
}

func TestCompressErrors(t *testing.T) {
	zipDst := filepath.Join(t.TempDir(), "out.zip")
	txtDst := filepath.Join(t.TempDir(), "out.txt")

	tests := []struct {
		name        string
		sources     []string
		destination string
	}{
		{"소스가 비어 있으면 에러를 반환한다", nil, zipDst},
		{"대상 경로가 비어 있으면 에러를 반환한다", []string{"dummy"}, ""},
		{"지원하지 않는 확장자면 에러를 반환한다", []string{"dummy"}, txtDst},
	}

	for _, tt := range tests {
		if err := Compress(t.Context(), tt.sources, tt.destination); err == nil {
			t.Errorf("%s: Compress(%v, %q) = nil, want error", tt.name, tt.sources, tt.destination)
		}
	}
}

func TestUncompressErrors(t *testing.T) {
	tmpDir := t.TempDir()

	// 압축 형식이 아닌 일반 텍스트 파일을 준비한다.
	plainFile := filepath.Join(tmpDir, "plain.txt")
	if err := files.WriteString(plainFile, "not an archive"); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		source      string
		destination string
	}{
		{"소스 경로가 비어 있으면 에러를 반환한다", "", filepath.Join(tmpDir, "out1")},
		{"존재하지 않는 파일이면 에러를 반환한다", filepath.Join(tmpDir, "missing.zip"), filepath.Join(tmpDir, "out2")},
		{"압축 형식이 아니면 에러를 반환한다", plainFile, filepath.Join(tmpDir, "out3")},
	}

	for _, tt := range tests {
		if err := Uncompress(t.Context(), tt.source, tt.destination); err == nil {
			t.Errorf("%s: Uncompress(%q) = nil, want error", tt.name, tt.source)
		}
	}
}
