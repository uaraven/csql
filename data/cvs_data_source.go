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
	"sync/atomic"
)

type csvDataSource struct {
	io.Closer
	lock          sync.Mutex
	path          string
	name          string
	index         atomic.Int64
	file          *os.File
	reader        *csv.Reader
	headers       DataSourceHeader
	currentValues []string
}

func NewCsvDataSource(csvFile string) (DataSource, error) {
	return NewCsvDataSourceWithAlias(csvFile, "")
}

func NewCsvDataSourceWithAlias(csvFile string, alias string) (DataSource, error) {
	_, nameExt := filepath.Split(csvFile)
	var justName string
	if alias != "" {
		justName = alias
	} else {
		justName = strings.TrimSuffix(nameExt, filepath.Ext(nameExt))
	}
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
		path:   csvFile,
		file:   file,
		name:   justName,
		reader: csvReader,
	}
	cds.index.Store(-1)
	cds.headers = NewHeadersFromSlice(cds, headers)
	return cds, nil
}

func (cds *csvDataSource) Header() DataSourceHeader {
	return cds.headers
}

func (cds *csvDataSource) GetName() string {
	return cds.name
}

func (cds *csvDataSource) MatchesName(s string) bool {
	return util.EqualsIgnoreCase(s, cds.name)
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
	cds.currentValues = values
	return newRowWithId(cds.index.Add(1), cds, values), nil
}

func (cds *csvDataSource) CurrentRow() (Row, error) {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	if cds.currentValues == nil {
		return nil, nil
	} else {
		return newRowWithId(cds.index.Load(), cds, cds.currentValues), nil
	}
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
	cds.index.Store(-1)
	cds.currentValues = nil
	_, err = cds.reader.Read()
	return err
}

func (cds *csvDataSource) Close() error {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	return cds.file.Close()
}
