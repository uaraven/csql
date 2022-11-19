package data

import (
	"bufio"
	"csql/util"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type csvDataSource struct {
	io.Closer
	lock    sync.Mutex
	path    string
	name    string
	alias   string
	index   int64
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
		index:  -1,
		path:   csvFile,
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
	cds.lock.Lock()
	defer cds.lock.Unlock()
	values, err := cds.reader.Read()
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	cds.index += 1
	return &rowImpl{
		id:     cds.index,
		parent: cds,
		values: values,
	}, nil
}

func (cds *csvDataSource) Rewind() error {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	err := cds.file.Close()
	if err != nil {
		return err
	}
	file, err := os.Open(cds.path)
	if err != nil {
		return err
	}
	cds.file = file
	cds.reader = csv.NewReader(bufio.NewReader(cds.file))
	cds.reader.TrimLeadingSpace = true
	cds.index = -1
	_, err = cds.reader.Read()
	return err
}

func (cds *csvDataSource) ReadAll() ([]Row, error) {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	result := make([]Row, 0)
	row, err := cds.NextRow()
	if err != nil {
		return nil, err
	}
	for row != nil {
		result = append(result, row)
		row, err = cds.NextRow()
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (cds *csvDataSource) Close() error {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	return cds.file.Close()
}
