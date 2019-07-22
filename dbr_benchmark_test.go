package db_sql_benchmark

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	dbrDialect "github.com/gocraft/dbr/dialect"
)

//
// Select benchmarks
//

func dbrToSQL(dialect dbr.Dialect, b dbr.Builder) (query string, args []interface{}) {
	// As ToSql method seems to be dropped, we use a trimmed version
	// of interpolator.encodePlaceholder method dbr calls under the hood.
	pbuf := dbr.NewBuffer()
	b.Build(dialect, pbuf)
	return pbuf.String(), pbuf.Value()
}

func dbrSelectSimple(b *testing.B, dialect dbr.Dialect) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dbrToSQL(dialect, dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam"))
	}
}

func dbrSelectConditional(b *testing.B, dialect dbr.Dialect) {

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		qb := dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		if n%2 == 0 {
			qb.GroupBy("subdomain_id").
				Having("number = ?", 1).
				OrderBy("state").
				Limit(7).
				Offset(8)
		}

		dbrToSQL(dialect, qb)
	}
}
func dbrSelectComplex(b *testing.B, dialect dbr.Dialect) {

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dbrToSQL(dialect, dbr.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			// Where(dbr.Eq{"h": []int{1, 2, 3}}).
			GroupBy("i").
			GroupBy("ii").
			GroupBy("iii").
			Having("j = k").
			Having("jj = ?", 1).
			Having("jjj = ?", 2).
			OrderBy("l").
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8))
	}
}

func dbrSelectSubquery(b *testing.B, dialect dbr.Dialect) {

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		subQuery, _ := dbrToSQL(dialect, dbr.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam"))

		dbrToSQL(dialect, dbr.Select("a", "b", fmt.Sprintf("(%s) AS subq", subQuery)).
			From("c").
			Distinct().
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8))
	}
}

func BenchmarkDbrSelectSimple(b *testing.B) {
	dbrSelectSimple(b, dbrDialect.SQLite3)
}

func BenchmarkDbrSelectSimplePostgreSQL(b *testing.B) {
	dbrSelectSimple(b, dbrDialect.PostgreSQL)
}

func BenchmarkDbrSelectConditional(b *testing.B) {
	dbrSelectConditional(b, dbrDialect.SQLite3)
}

func BenchmarkDbrSelectConditionalPostgreSQL(b *testing.B) {
	dbrSelectConditional(b, dbrDialect.PostgreSQL)
}

func BenchmarkDbrSelectComplex(b *testing.B) {
	dbrSelectComplex(b, dbrDialect.SQLite3)
}

func BenchmarkDbrSelectComplexPostgreSQL(b *testing.B) {
	dbrSelectComplex(b, dbrDialect.PostgreSQL)
}

func BenchmarkDbrSelectSubquery(b *testing.B) {
	dbrSelectSubquery(b, dbrDialect.SQLite3)
}

func BenchmarkDbrSelectSubqueryPostgreSQL(b *testing.B) {
	dbrSelectSubquery(b, dbrDialect.PostgreSQL)
}

//
// Insert benchmark
//
func dbrInsert(b *testing.B, dialect dbr.Dialect) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dialect, dbr.InsertInto("mytable").
			Columns("id", "a", "b", "price", "created", "updated").
			Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05"))
	}
}

func BenchmarkDbrInsert(b *testing.B) {
	dbrInsert(b, dbrDialect.SQLite3)
}

func BenchmarkDbrInsertPostgreSQL(b *testing.B) {
	dbrInsert(b, dbrDialect.PostgreSQL)
}

//
// Update benchmark
//
func dbrUpdateSetColumns(b *testing.B, dialect dbr.Dialect) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dialect, dbr.Update("mytable").
			Set("foo", 1).
			Set("bar", dbr.Expr("COALESCE(bar, 0) + 1")).
			Set("c", 2).
			Where("id = ?", 9).
			Limit(10))
	}
}

func dbrUpdateSetMap(b *testing.B, dialect dbr.Dialect) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dialect, dbr.Update("mytable").
			SetMap(map[string]interface{}{"b": 1, "c": 2, "bar": dbr.Expr("COALESCE(bar, 0) + 1")}).
			Where("id = ?", 9).
			Limit(10))
	}
}

func BenchmarkDbrUpdateSetColumns(b *testing.B) {
	dbrUpdateSetColumns(b, dbrDialect.SQLite3)
}

func BenchmarkDbrUpdateSetColumnsPostgreSQL(b *testing.B) {
	dbrUpdateSetColumns(b, dbrDialect.PostgreSQL)
}

func BenchmarkDbrUpdateSetMap(b *testing.B) {
	dbrUpdateSetMap(b, dbrDialect.SQLite3)
}

func BenchmarkDbrUpdateSetMapPostgreSQL(b *testing.B) {
	dbrUpdateSetMap(b, dbrDialect.SQLite3)
}

//
// Delete benchmark
//
func dbrDelete(b *testing.B, dialect dbr.Dialect) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		dbrToSQL(dialect, dbr.DeleteFrom("test_table").
			Where("b = ?", 1).
			Limit(2))
	}
}

func BenchmarkDbrDelete(b *testing.B) {
	dbrDelete(b, dbrDialect.SQLite3)
}

func BenchmarkDbrDeletePostgreSQL(b *testing.B) {
	dbrDelete(b, dbrDialect.PostgreSQL)
}
