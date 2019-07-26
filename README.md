golang-sql-builder-benchmark
====================

A comparison of popular Go SQL query builders. Provides feature list and benchmarks

# Builders

1. dbr: https://github.com/gocraft/dbr
2. squirrel: https://github.com/lann/squirrel
3. sqrl: https://github.com/elgris/sqrl
4. gocu: https://github.com/doug-martin/goqu - just for SELECT query
5. sqlf: https://github.com/leporo/sqlf
6. godb: https://github.com/samonzeweb/godb
7. xorm-builder: https://github.com/go-xorm/builder


# Feature list

| feature                    | dbr | squirrel | sqrl | goqu | sqlf | godb | xorm-builder |
|----------------------------|-----|----------|------|------|------|------|--------------|
| SelectBuilder              | +   | +        | +    | +    | +    | +    | +            |
| InsertBuilder              | +   | +        | +    | +    | +    | +    | +            |
| UpdateBuilder              | +   | +        | +    | +    | +    | +    | +            |
| DeleteBuilder              | +   | +        | +    | +    | +    | +    | +            |
| PostgreSQL support         | +   | +        | +    | +    | +    | +    | +            |
| MS SQL support             | -   | -        | -    | -    | -    | +    | +            |
| Custom placeholders        | +   | +        | +    | +    | +    | -    | +            |
| JOINs support              | +   | +        | +    | +    | +    | +    | +            |
| Subquery in query builder  |     | +        | +    | +    | -    | -*   | -*           |
| Aliases for columns        |     | +        | +    | +    | -    | -    | -            | 
| CASE expression            |     | +        | +    | +    | -    | -    | -            |
\* subqueries could be implemented only via raw/bounded SQL

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

`go test -bench=. -benchmem | column -t` on 2.2 GHz i7 Macbook Pro:

```
BenchmarkDbrSelectSimple-12                  3000000                                         537       ns/op  728       B/op  12      allocs/op
BenchmarkDbrSelectSimplePostgreSQL-12        3000000                                         545       ns/op  728       B/op  12      allocs/op
BenchmarkDbrSelectConditional-12             2000000                                         799       ns/op  976       B/op  17      allocs/op
BenchmarkDbrSelectConditionalPostgreSQL-12   2000000                                         807       ns/op  976       B/op  17      allocs/op
BenchmarkDbrSelectComplex-12                 1000000                                         2110      ns/op  2360      B/op  38      allocs/op
BenchmarkDbrSelectComplexPostgreSQL-12       1000000                                         2109      ns/op  2360      B/op  38      allocs/op
BenchmarkDbrSelectSubquery-12                1000000                                         1479      ns/op  2161      B/op  29      allocs/op
BenchmarkDbrSelectSubqueryPostgreSQL-12      1000000                                         1477      ns/op  2161      B/op  29      allocs/op
BenchmarkDbrInsert-12                        1000000                                         1375      ns/op  1136      B/op  26      allocs/op
BenchmarkDbrInsertPostgreSQL-12              1000000                                         1406      ns/op  1136      B/op  26      allocs/op
BenchmarkDbrUpdateSetColumns-12              1000000                                         1521      ns/op  1297      B/op  28      allocs/op
BenchmarkDbrUpdateSetColumnsPostgreSQL-12    1000000                                         1530      ns/op  1297      B/op  28      allocs/op
BenchmarkDbrUpdateSetMap-12                  1000000                                         1652      ns/op  1296      B/op  28      allocs/op
BenchmarkDbrUpdateSetMapPostgreSQL-12        1000000                                         1652      ns/op  1296      B/op  28      allocs/op
BenchmarkDbrDelete-12                        3000000                                         526       ns/op  496       B/op  12      allocs/op
BenchmarkDbrDeletePostgreSQL-12              3000000                                         525       ns/op  496       B/op  12      allocs/op
BenchmarkGodbSelectSimple-12                 2000000                                         853       ns/op  760       B/op  15      allocs/op
BenchmarkGodbSelectSimplePostgreSQL-12       2000000                                         858       ns/op  760       B/op  15      allocs/op
BenchmarkGodbSelectConditional-12            1000000                                         1185      ns/op  1020      B/op  20      allocs/op
BenchmarkGodbSelectConditionalPostgreSQL-12  1000000                                         1162      ns/op  1020      B/op  20      allocs/op
BenchmarkGodbSelectComplex-12                500000                                          2992      ns/op  2072      B/op  52      allocs/op
BenchmarkGodbSelectComplexPostgreSQL-12      500000                                          3000      ns/op  2072      B/op  52      allocs/op
BenchmarkGodbInsert-12                       2000000                                         910       ns/op  1008      B/op  10      allocs/op
BenchmarkGodbInsertPostgreSQL-12             2000000                                         903       ns/op  1008      B/op  10      allocs/op
BenchmarkGodbUpdateSetColumns-12             2000000                                         850       ns/op  656       B/op  18      allocs/op
BenchmarkGodbUpdateSetColumnsPostgreSQL-12   2000000                                         853       ns/op  656       B/op  18      allocs/op
BenchmarkGodbDelete-12                       3000000                                         421       ns/op  338       B/op  10      allocs/op
BenchmarkGodbDeletePostgreSQL-12             3000000                                         425       ns/op  338       B/op  10      allocs/op
BenchmarkGoquSelectSimple-12                 500000                                          3445      ns/op  3360      B/op  38      allocs/op
BenchmarkGoquSelectConditional-12            300000                                          3938      ns/op  3804      B/op  49      allocs/op
BenchmarkGoquSelectComplex-12                200000                                          11977     ns/op  9464      B/op  169     allocs/op
BenchmarkSqlfSelectSimple-12                 5000000                                         267       ns/op  0         B/op  0       allocs/op
BenchmarkSqlfSelectSimplePostgreSQL-12       3000000                                         395       ns/op  0         B/op  0       allocs/op
BenchmarkSqlfSelectConditional-12            3000000                                         428       ns/op  8         B/op  1       allocs/op
BenchmarkSqlfSelectConditionalPostgreSQL-12  3000000                                         593       ns/op  8         B/op  1       allocs/op
BenchmarkSqlfSelectComplex-12                2000000                                         735       ns/op  16        B/op  2       allocs/op
BenchmarkSqlfSelectComplexPostgreSQL-12      2000000                                         945       ns/op  16        B/op  2       allocs/op
BenchmarkSqlfSelectSubquery-12               2000000                                         943       ns/op  128       B/op  4       allocs/op
BenchmarkSqlfSelectSubqueryPostgreSQL-12     1000000                                         1132      ns/op  128       B/op  4       allocs/op
BenchmarkSqlfInsert-12                       2000000                                         819       ns/op  0         B/op  0       allocs/op
BenchmarkSqlfInsertPostgreSQL-12             2000000                                         959       ns/op  0         B/op  0       allocs/op
BenchmarkSqlfUpdateSetColumns-12             3000000                                         482       ns/op  8         B/op  1       allocs/op
BenchmarkSqlfUpdateSetColumnsPostgreSQL-12   2000000                                         624       ns/op  8         B/op  1       allocs/op
BenchmarkSqlfDelete-12                       5000000                                         282       ns/op  8         B/op  1       allocs/op
BenchmarkSqlfDeletePostgreSQL-12             5000000                                         342       ns/op  8         B/op  1       allocs/op
BenchmarkSqrlSelectSimple-12                 2000000                                         783       ns/op  704       B/op  15      allocs/op
BenchmarkSqrlSelectConditional-12            1000000                                         1020      ns/op  848       B/op  18      allocs/op
BenchmarkSqrlSelectComplex-12                300000                                          5707      ns/op  4354      B/op  87      allocs/op
BenchmarkSqrlSelectSubquery-12               300000                                          4093      ns/op  3354      B/op  65      allocs/op
BenchmarkSqrlSelectMoreComplex-12            200000                                          8623      ns/op  6963      B/op  138     allocs/op
BenchmarkSqrlInsert-12                       1000000                                         1187      ns/op  992       B/op  17      allocs/op
BenchmarkSqrlUpdateSetColumns-12             1000000                                         1476      ns/op  1056      B/op  25      allocs/op
BenchmarkSqrlUpdateSetMap-12                 1000000                                         1784      ns/op  1131      B/op  27      allocs/op
BenchmarkSqrlDelete-12                       3000000                                         412       ns/op  304       B/op  9       allocs/op
BenchmarkSquirrelSelectSimple-12             300000                                          4055      ns/op  2512      B/op  49      allocs/op
BenchmarkSquirrelSelectConditional-12        200000                                          6524      ns/op  3757      B/op  79      allocs/op
BenchmarkSquirrelSelectComplex-12            100000                                          17951     ns/op  10164     B/op  224     allocs/op
BenchmarkSquirrelSelectSubquery-12           100000                                          14967     ns/op  8645      B/op  182     allocs/op
BenchmarkSquirrelSelectMoreComplex-12        50000                                           30914     ns/op  17160     B/op  384     allocs/op
BenchmarkSquirrelInsert-12                   300000                                          4899      ns/op  3041      B/op  67      allocs/op
BenchmarkSquirrelUpdateSetColumns-12         200000                                          7631      ns/op  4466      B/op  103     allocs/op
BenchmarkSquirrelUpdateSetMap-12             200000                                          8178      ns/op  4546      B/op  105     allocs/op
BenchmarkSquirrelDelete-12                   300000                                          4799      ns/op  2593      B/op  63      allocs/op
BenchmarkXormSelectSimple-12                 10000                                           9767864   ns/op  3582102   B/op  61294   allocs/op
BenchmarkXormSelectSimplePostgreSQL-12       10000                                           14201616  ns/op  5593382   B/op  107232  allocs/op
BenchmarkXormSelectConditional-12            50                                              20500669  ns/op  7000353   B/op  121617  allocs/op
BenchmarkXormSelectConditionalPostgreSQL-12  50                                              28470866  ns/op  10904631  B/op  212797  allocs/op
BenchmarkXormSelectComplex-12                100                                             20151201  ns/op  7115083   B/op  122536  allocs/op
BenchmarkXormSelectComplexPostgreSQL-12      50                                              28509490  ns/op  10938281  B/op  213873  allocs/op
```

# Conclusion

If your queries are very simple, pick `dbr`, the fastest one.

If really need immutability of query builder and you're ready to sacrifice extra memory, use `squirrel`, the slowest but most reliable one.

If you like those sweet helpers that `squirrel` provides to ease your query building or if you plan to use the same builder for `PostgreSQL`, take `sqrl` as it's balanced between performance and features.

`goqu` has LOTS of features and ways to build queries. Although it requires stubbing sql connection if you need just to build a query. It can be done with [sqlmock](http://github.com/DATA-DOG/go-sqlmock). Disadvantage: the builder is slow and has TOO MANY features, so building a query may become a nightmare. But if you need total control on everything - this is your choice.