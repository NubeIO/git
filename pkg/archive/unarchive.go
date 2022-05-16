package archive

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	// ErrNotSupportFile does not support file extensions.
	ErrNotSupportFile = errors.New("not support file extension")
)

// UnArchiver is a type that can extract archive files into a folder.
type UnArchiver interface {
	UnArchive(source, destination string) error
}

// UnArchive UnArchives the given archive file into the destination folder.
// The archive format is selected implicitly.
func UnArchive(source, destination string) error {
	unArchiver, err := byExtension(source)
	if err != nil {
		return fmt.Errorf("UnArchive `%s` error: %w", source, err)
	}

	if err := unArchiver.UnArchive(source, destination); err != nil {
		return fmt.Errorf("UnArchive `%s` error: %w", source, err)
	}
	return nil
}

// Support check for handle the archive file format.
func Support(filePath string) bool {
	_, err := byExtension(filePath)
	return err == nil
}

func byExtension(filePath string) (UnArchiver, error) {
	switch {
	case strings.HasSuffix(filePath, ".zip"):
		return Zip{}, nil

	case strings.HasSuffix(filePath, ".tar.gz"),
		strings.HasSuffix(filePath, ".tgz"):
		return TarGz{}, nil

	default:
		return nil, ErrNotSupportFile
	}
}

func mkdir(destPath string, mode os.FileMode) error {
	err := os.MkdirAll(destPath, mode)
	if err != nil {
		return fmt.Errorf("mkdir `%s` error: %w", destPath, err)
	}
	return nil
}

func writeNewFile(filePath string, in io.Reader, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("mkdir `%s` for file error: %w", filePath, err)
	}

	out, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("create file `%s` error: %w", filePath, err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("write file `%s`: %w", filePath, err)
	}
	return nil
}
