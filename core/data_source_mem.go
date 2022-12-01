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

func NewMemDataSourceFromCsv(csvFile string) DataSource {
	return NewMemDataSourceWithAlias(csvFile, "")
}

func NewMemDataSource(name string, headers []ColumnMetadata, rows []Row) DataSource {
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

	return ds
}

func NewMemDataSourceWithAlias(csvFile string, alias string) DataSource {
	_, nameExt := filepath.Split(csvFile)
	var justName string
	if alias != "" {
		justName = alias
	} else {
		justName = strings.TrimSuffix(nameExt, filepath.Ext(nameExt))
	}
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.TrimLeadingSpace = true
	headers, err := csvReader.Read()
	if err != nil {
		panic(err)
	}
	cds := &memDataSource{
		name:  justName,
		index: -1,
	}
	cds.headers = NewHeadersFromSlice(cds, headers)
	err = cds.loadCsv(csvReader)
	if err != nil {
		panic(err)
	}
	return cds
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

func (cds *memDataSource) NextRow() Row {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	if cds.index+1 >= len(cds.data) {
		return nil
	}
	cds.index++
	return cds.data[cds.index]
}

func (cds *memDataSource) CurrentRow() Row {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	if cds.index < 0 {
		return nil
	}
	return cds.data[cds.index]
}

func (cds *memDataSource) Rewind() {
	cds.lock.Lock()
	defer cds.lock.Unlock()
	cds.index = -1
}

func (cds *memDataSource) GetRows() []Row {
	dest := make([]Row, len(cds.data))
	copy(dest, cds.data)
	return dest
}
