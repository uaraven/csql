package data

func loadTestDatasource() DataSource {
	ds, err := NewCsvDataSource("../test-data/employees.csv")
	if err != nil {
		panic(err)
	}
	return ds
}
