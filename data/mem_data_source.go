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

type memDataSource struct {
	io.Closer
	lock    sync.Mutex
	path    string
	name    string
	index   atomic.Int64
	data    []Row
	headers DataSourceHeader
}

func NewMemDataSource(csvFile string) (DataSource, error) {
	return NewCsvDataSourceWithAlias(csvFile, "")
}

func NewMemDataSourceWithAlias(csvFile string, alias string) (DataSource, error) {
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
	defer func() { _ = file.Close() }()
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.TrimLeadingSpace = true
	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	cds := &memDataSource{
		path: csvFile,
		name: justName,
	}
	err = cds.loadCsv(csvReader)
	if err != nil {
		return nil, err
	}
	cds.index.Store(-1)
	cds.headers = NewHeadersFromSlice(cds, headers)
	return cds, nil
}

func (cds *memDataSource) loadCsv(csvReader *csv.Reader) error {
	rows := make([]Row, 0)
	index := int64(0)
	for {
		csvRow, err := csvReader.Read()
		r := newRowWithId(index, cds, csvRow)
		index++
		rows = append(rows, r)
		if err == io.EOF {
			cds.data = rows
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (cds *memDataSource) Header() DataSourceHeader {
	return cds.headers
}

func (cds *memDataSource) GetName() string {
	return cds.name
}

func (cds *memDataSource) MatchesName(s string) bool {
	return util.EqualsIgnoreCase(s, cds.name)
}

func (cds *memDataSource) NextRow() (Row, error) {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	readIndex := cds.index.Add(1)
	if readIndex >= int64(len(cds.data)) {
		return nil, nil
	}
	return cds.data[readIndex], nil
}

func (cds *memDataSource) Rewind() error {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	cds.index.Store(-1)
	return nil
}

func (cds *memDataSource) ReadAll() ([]Row, error) {
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

func (cds *memDataSource) Close() error {
	return nil
}
