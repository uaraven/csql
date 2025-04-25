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
- Aggregate functions and HAVING expression
- LIMIT and ORDER BY
- Table and column aliases (`SELECT t.column as col FROM "Table.csv" T`)

csql automatically converts values to integer and float types and nulls. 

See [this document](docs/syntax.md) for detailed description of Csql syntax.

---
<sup>*</sup> IN (select) is not supported yet in version 0.2.x


## Usage

For interactive mode run `csql` command and the CSQL shell will start.

You can type SELECT queries or commands. Queries must start with word `select` and terminate with a semicolon (`;`).

Supported commands are:
   - help - display help message
   - clear - clear screen
   - exit - exit csql shell
   - ls - list files in the current directory
   - pwd - show current directory
   - cd - change directory
   - csv - set CSV parsing options
   - inspect - inspect a CSV file - print out column names and data types of the first row
                    
CSQL can also execute queries passed from the command line

    $ csql -c "SELECT * FROM file.csv WHERE column = 'value'" -o result.csv

## Building

csql is a simple Go project, so it can be built as any other Go application

    $ go test ./...
    $ go build

Or use `build.sh` script to build for all supported platforms. Results will be placed in the `/bin` folder. 

### Changing grammar

csql uses ANTLR4 parser generator. ANTLR is a Java application and so you will need Java if you changed the grammar.
Run `antlr/generate.sh` from the project root to build parser from the `Csql.g4` grammar file.
