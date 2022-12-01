package core

import (
	"fmt"
	"github.com/uaraven/csql/funky"
)

func loadTestDatasource() DataSource {
	return NewMemDataSourceFromCsv("../test-data/employees.csv")
}

func loadDefaultTestMemDatasource() DataSource {
	return NewMemDataSourceFromCsv("../test-data/employees.csv")
}

func loadTestMemDatasource(name string) DataSource {
	return NewMemDataSourceFromCsv(fmt.Sprintf("../test-data/%s.csv", name))
}

func loadDsWithAlias(name string, alias string) DataSource {
	return NewMemDataSourceWithAlias(fmt.Sprintf("../test-data/%s.csv", name), alias)
}

func loadTestMemDatasourceWithAlias(alias string) DataSource {
	return NewMemDataSourceWithAlias("../test-data/employees.csv", alias)
}

type mapContext struct {
	EvaluationContext
	id     int
	values map[string]string
}

func newMapContext() *mapContext {
	return &mapContext{
		id:     0,
		values: make(map[string]string),
	}
}

func (ec *mapContext) Get(name string) funky.Option[Value] {
	if val, present := ec.values[name]; present {
		return funky.SomeOf(decodeToValue(val))
	}
	return funky.NoneOf[Value]()
}

func (ec *mapContext) Id() int {
	return ec.id
}

func (ec *mapContext) SetId(newId int) {
	ec.id = newId
}

func (ec *mapContext) SetValue(name string, value string) {
	ec.values[name] = value
}

func (ec *mapContext) Clear() {
	ec.values = make(map[string]string)
}
