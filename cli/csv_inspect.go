package cli

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/uaraven/csql/core"
	"os"
)

type CsvColumn struct {
	Name string
	Type string
}

var typeNames = map[core.DataType]string{
	core.TypeInt:     "Number(Int)",
	core.TypeFloat:   "Number(Float)",
	core.TypeString:  "String",
	core.TypeBoolean: "Bool",
	core.TypeNull:    "Null",
}

func InspectCsv(csvFile string) ([]CsvColumn, error) {
	csvFile = core.ExpandCsvFileName(csvFile)
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, fmt.Errorf("file doesn't exist: %s\n%v", csvFile, err)
	}
	defer func() { _ = file.Close() }()
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.TrimLeadingSpace = true
	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading from file: %s\n%v", csvFile, err)
	}
	csvRow, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading from file: %s\n%v", csvFile, err)
	}
	if len(csvRow) != len(headers) {
		return nil, fmt.Errorf("number of columns doesn't match the number of headers in %s", csvFile)
	}
	result := make([]CsvColumn, len(headers))
	for i, column := range headers {
		result[i].Name = column
		v := core.ParseStringToValue(csvRow[i])
		result[i].Type = typeNames[v.Type()]
	}
	return result, nil
}
