package core

import (
	"errors"
	"os"
	"path/filepath"
)

const csvExt = ".csv"

// ExpandCsvFileName checks if the provided file name doesn't have an extension.
// If there's no extension, it tries to add ".csv" extension. If the file with ".csv" extension exists it returns
// update file name, otherwise it returns original file name.
//
// This function will not fail if the original or original +".csv" files do not exist
func ExpandCsvFileName(csvFile string) string {
	if filepath.Ext(csvFile) == "" {
		_, err := os.Stat(csvFile)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			_, err = os.Stat(csvFile + csvExt)
			if err == nil {
				csvFile = csvFile + csvExt
			}
		}
	}
	return csvFile
}
