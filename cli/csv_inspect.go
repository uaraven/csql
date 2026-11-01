package cli

import (
	"fmt"
	"github.com/uaraven/csql/core"
	"io"
	"os"
)

type CsvColumn struct {
	Name string
	Type string
}

const unknownType = "Unknown"

var typeNames = map[core.DataType]string{
	core.TypeInt:     "Number(Int)",
	core.TypeFloat:   "Number(Float)",
	core.TypeString:  "String",
	core.TypeBoolean: "Bool",
	core.TypeNull:    "Null",
}

// inspectRows is the number of data rows that are used to find out the type of a column
const inspectRows = 100

// InspectCsv reads the header and the first rows of the CSV file. The type of a column is the type of its first
// value that is not NULL. The type is "Unknown" if there is no such value.
func InspectCsv(csvFile string) ([]CsvColumn, error) {
	csvFile = core.ExpandCsvFileName(csvFile)
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, fmt.Errorf("file doesn't exist: %s\n%v", csvFile, err)
	}
	defer func() { _ = file.Close() }()
	csvReader := core.NewCsvReader(file)
	headers, err := csvReader.Read()
	if err == io.EOF {
		return nil, fmt.Errorf("file is empty, expected a header row: %s", csvFile)
	} else if err != nil {
		return nil, fmt.Errorf("error reading from file: %s\n%v", csvFile, err)
	}
	result := make([]CsvColumn, len(headers))
	for i, column := range headers {
		result[i] = CsvColumn{Name: column, Type: unknownType}
	}
	for row := 0; row < inspectRows; row++ {
		csvRow, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading from file: %s\n%v", csvFile, err)
		}
		for i := range result {
			if result[i].Type != unknownType {
				continue
			}
			if v := core.ParseStringToValue(csvRow[i]); v.Type() != core.TypeNull {
				result[i].Type = typeNames[v.Type()]
			}
		}
	}
	return result, nil
}
