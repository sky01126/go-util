package compress

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mholt/archiver/v4"
)

// Uncompress 압축 파일을 대상 디렉토리에 해제한다.
// 파일 이름을 통해 압축 형식을 자동으로 감지한다.
func Uncompress(source, destination string) error {
	if source == "" {
		return fmt.Errorf("source path is empty")
	}

	// 파일 열기
	f, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(f)

	ctx := context.Background()

	// 형식 식별
	format, _, err := archiver.Identify(ctx, source, f)
	if err != nil {
		return fmt.Errorf("failed to identify format for %s: %w", source, err)
	}

	ex, ok := format.(archiver.Extractor)
	if !ok {
		return fmt.Errorf("format %s does not support extraction", source)
	}

	// 대상 디렉토리가 없으면 생성
	if err := os.MkdirAll(destination, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// 압축 해제 핸들러
	handler := func(_ context.Context, archFile archiver.FileInfo) error {
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
		defer func(out *os.File) {
			err := out.Close()
			if err != nil {
				fmt.Printf("Error closing file: %v\n", err)
			}
		}(out)

		in, err := archFile.Open()
		if err != nil {
			return err
		}
		defer func(in fs.File) {
			err := in.Close()
			if err != nil {
				fmt.Printf("Error closing file: %v\n", err)
			}
		}(in)

		_, err = io.Copy(out, in)
		return err
	}

	if err := ex.Extract(ctx, f, handler); err != nil {
		return fmt.Errorf("failed to extract %s: %w", source, err)
	}

	return nil
}

// Compress 파일 또는 디렉토리들을 압축한다.
// 대상 파일의 확장자를 통해 압축 형식을 자동으로 결정한다.
func Compress(sources []string, destination string) error {
	if len(sources) == 0 {
		return fmt.Errorf("no sources provided for compression")
	}
	if destination == "" {
		return fmt.Errorf("destination path is empty")
	}

	// v4에서 확장자로 포맷 결정
	var format archiver.Format
	ext := filepath.Ext(destination)
	switch ext {
	case ".zip":
		format = archiver.Zip{}
	case ".tar":
		format = archiver.Tar{}
	case ".gz", ".tgz":
		format = archiver.Archive{
			Compression: archiver.Gz{},
			Archival:    archiver.Tar{},
		}
	case ".bz2":
		format = archiver.Archive{
			Compression: archiver.Bz2{},
			Archival:    archiver.Tar{},
		}
	case ".xz":
		format = archiver.Archive{
			Compression: archiver.Xz{},
			Archival:    archiver.Tar{},
		}
	case ".lz4":
		format = archiver.Archive{
			Compression: archiver.Lz4{},
			Archival:    archiver.Tar{},
		}
	default:
		return fmt.Errorf("unsupported archive format: %s", ext)
	}

	ar, ok := format.(archiver.Archiver)
	if !ok {
		return fmt.Errorf("format %s does not support archiving", destination)
	}

	// 파일 맵 생성
	var files []archiver.FileInfo
	for _, src := range sources {
		absSrc, err := filepath.Abs(src)
		if err != nil {
			return fmt.Errorf("failed to get absolute path for %s: %w", src, err)
		}

		err = filepath.Walk(absSrc, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relPath, err := filepath.Rel(filepath.Dir(absSrc), path)
			if err != nil {
				return err
			}

			currPath := path
			files = append(files, archiver.FileInfo{
				FileInfo:      info,
				NameInArchive: relPath,
				Open: func() (fs.File, error) {
					return os.Open(currPath)
				},
			})
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to walk source %s: %w", src, err)
		}
	}

	out, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(out)

	if err := ar.Archive(context.Background(), out, files); err != nil {
		return fmt.Errorf("failed to create archive %s: %w", destination, err)
	}

	return nil
}

// IsSupported 지원되는 압축 형식인지 확인한다.
func IsSupported(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".zip", ".tar", ".gz", ".tgz", ".bz2", ".xz", ".lz4", ".sz", ".br":
		return true
	}
	return false
}

// GetExtension 지원되는 압축 파일의 확장자를 반환한다.
func GetExtension(filename string) string {
	return filepath.Ext(filename)
}
