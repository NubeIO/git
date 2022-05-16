package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// TarGz un archives tar.gz(tgz) archive file.
type TarGz struct{}

// UnArchive unpacks the .zip file at source to destination.
func (t TarGz) UnArchive(source, destination string) error {
	sf, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("open source file error: %w", err)
	}
	defer sf.Close()

	gr, err := gzip.NewReader(sf)
	if err != nil {
		return fmt.Errorf("open gzip reader error: %w", err)
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	if err := mkdir(destination, os.ModePerm); err != nil {
		return err
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("reader next error: %w", err)
		}

		filePath := filepath.Join(destination, header.Name)
		fileMode := os.FileMode(header.Mode)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := mkdir(filePath, fileMode); err != nil {
				return err
			}

		case tar.TypeReg, tar.TypeRegA, tar.TypeChar, tar.TypeBlock, tar.TypeFifo, tar.TypeGNUSparse:
			if err := writeNewFile(filePath, tr, fileMode); err != nil {
				return err
			}

		case tar.TypeXGlobalHeader, tar.TypeSymlink, tar.TypeLink: // ignore

		default:
			return fmt.Errorf("unknown type error: %v", header.Typeflag)
		}
	}
	return nil
}
