# Csql Syntax

Csql only supports SELECT statement

## Syntax

Reserved words are typed in ALL CAPS. Expressions in square brackets are optional. 
Asterisk `*` means that previous expression can be repeated zero or more times.
Pipe symbol `|` between two expressions means that either left or right expression can be used.


```
SELECT projection FROM  
    table [, table]* | join
WHERE
    conditions
[ UNION ...]
ORDER BY order_by
LIMIT number
[INTO destination]
```

```
projection:
    [distinct] column [,column]* 

column:
    * | (column_name | expression AS column_alias)
    
column_name:
    [table_name.]file_column_name
```
_expression_ is anything that can be evaluated to a value, i.e. 1+2, 'string' or a function <sup>*</sup>
Asterisk can be prepended with table name to select all columns from that table.

```
table:
    file_name [AS] table_alias
```
_table_ is the name of the CSV file. If the file name includes special characters or file path it must be enclosed 
in double quotes.

Several tables listed with separating comma are equivalent of cross joining these tables. 

```
join:
   join_source CROSS JOIN join_source |
   join_source join_expression join_source ON conditions 
   
join_source:
   table | join 
```
_join_expression_ is one of

- INNER JOIN
- LEFT [OUTER] JOIN
- RIGHT [OUTER] JOIN
- FULL [OUTER] JOIN

CROSS JOIN does not support `ON` clause.

```
conditions:
    expression |
    '(' expression ') |
    expression =|<>|!=|>|<|>=|<= expression
    expression [NOT] LIKE like_pattern
    expression [NOT] MATCH regular_expression
    expression BETWEEN expression AND expression
    expression OR expression
    expression AND expression
    NOT expression
```
Conditions must always evaluate to either TRUE or FALSE.

```
order_by:
   order_by_expr [, order_by_expr]*
   
order_by_expr:
    column_name | column_index [ASC|DESC]

```

## Literals

Csql supports following literals:
- numbers
- strings
- NULL

Numbers can be either integer or floating points and can include unary `+` or `-` sign.

Strings are enclosed in single quotes `'`.

## Expressions

Supported expressions include literals, columns, binary arithmetic operations on other expressions,
string concatenation operator `||` and function calls

Supported operations:
 - `+` addition
 - `-` subtraction
 - `*` multiplication
 - `/` division
 - `%` modulo
 - `||` string concatenation

Supported functions:
 - Math
   - [x] ROUND
   - [x] TRUNC
   - [x] FRAC
   - [x] POW
   - [x] SQRT
 - Type cast:
   - [x] TO_STRING
   - [x] TO_FLOAT
   - [x] TO_INT
 - String:
   - [x] SUBSTRING
   - [ ] INDEX_OF
   - [x] LEN
   - [x] TO_UPPER
   - [x] TO_LOWER
 - Others
   - [x] COALESCE
                       
Supported Aggregate functions:
 - SUM
 - COUNT
 - MIN
 - MAX
 - AVG

See more details on function syntax in the [Functions](#functions) section.

## CSV format

First line of the file must contain column names. Strings that contain commas, must be enclosed in double quotes.
Whole file is read into memory at the moment it is declared in the FROM clause.

## Data types

Csql supports integers, floating point numbers and strings. Type of each value is determined at the time of reading of
CSV file. Operations are type-sensitive, i.e., for example, number cannot be added to a string.

Special value `null` is read as NULL value. The string representing null value can be changed using shell command `cvs --null="null value"` or
command line parameter `--null`.

## Indices

As CSV files does not have indices, Csql does not support indices. All operations require full table scan and joins have
M*N time complexity. In short, don't expect fabulous performance, especially on large files.

## Functions

TBD
