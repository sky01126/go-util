package compress

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mholt/archives"
)

// Uncompress 압축 파일을 대상 디렉토리에 해제한다.
// 파일 이름을 통해 압축 형식을 자동으로 감지한다.
func Uncompress(ctx context.Context, source, destination string) error {
	if source == "" {
		return fmt.Errorf("source path is empty")
	}

	// 파일 열기
	f, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func() {
		_ = f.Close()
	}()

	// 형식 식별 (반환된 stream은 식별 과정에서 읽은 바이트가 복원된 reader)
	format, stream, err := archives.Identify(ctx, source, f)
	if err != nil {
		return fmt.Errorf("failed to identify format for %s: %w", source, err)
	}

	ex, ok := format.(archives.Extractor)
	if !ok {
		return fmt.Errorf("format %s does not support extraction", source)
	}

	// 대상 디렉토리가 없으면 생성
	if err := os.MkdirAll(destination, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// 압축 해제 핸들러
	handler := func(ctx context.Context, archFile archives.FileInfo) error {
		return extractArchiveFile(destination, archFile)
	}

	if err := ex.Extract(ctx, stream, handler); err != nil {
		return fmt.Errorf("failed to extract %s: %w", source, err)
	}

	return nil
}

// extractArchiveFile 압축 항목 하나를 대상 경로에 기록한다.
func extractArchiveFile(destination string, archFile archives.FileInfo) error {
	targetPath := filepath.Join(destination, archFile.NameInArchive)

	if archFile.IsDir() {
		return os.MkdirAll(targetPath, archFile.Mode())
	}

	// 부모 디렉토리 생성
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	out, err := os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, archFile.Mode())
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	in, err := archFile.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = in.Close()
	}()

	_, err = io.Copy(out, in)
	return err
}

// Compress 파일 또는 디렉토리들을 압축한다.
// 대상 파일의 확장자를 통해 압축 형식을 자동으로 결정한다.
func Compress(ctx context.Context, sources []string, destination string) error {
	if len(sources) == 0 {
		return fmt.Errorf("no sources provided for compression")
	}
	if destination == "" {
		return fmt.Errorf("destination path is empty")
	}

	ar, err := resolveArchiveFormat(destination)
	if err != nil {
		return err
	}

	files, err := buildFileList(sources)
	if err != nil {
		return err
	}

	out, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func() {
		_ = out.Close()
	}()

	if err := ar.Archive(ctx, out, files); err != nil {
		return fmt.Errorf("failed to create archive %s: %w", destination, err)
	}

	return nil
}

// resolveArchiveFormat 대상 파일 확장자로 압축 형식을 결정한다.
func resolveArchiveFormat(destination string) (archives.Archiver, error) {
	var format archives.Format
	ext := filepath.Ext(destination)
	switch ext {
	case ".zip":
		format = archives.Zip{}
	case ".tar":
		format = archives.Tar{}
	case ".gz", ".tgz":
		format = archives.CompressedArchive{Compression: archives.Gz{}, Archival: archives.Tar{}}
	case ".bz2":
		format = archives.CompressedArchive{Compression: archives.Bz2{}, Archival: archives.Tar{}}
	case ".xz":
		format = archives.CompressedArchive{Compression: archives.Xz{}, Archival: archives.Tar{}}
	case ".lz4":
		format = archives.CompressedArchive{Compression: archives.Lz4{}, Archival: archives.Tar{}}
	default:
		return nil, fmt.Errorf("unsupported archive format: %s", ext)
	}

	ar, ok := format.(archives.Archiver)
	if !ok {
		return nil, fmt.Errorf("format %s does not support archiving", destination)
	}
	return ar, nil
}

// buildFileList 소스 경로들을 순회해 압축 대상 파일 목록을 만든다.
func buildFileList(sources []string) ([]archives.FileInfo, error) {
	var files []archives.FileInfo
	for _, src := range sources {
		absSrc, err := filepath.Abs(src)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path for %s: %w", src, err)
		}

		err = filepath.Walk(absSrc, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relPath, err := filepath.Rel(filepath.Dir(absSrc), path)
			if err != nil {
				return err
			}

			files = append(files, archives.FileInfo{
				FileInfo:      info,
				NameInArchive: relPath,
				Open: func() (fs.File, error) {
					return os.Open(path)
				},
			})
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to walk source %s: %w", src, err)
		}
	}
	return files, nil
}

// IsSupported 지원되는 압축 형식인지 확인한다.
func IsSupported(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".zip", ".tar", ".gz", ".tgz", ".bz2", ".xz", ".lz4":
		return true
	}
	return false
}

// GetExtension 지원되는 압축 파일의 확장자를 반환한다.
func GetExtension(filename string) string {
	return filepath.Ext(filename)
}
