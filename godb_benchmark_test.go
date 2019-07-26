package db_sql_benchmark

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/postgresql"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

var mockDriver *sql.DB

func init() {
	db, _, _ := sqlmock.New()
	mockDriver = db
}

var (
	godbSqlite3    = godb.Wrap(sqlite.Adapter, mockDriver)
	godbPostgreSQL = godb.Wrap(postgresql.Adapter, mockDriver)
)

func godbSelectSimple(b *testing.B, builder *godb.DB) {
	for n := 0; n < b.N; n++ {

		_, _, err := builder.SelectFrom("tickets").
			Columns("id").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}

	}
}

func godbSelectConditional(b *testing.B, builder *godb.DB) {
	for n := 0; n < b.N; n++ {

		q := builder.SelectFrom("tickets").
			Columns("id").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		if n%2 == 0 {
			q = q.GroupBy("subdomain_id").
				Having("number = 1").
				OrderBy("state").
				Limit(7).
				Offset(8)
		}

		_, _, err := q.ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}

	}
}

func godbSelectComplex(b *testing.B, builder *godb.DB) {
	for n := 0; n < b.N; n++ {
		q := builder.SelectFrom("c").
			Columns("a", "b", "z", "y", "x").
			Distinct().
			Where("d = ? OR e = ?", 1, "wat").
			Where("g = ?", 3).
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
			Offset(8)
		_, _, err := q.ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}

	}
}

func BenchmarkGodbSelectSimple(b *testing.B) {
	godbSelectSimple(b, godbSqlite3)
}

func BenchmarkGodbSelectSimplePostgreSQL(b *testing.B) {
	godbSelectSimple(b, godbPostgreSQL)
}

func BenchmarkGodbSelectConditional(b *testing.B) {
	godbSelectConditional(b, godbSqlite3)
}

func BenchmarkGodbSelectConditionalPostgreSQL(b *testing.B) {
	godbSelectConditional(b, godbPostgreSQL)
}

func BenchmarkGodbSelectComplex(b *testing.B) {
	godbSelectComplex(b, godbSqlite3)
}

func BenchmarkGodbSelectComplexPostgreSQL(b *testing.B) {
	godbSelectComplex(b, godbPostgreSQL)
}

//
// Insert benchmark
//
func godbInsert(b *testing.B, builder *godb.DB) {
	for n := 0; n < b.N; n++ {
		q := builder.InsertInto("mytable").
			Columns("id", "a", "b", "price", "created", "updated").
			Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05")
		_, _, err := q.ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}

	}
}

func BenchmarkGodbInsert(b *testing.B) {
	godbInsert(b, godbSqlite3)
}

func BenchmarkGodbInsertPostgreSQL(b *testing.B) {
	godbInsert(b, godbPostgreSQL)
}

//
// Update benchmark
//
func godbUpdateSetColumns(b *testing.B, builder *godb.DB) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		q := builder.UpdateTable("mytable").
			Set("foo", 1).
			SetRaw("bar, (COALESCE(bar, 0) + 1)").
			Set("c", 2).
			Where("id = ?", 9)
		_, _, err := q.ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}
	}
}

func BenchmarkGodbUpdateSetColumns(b *testing.B) {
	godbUpdateSetColumns(b, godbSqlite3)
}

func BenchmarkGodbUpdateSetColumnsPostgreSQL(b *testing.B) {
	godbUpdateSetColumns(b, godbPostgreSQL)
}

//
// Delete benchmark
//
func godbDelete(b *testing.B, builder *godb.DB) {
	for n := 0; n < b.N; n++ {
		q := builder.DeleteFrom("test_table").
			Where("b = ?", 1)
		_, _, err := q.ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}
	}
}

func BenchmarkGodbDelete(b *testing.B) {
	godbDelete(b, godbSqlite3)
}

func BenchmarkGodbDeletePostgreSQL(b *testing.B) {
	godbDelete(b, godbPostgreSQL)
}
