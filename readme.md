# csql - Query language for comma-separated values

csql is a small tool allowing to run SQL queries on CSV files.

csql mostly adheres to ANSI SQL syntax<sup> * </sup> and supports all basic operations, including
joins and aggregations<sup> ** </sup>

---
<sup> * </sup> Only SELECT queries are supported, not every ANSI SQL expression and function is supported  

<sup>**</sup> Full outer join and aggregations are not supported yet in version 0.1.x