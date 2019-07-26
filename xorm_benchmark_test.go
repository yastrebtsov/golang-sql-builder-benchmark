package db_sql_benchmark

import (
	"testing"

	"github.com/go-xorm/builder"
)

var (
	sqllDialect  = builder.Dialect(builder.SQLITE)
	pgSQLDialect = builder.Dialect(builder.POSTGRES)
)

func xormSelectSimple(b *testing.B, bld *builder.Builder) {
	for n := 0; n < b.N; n++ {
		_, _, err := bld.Select("id").From("tickets").
			Where(builder.Eq{"subdomain_id": 1}.And(builder.Eq{"state": "open"}.Or(builder.Eq{"state": "spam"}))).
			ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}

	}
}

func xormSelectConditional(b *testing.B, bld *builder.Builder) {
	for n := 0; n < b.N; n++ {

		q := bld.Select("id").From("tickets").
			Where(builder.Eq{"subdomain_id": 1}.And(builder.Eq{"state": "open"}.Or(builder.Eq{"state": "spam"})))

		if n%2 == 0 {
			q = q.GroupBy("subdomain_id").
				Having("number = 1").
				OrderBy("state").
				Limit(7, 8)
		}

		_, _, err := q.ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}

	}
}

func xormSelectComplex(b *testing.B, bld *builder.Builder) {
	for n := 0; n < b.N; n++ {
		q := bld.Select("DISTINCT a", "b", "z", "y", "x").
			From("c").
			Where(builder.Eq{"d": 1}.Or(builder.Eq{"e": "wat"})).
			Where(builder.Eq{"g": 3}).
			GroupBy("i").
			GroupBy("ii").
			GroupBy("iii").
			Having("j=k").
			Having("jj=1").
			Having("jjj=2").
			OrderBy("l").
			OrderBy("l").
			OrderBy("l").
			Limit(7, 8)
		_, _, err := q.ToSQL()
		if err != nil {
			b.Fatal("Error in sql", err)
		}

	}
}

// func xormSelectSubquery(b *testing.B, bld *builder.Builder) {
// 	for n := 0; n < b.N; n++ {
// 		sq := bld.Select("id").
// 			From("tickets").
// 			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")
// 		subQuery, _ := sq.ToBoundSQL()

// 		q := bld.Select("DISTINCT a, b").
// 			Select(fmt.Sprintf("(%s) AS subq", subQuery)).
// 			From("c").
// 			// Distinct().
// 			// Where(dbr.Eq{"f": 2, "x": "hi"}).
// 			Where("g = ?", 3).
// 			OrderBy("l").
// 			OrderBy("l").
// 			Limit(7).
// 			Offset(8)
// 			_, _, err := q.ToSQL()
// 	}
// }

func BenchmarkXormSelectSimple(b *testing.B) {
	xormSelectSimple(b, sqllDialect)
}

func BenchmarkXormSelectSimplePostgreSQL(b *testing.B) {
	xormSelectSimple(b, pgSQLDialect)
}

func BenchmarkXormSelectConditional(b *testing.B) {
	xormSelectConditional(b, sqllDialect)
}

func BenchmarkXormSelectConditionalPostgreSQL(b *testing.B) {
	xormSelectConditional(b, pgSQLDialect)
}

func BenchmarkXormSelectComplex(b *testing.B) {
	xormSelectComplex(b, sqllDialect)
}

func BenchmarkXormSelectComplexPostgreSQL(b *testing.B) {
	xormSelectComplex(b, pgSQLDialect)
}

// func BenchmarkXormSelectSubquery(b *testing.B) {
// 	xormSelectSubquery(b, sqllDialect)
// }

// func BenchmarkXormSelectSubqueryPostgreSQL(b *testing.B) {
// 	xormSelectSubquery(b, pgSQLDialect)
// }

//
// Insert benchmark
//
// func xormInsert(b *testing.B, builder *builder.Builder) {
// 	for n := 0; n < b.N; n++ {
// 		q := builder.InsertInto("mytable").
// 			Columns("id", "a", "b", "price", "created", "updated").
// 			Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05")
// 		_, _, err := q.ToSQL()
// 		if err != nil {
// 			b.Fatal("Error in sql", err)
// 		}

// 	}
// }

// func BenchmarkXormInsert(b *testing.B) {
// 	xormInsert(b, sqllDialect)
// }

// func BenchmarkXormInsertPostgreSQL(b *testing.B) {
// 	xormInsert(b, pgSQLDialect)
// }

//
// Update benchmark
//
// func xormUpdateSetColumns(b *testing.B, builder *builder.Builder) {

// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		q := builder.UpdateTable("mytable").
// 			Set("foo", 1).
// 			SetRaw("bar, (COALESCE(bar, 0) + 1)").
// 			Set("c", 2).
// 			Where("id = ?", 9)
// 		_, _, err := q.ToSQL()
// 		if err != nil {
// 			b.Fatal("Error in sql", err)
// 		}
// 	}
// }

// func BenchmarkXormUpdateSetColumns(b *testing.B) {
// 	xormUpdateSetColumns(b, sqllDialect)
// }

// func BenchmarkXormUpdateSetColumnsPostgreSQL(b *testing.B) {
// 	xormUpdateSetColumns(b, pgSQLDialect)
// }

//
// Delete benchmark
//
// func xormDelete(b *testing.B, builder *builder.Builder) {
// 	for n := 0; n < b.N; n++ {
// 		q := builder.DeleteFrom("test_table").
// 			Where("b = ?", 1)
// 		_, _, err := q.ToSQL()
// 		if err != nil {
// 			b.Fatal("Error in sql", err)
// 		}
// 	}
// }

// func BenchmarkXormDelete(b *testing.B) {
// 	xormDelete(b, sqllDialect)
// }

// func BenchmarkXormDeletePostgreSQL(b *testing.B) {
// 	xormDelete(b, pgSQLDialect)
// }
