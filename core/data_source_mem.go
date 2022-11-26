package core

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type memDataSource struct {
	lock    sync.Mutex
	name    string
	index   int
	data    []Row
	headers DataSourceHeader
}

func NewMemDataSourceFromCsv(csvFile string) (DataSource, error) {
	return NewMemDataSourceWithAlias(csvFile, "")
}

func NewMemDataSource(name string, headers []ColumnMetadata, rows []Row) (DataSource, error) {
	ds := &memDataSource{
		lock:  sync.Mutex{},
		name:  name,
		index: -1,
		data:  rows,
	}

	ds.headers = dataSourceHeader{
		parent:  ds,
		columns: headers,
	}

	return ds, nil
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
		name:  justName,
		index: -1,
	}
	cds.headers = NewHeadersFromSlice(cds, headers)
	err = cds.loadCsv(csvReader)
	if err != nil {
		return nil, err
	}
	return cds, nil
}

func (cds *memDataSource) loadCsv(csvReader *csv.Reader) error {
	rows := make([]Row, 0)
	index := 0
	for {
		csvRow, err := csvReader.Read()
		if err == io.EOF {
			cds.data = rows
			return nil
		} else if err != nil {
			return err
		}
		r := parseRowWithId(index, cds.Header(), csvRow)
		index++
		rows = append(rows, r)
	}
}

func (cds *memDataSource) Header() DataSourceHeader {
	return cds.headers
}

func (cds *memDataSource) GetName() string {
	return cds.name
}

func (cds *memDataSource) NextRow() (Row, error) {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	if cds.index+1 >= len(cds.data) {
		return nil, nil
	}
	cds.index++
	return cds.data[cds.index], nil
}

func (cds *memDataSource) CurrentRow() (Row, error) {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	if cds.index < 0 {
		return nil, nil
	}
	return cds.data[cds.index], nil
}

func (cds *memDataSource) Rewind() error {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	cds.index = -1
	return nil
}

func (cds *memDataSource) GetRows() []Row {
	dest := make([]Row, len(cds.data))
	copy(dest, cds.data)
	return dest
}
