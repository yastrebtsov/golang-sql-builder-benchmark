golang-sql-builder-benchmark
====================

A comparison of popular Go SQL query builders. Provides feature list and benchmarks

# Builders

1. dbr: https://github.com/gocraft/dbr
2. squirrel: https://github.com/lann/squirrel
3. sqrl: https://github.com/elgris/sqrl
4. gocu: https://github.com/doug-martin/goqu - just for SELECT query
5. sqlf: https://github.com/leporo/sqlf


# Feature list

| feature                    | dbr | squirrel | sqrl | goqu | sqlf |
|----------------------------|-----|----------|------|------|------|
| SelectBuilder              | +   | +        | +    | +    | +    |
| InsertBuilder              | +   | +        | +    | +    | +    |
| UpdateBuilder              | +   | +        | +    | +    | +    |
| DeleteBuilder              | +   | +        | +    | +    | +    |
| PostgreSQL support         | +   | +        | +    | +    | +    |
| Custom placeholders        | +   | +        | +    | +    | +    |
| JOINs support              | +   | +        | +    | +    |      |
| Subquery in query builder  |     | +        | +    | +    |      |
| Aliases for columns        |     | +        | +    | +    |      |
| CASE expression            |     | +        | +    | +    |      |

Some explanations here:
- `Custom placeholders` - ability to use not only `?` placeholders, Useful for PostgreSQL
- `JOINs support` - ability to build JOINs in SELECT queries like `Select("*").From("a").Join("b")`
- `Subquery in query builder` - when you prepare a subquery with one builder and then pass it to another. Something like:
```go
subQ := Select("aa", "bb").From("dd")
qb := Select().Column(subQ).From("a")
```
- `Aliases for columns` - easy way to alias a column, especially if column is specified by subquery:
```go
subQ := Select("aa", "bb").From("dd")
qb := Select().Column(Alias(subQ, "alias")).From("a")
```
- `CASE expression` - syntactic sugar for [CASE expressions](http://dev.mysql.com/doc/refman/5.7/en/case.html)

# Benchmarks

`go test -bench=. -benchmem | column -t` on 2.6 GHz i5 Macbook Pro:

```
BenchmarkDbrSelectSimple-4                  2000000     812      ns/op  728    B/op  12   allocs/op
BenchmarkDbrSelectSimplePostgreSQL-4        2000000     1015     ns/op  728    B/op  12   allocs/op
BenchmarkDbrSelectConditional-4             1000000     1195     ns/op  976    B/op  17   allocs/op
BenchmarkDbrSelectConditionalPostgreSQL-4   1000000     1202     ns/op  976    B/op  17   allocs/op
BenchmarkDbrSelectComplex-4                 500000      4729     ns/op  2360   B/op  38   allocs/op
BenchmarkDbrSelectComplexPostgreSQL-4       500000      3401     ns/op  2360   B/op  38   allocs/op
BenchmarkDbrSelectSubquery-4                1000000     4212     ns/op  2160   B/op  29   allocs/op
BenchmarkDbrSelectSubqueryPostgreSQL-4      500000      3812     ns/op  2160   B/op  29   allocs/op
BenchmarkDbrInsert-4                        1000000     2737     ns/op  1136   B/op  26   allocs/op
BenchmarkDbrInsertPostgreSQL-4              500000      3253     ns/op  1136   B/op  26   allocs/op
BenchmarkDbrUpdateSetColumns-4              1000000     2979     ns/op  1297   B/op  28   allocs/op
BenchmarkDbrUpdateSetColumnsPostgreSQL-4    500000      3914     ns/op  1297   B/op  28   allocs/op
BenchmarkDbrUpdateSetMap-4                  500000      2756     ns/op  1296   B/op  28   allocs/op
BenchmarkDbrUpdateSetMapPostgreSQL-4        500000      2832     ns/op  1296   B/op  28   allocs/op
BenchmarkDbrDelete-4                        2000000     819      ns/op  496    B/op  12   allocs/op
BenchmarkDbrDeletePostgreSQL-4              2000000     838      ns/op  496    B/op  12   allocs/op


BenchmarkGoquSelectSimple-4                 300000      5610     ns/op  3360   B/op  38   allocs/op
BenchmarkGoquSelectConditional-4            200000      6043     ns/op  3804   B/op  49   allocs/op
BenchmarkGoquSelectComplex-4                100000      19087    ns/op  9464   B/op  169  allocs/op


BenchmarkSqlfSelectSimple-4                 3000000     405      ns/op  80     B/op  1    allocs/op
BenchmarkSqlfSelectSimplePostgreSQL-4       2000000     853      ns/op  224    B/op  4    allocs/op
BenchmarkSqlfSelectConditional-4            1000000     1024     ns/op  148    B/op  3    allocs/op
BenchmarkSqlfSelectConditionalPostgreSQL-4  1000000     1539     ns/op  280    B/op  6    allocs/op
BenchmarkSqlfSelectComplex-4                1000000     2195     ns/op  192    B/op  5    allocs/op
BenchmarkSqlfSelectComplexPostgreSQL-4      500000      2457     ns/op  208    B/op  5    allocs/op
BenchmarkSqlfSelectSubquery-4               1000000     1845     ns/op  384    B/op  8    allocs/op
BenchmarkSqlfSelectSubqueryPostgreSQL-4     1000000     2346     ns/op  528    B/op  11   allocs/op
BenchmarkSqlfInsert-4                       1000000     1355     ns/op  96     B/op  1    allocs/op
BenchmarkSqlfInsertPostgreSQL-4             1000000     1578     ns/op  96     B/op  1    allocs/op
BenchmarkSqlfUpdateSetColumns-4             2000000     1039     ns/op  136    B/op  6    allocs/op
BenchmarkSqlfUpdateSetColumnsPostgreSQL-4   1000000     1372     ns/op  136    B/op  6    allocs/op
BenchmarkSqlfDelete-4                       2000000     557      ns/op  64     B/op  3    allocs/op
BenchmarkSqlfDeletePostgreSQL-4             2000000     625      ns/op  64     B/op  3    allocs/op


BenchmarkSqrlSelectSimple-4                 1000000     1139     ns/op  704    B/op  15   allocs/op
BenchmarkSqrlSelectConditional-4            1000000     1485     ns/op  848    B/op  18   allocs/op
BenchmarkSqrlSelectComplex-4                200000      8107     ns/op  4352   B/op  87   allocs/op
BenchmarkSqrlSelectSubquery-4               300000      6371     ns/op  3352   B/op  65   allocs/op
BenchmarkSqrlSelectMoreComplex-4            100000      12307    ns/op  6961   B/op  138  allocs/op
BenchmarkSqrlInsert-4                       1000000     1771     ns/op  992    B/op  17   allocs/op
BenchmarkSqrlUpdateSetColumns-4             500000      2129     ns/op  1056   B/op  25   allocs/op
BenchmarkSqrlUpdateSetMap-4                 500000      2621     ns/op  1130   B/op  27   allocs/op
BenchmarkSqrlDelete-4                       2000000     615      ns/op  304    B/op  9    allocs/op


BenchmarkSquirrelSelectSimple-4             300000      6831     ns/op  2512   B/op  49   allocs/op
BenchmarkSquirrelSelectConditional-4        200000      13046    ns/op  3756   B/op  79   allocs/op
BenchmarkSquirrelSelectComplex-4            50000       31387    ns/op  10162  B/op  224  allocs/op
BenchmarkSquirrelSelectSubquery-4           50000       25838    ns/op  8642   B/op  182  allocs/op
BenchmarkSquirrelSelectMoreComplex-4        30000       50207    ns/op  17155  B/op  384  allocs/op
BenchmarkSquirrelInsert-4                   200000      8589     ns/op  3040   B/op  67   allocs/op
BenchmarkSquirrelUpdateSetColumns-4         100000      12964    ns/op  4464   B/op  103  allocs/op
BenchmarkSquirrelUpdateSetMap-4             100000      13674    ns/op  4544   B/op  105  allocs/op
BenchmarkSquirrelDelete-4                   200000      8116     ns/op  2592   B/op  63   allocs/op
```

# Conclusion

If your queries are very simple, pick `dbr`, the fastest one.

If really need immutability of query builder and you're ready to sacrifice extra memory, use `squirrel`, the slowest but most reliable one.

If you like those sweet helpers that `squirrel` provides to ease your query building or if you plan to use the same builder for `PostgreSQL`, take `sqrl` as it's balanced between performance and features.

`goqu` has LOTS of features and ways to build queries. Although it requires stubbing sql connection if you need just to build a query. It can be done with [sqlmock](http://github.com/DATA-DOG/go-sqlmock). Disadvantage: the builder is slow and has TOO MANY features, so building a query may become a nightmare. But if you need total control on everything - this is your choice.