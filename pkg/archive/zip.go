package archive

import (
	"archive/zip"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Zip un archives zip archive file.
type Zip struct{}

// UnArchive unpacks the .zip file at source to destination.
func (z Zip) UnArchive(source, destination string) error {
	r, err := zip.OpenReader(source)
	if err != nil {
		return fmt.Errorf("open reader error: %w", err)
	}
	defer r.Close()
	log.Infof("unzip asset: [source: %s, destination: %s]", source, destination)
	if err := mkdir(destination, os.ModePerm); err != nil {
		return err
	}
	for _, zf := range r.File {
		f, err := zf.Open()
		if err != nil {
			return fmt.Errorf("open file `%s` error: %w", zf.Name, err)
		}
		fileInfo := zf.FileInfo()
		filePath := filepath.Join(destination, zf.Name)
		log.Infof("unzip files: [destination: %s, file: %s]", source, zf.Name)
		if fileInfo.IsDir() {
			if err := mkdir(filePath, fileInfo.Mode()); err != nil {
				return err
			}
			continue
		}

		if err := writeNewFile(filePath, f, fileInfo.Mode()); err != nil {
			return err
		}
	}
	return nil
}
