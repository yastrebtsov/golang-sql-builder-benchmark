package db_sql_benchmark

import (
	"fmt"
	"testing"

	"github.com/leporo/sqlf"
)

var (
	sqlfSqlite3    = sqlf.NewBuilder(sqlf.NoDialect())
	sqlfPostgreSQL = sqlf.NewBuilder(sqlf.PostgreSQL())
)

func sqlfSelectSimple(b *testing.B, builder *sqlf.Builder) {
	for n := 0; n < b.N; n++ {
		q := builder.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")
		q.Build()
		q.Close()
	}
}

func sqlfSelectConditional(b *testing.B, builder *sqlf.Builder) {
	for n := 0; n < b.N; n++ {
		q := builder.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		if n%2 == 0 {
			q.GroupBy("subdomain_id").
				Having("number = ?", 1).
				OrderBy("state").
				Limit(7).
				Offset(8)
		}

		q.Build()
		q.Close()
	}
}

func sqlfSelectComplex(b *testing.B, builder *sqlf.Builder) {
	for n := 0; n < b.N; n++ {
		q := builder.Select("DITINCT a, b, z, y, x").
			// Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where("g = ?", 3).
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
			Offset(8)
		q.Build()
		q.Close()
	}
}

func sqlfSelectSubquery(b *testing.B, builder *sqlf.Builder) {
	for n := 0; n < b.N; n++ {
		sq := builder.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")
		subQuery, _ := sq.Build()

		q := builder.Select("DITINCT a, b").
			Select(fmt.Sprintf("(%s) AS subq", subQuery)).
			From("c").
			// Distinct().
			// Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where("g = ?", 3).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8)
		q.Build()
		q.Close()
		sq.Close()
	}
}

func BenchmarkSqlfSelectSimple(b *testing.B) {
	sqlfSelectSimple(b, sqlfSqlite3)
}

func BenchmarkSqlfSelectSimplePostgreSQL(b *testing.B) {
	sqlfSelectSimple(b, sqlfPostgreSQL)
}

func BenchmarkSqlfSelectConditional(b *testing.B) {
	sqlfSelectConditional(b, sqlfSqlite3)
}

func BenchmarkSqlfSelectConditionalPostgreSQL(b *testing.B) {
	sqlfSelectConditional(b, sqlfPostgreSQL)
}

func BenchmarkSqlfSelectComplex(b *testing.B) {
	sqlfSelectComplex(b, sqlfSqlite3)
}

func BenchmarkSqlfSelectComplexPostgreSQL(b *testing.B) {
	sqlfSelectComplex(b, sqlfPostgreSQL)
}

func BenchmarkSqlfSelectSubquery(b *testing.B) {
	sqlfSelectSubquery(b, sqlfSqlite3)
}

func BenchmarkSqlfSelectSubqueryPostgreSQL(b *testing.B) {
	sqlfSelectSubquery(b, sqlfPostgreSQL)
}

//
// Insert benchmark
//
func sqlfInsert(b *testing.B, builder *sqlf.Builder) {
	for n := 0; n < b.N; n++ {
		q := builder.InsertInto("mytable").
			Set("id", 1).
			Set("a", "test_a").
			Set("b", "test_b").
			Set("price", 100.05).
			Set("created", "2014-01-05").
			Set("updated", "2015-01-05")
		q.Build()
		q.Close()
	}
}

func BenchmarkSqlfInsert(b *testing.B) {
	sqlfInsert(b, sqlfSqlite3)
}

func BenchmarkSqlfInsertPostgreSQL(b *testing.B) {
	sqlfInsert(b, sqlfPostgreSQL)
}

//
// Update benchmark
//
func sqlfUpdateSetColumns(b *testing.B, builder *sqlf.Builder) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		q := builder.Update("mytable").
			Set("foo", 1).
			SetExpr("bar", "COALESCE(bar, 0) + 1").
			Set("c", 2).
			Where("id = ?", 9).
			Limit(10)
		q.Build()
		q.Close()
	}
}

func BenchmarkSqlfUpdateSetColumns(b *testing.B) {
	sqlfUpdateSetColumns(b, sqlfSqlite3)
}

func BenchmarkSqlfUpdateSetColumnsPostgreSQL(b *testing.B) {
	sqlfUpdateSetColumns(b, sqlfPostgreSQL)
}

//
// Delete benchmark
//
func sqlfDelete(b *testing.B, builder *sqlf.Builder) {
	for n := 0; n < b.N; n++ {
		q := builder.DeleteFrom("test_table").
			Where("b = ?", 1).
			Limit(2)
		q.Build()
		q.Close()
	}
}

func BenchmarkSqlfDelete(b *testing.B) {
	sqlfDelete(b, sqlfSqlite3)
}

func BenchmarkSqlfDeletePostgreSQL(b *testing.B) {
	sqlfDelete(b, sqlfPostgreSQL)
}
