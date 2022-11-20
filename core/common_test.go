package core

import (
	"csql/funky"
	"fmt"
)

func loadTestDatasource() DataSource {
	ds, err := NewCsvDataSource("../test-data/employees.csv")
	if err != nil {
		panic(err)
	}
	return ds
}

func loadDefaultTestMemDatasource() DataSource {
	ds, err := NewMemDataSource("../test-data/employees.csv")
	if err != nil {
		panic(err)
	}
	return ds
}

func loadTestMemDatasource(name string) DataSource {
	ds, err := NewMemDataSource(fmt.Sprintf("../test-data/%s.csv", name))
	if err != nil {
		panic(err)
	}
	return ds
}

func loadTestMemDatasourceWithAlias(alias string) DataSource {
	ds, err := NewMemDataSourceWithAlias("../test-data/employees.csv", alias)
	if err != nil {
		panic(err)
	}
	return ds
}

func ReadAllRows(ds DataSource) ([]Row, error) {
	result := make([]Row, 0)
	row, err := ds.NextRow()
	if err != nil {
		return nil, err
	}
	for row != nil {
		result = append(result, row)
		row, err = ds.NextRow()
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type mapContext struct {
	EvaluationContext
	id     int64
	values map[string]string
}

func newMapContext() *mapContext {
	return &mapContext{
		id:     0,
		values: make(map[string]string),
	}
}

func newMapContextWithId(id int64) *mapContext {
	return &mapContext{
		id:     id,
		values: make(map[string]string),
	}
}

func (ec *mapContext) Get(name string) funky.Option[Value] {
	if val, present := ec.values[name]; present {
		return funky.SomeOf(decodeToValue(val))
	}
	return funky.NoneOf[Value]()
}

func (ec *mapContext) Id() int64 {
	return ec.id
}

func (ec *mapContext) SetId(newId int64) {
	ec.id = newId
}

func (ec *mapContext) SetValue(name string, value string) {
	ec.values[name] = value
}

func (ec *mapContext) Clear() {
	ec.values = make(map[string]string)
}
