package files

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Exists 파일 또는 디렉토리가 존재하는지 확인한다.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// IsFile 경로가 일반 파일인지 확인한다.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsDir 경로가 디렉토리인지 확인한다.
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// UserHomeDir 현재 사용자의 홈 디렉토리 경로를 반환한다.
func UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

// AddDirectories 디렉토리를 생성한다.
func AddDirectories(path string) error {
	return os.MkdirAll(path, 0755)
}

// Remove 파일 또는 디렉토리를 삭제한다.
func Remove(path string) error {
	return os.RemoveAll(path)
}

// ReadString 파일의 내용을 문자열로 읽는다.
func ReadString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read file %s: %w", path, err)
	}
	return string(data), nil
}

// WriteString 파일에 문자열을 작성한다.
func WriteString(path string, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("write file %s: %w", path, err)
	}
	return nil
}

// PathSeparator OS별 경로 구분자를 반환한다.
func PathSeparator() string {
	return string(os.PathSeparator)
}

// Join 여러 경로 요소를 하나로 합친다.
func Join(elements ...string) string {
	return filepath.Join(elements...)
}

// Copy 파일을 복사한다.
func Copy(src, dst string) (err error) {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open source file %s: %w", src, err)
	}
	defer func() {
		err := sourceFile.Close()
		if err != nil {
			fmt.Printf("Error closing source file: %v\n", err)
		}
	}()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create destination file %s: %w", dst, err)
	}
	defer func() {
		closeErr := destFile.Close()
		if err == nil && closeErr != nil {
			err = fmt.Errorf("close destination file %s: %w", dst, closeErr)
		}
	}()

	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("copy content from %s to %s: %w", src, dst, err)
	}
	return err
}
