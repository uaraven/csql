package data

import "csql/funky"

func loadTestDatasource() DataSource {
	ds, err := NewCsvDataSource("../test-data/employees.csv")
	if err != nil {
		panic(err)
	}
	return ds
}

func loadTestMemDatasource() DataSource {
	ds, err := NewMemDataSource("../test-data/employees.csv")
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
