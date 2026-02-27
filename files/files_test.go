package files

import (
	"os"
	"testing"
)

func TestFiles(t *testing.T) {
	testFile := "test.txt"
	testContent := "hello world"

	// 1. WriteString 테스트
	err := WriteString(testFile, testContent)
	if err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(testFile)

	// 2. Exists, IsFile, IsDir 테스트
	if !Exists(testFile) {
		t.Errorf("Exists failed: file should exist")
	}
	if !IsFile(testFile) {
		t.Errorf("IsFile failed: should be a file")
	}
	if IsDir(testFile) {
		t.Errorf("IsDir failed: should not be a directory")
	}

	// 3. ReadString 테스트
	content, err := ReadString(testFile)
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if content != testContent {
		t.Errorf("GetString content mismatch: expected %s, got %s", testContent, content)
	}

	// 4. Copy 테스트
	copyFile := "test_copy.txt"
	err = Copy(testFile, copyFile)
	if err != nil {
		t.Fatalf("Copy failed: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(copyFile)

	if !Exists(copyFile) {
		t.Errorf("Copy failed: copy file should exist")
	}

	// 5. Remove 테스트
	err = Remove(copyFile)
	if err != nil {
		t.Fatalf("Remove failed: %v", err)
	}
	if Exists(copyFile) {
		t.Errorf("Remove failed: file should be deleted")
	}
}

func TestDirs(t *testing.T) {
	testDir := "test_dir/sub_dir"

	// 1. AddDirectories, IsDir, IsFile 테스트
	err := AddDirectories(testDir)
	if err != nil {
		t.Fatalf("AddDirectories failed: %v", err)
	}
	defer func() {
		_ = os.RemoveAll("test_dir")
	}()

	if !Exists(testDir) {
		t.Errorf("AddDirectories failed: directory should exist")
	}
	if !IsDir(testDir) {
		t.Errorf("IsDir failed: should be a directory")
	}
	if IsFile(testDir) {
		t.Errorf("IsFile failed: should not be a file")
	}

	// 2. Remove 테스트
	err = Remove("test_dir")
	if err != nil {
		t.Fatalf("Remove failed: %v", err)
	}
	if Exists("test_dir") {
		t.Errorf("Remove failed: directory should be deleted")
	}
}
