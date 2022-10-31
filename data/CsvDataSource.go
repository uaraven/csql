package data

import (
	"bufio"
	"csql/util"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type csvDataSource struct {
	io.Closer
	name    string
	alias   string
	file    *os.File
	reader  *csv.Reader
	headers DataSourceHeader
}

func NewCsvDataSource(csvFile string) (DataSource, error) {
	_, nameExt := filepath.Split(csvFile)
	justName := strings.TrimSuffix(nameExt, filepath.Ext(nameExt))
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.TrimLeadingSpace = true
	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	cds := &csvDataSource{
		file:   file,
		name:   justName,
		reader: csvReader,
	}
	cds.headers = NewHeadersFromSlice(cds, headers)
	return cds, nil
}

func (cds *csvDataSource) Header() DataSourceHeader {
	return cds.headers
}

func (cds *csvDataSource) MatchesName(s string) bool {
	return util.EqualsIgnoreCase(s, cds.alias) || util.EqualsIgnoreCase(s, cds.name)
}

func (cds *csvDataSource) NextRow() (Row, error) {
	values, err := cds.reader.Read()
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &rowImpl{
		parent: cds,
		values: values,
	}, nil
}

func (cds *csvDataSource) Close() error {
	return cds.file.Close()
}
