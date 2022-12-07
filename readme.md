# csql - Query language for comma-separated values

csql is a small tool allowing to run SQL queries on CSV files.

csql supports basic<sup>*</sup> `SELECT` statement features, including:

- Joins
  - Left, Right, Full, Cross and Inner<sup>**</sup>
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

<sup>**</sup>Only JOIN ... ON ... is supported. JOIN USING is not supported.

## Usage

TBD

## Building

csql is a simple Go project, so it can be built as any other Go application

    $ go test ./...
    $ go build

Or use `build.sh` script to build for all supported platforms. Results will be placed in the `/bin` folder. 

### Changing grammar

csql uses ANTLR4 parser generator. ANTLR is a Java application and so you will need Java if you changed the grammar.
Run `antlr/generate.sh` from the project root to build parser from the `Csql.g4` grammar file.
