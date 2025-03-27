package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// this is an array (slice) of regular expressions
// (regexp.Regexp), designed to search for folder
// names containing a date in various formats.
var datePatterns = []*regexp.Regexp{
	regexp.MustCompile(`^(\d{2})\.(\d{2})$`),
	regexp.MustCompile(`^(\d{2})-(\d{2})$`),
	regexp.MustCompile(`^(\d{2})\.(\d{2})\.(\d{4})$`),
	regexp.MustCompile(`^(\d{2})\.(\d{2})\.(\d{2})$`),
	regexp.MustCompile(`^(\d{2})-(\d{2})-(\d{4})$`),
	regexp.MustCompile(`^(\d{2})-(\d{2})-(\d{2})$`),
}

// The FindLogFiles function searches for log files
// in the log/ or logs/ directories, regardless of
// the location (it searches in the main directory).
// These files are then sent to the mail with statistics.
func FindLogFiles(directory string, LOGFILES_BY_TIME_STYLES bool) ([]string, error) {
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
					if LOGFILES_BY_TIME_STYLES {
						today := time.Now()
						day := fmt.Sprintf("%02d", today.Day())
						month := fmt.Sprintf("%02d", int(today.Month()))
						year := fmt.Sprintf("%d", today.Year())
						shortYear := year[2:]

						parentDir := filepath.Base(filepath.Dir(filePath))

						for _, pattern := range datePatterns {
							if match := pattern.FindStringSubmatch(parentDir); match != nil {
								if match[1] == day && match[2] == month {
									if len(match) == 3 || match[3] == year || match[3] == shortYear {
										logFiles = append(logFiles, filePath)
									}
								}
							}
						}
					} else {
						logFiles = append(logFiles, filePath)
					}
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
