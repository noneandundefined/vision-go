package pkg

import (
	"os"
	"path/filepath"
	"strings"
)

// The FindLogFiles function searches for log files
// in the log/ or logs/ directories, regardless of
// the location (it searches in the main directory).
// These files are then sent to the mail with statistics.
func FindLogFiles(directory string) ([]string, error) {
	var logFiles []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && (info.Name() == "log" || info.Name() == "logs") {
			err := filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !fileInfo.IsDir() && (strings.HasSuffix(fileInfo.Name(), ".log") || strings.HasSuffix(fileInfo.Name(), ".logs")) {
					logFiles = append(logFiles, filePath)
				}

				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	return logFiles, err
}
