# csql - Query language for comma-separated values

csql is a small tool allowing to run SQL queries on CSV files.

csql supports basic<sup>*</sup> `SELECT` statement features, including:

- Joins
  - Left, Right, Full, Cross and Inner
- UNION and UNION ALL
- LIKE conditions
- MATCH conditions (like LIKE, but using regular expressions)
- BETWEEN conditions
- IN (...list...) conditions
- Logical operators `AND`, `OR` and `NOT`
- LIMIT and ORDER BY
- Table and column aliases (`SELECT t.column as col FROM "Table.csv" T`)

csql automatically converts values to integer and float types and nulls. 

---
<sup>*</sup> IN (select) and aggregations are not supported yet in version 0.1.x