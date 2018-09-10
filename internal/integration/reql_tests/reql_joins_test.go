// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/gorethink/gorethink.v4"
	"gopkg.in/gorethink/gorethink.v4/internal/compare"
)

// Tests that manipulation data in tables
func TestJoinsSuite(t *testing.T) {
	suite.Run(t, new(JoinsSuite))
}

type JoinsSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *JoinsSuite) SetupTest() {
	suite.T().Log("Setting up JoinsSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_joins").Exec(suite.session)
	err = r.DBCreate("db_joins").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Wait().Exec(suite.session)
	suite.Require().NoError(err)

	r.DB("db_joins").TableDrop("messages").Exec(suite.session)
	err = r.DB("db_joins").TableCreate("messages").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Table("messages").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_joins").TableDrop("otable_test_joins").Exec(suite.session)
	err = r.DB("db_joins").TableCreate("otable_test_joins").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Table("otable_test_joins").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_joins").TableDrop("otable_test_joins2").Exec(suite.session)
	err = r.DB("db_joins").TableCreate("otable_test_joins2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Table("otable_test_joins2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_joins").TableDrop("receivers").Exec(suite.session)
	err = r.DB("db_joins").TableCreate("receivers").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Table("receivers").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_joins").TableDrop("senders").Exec(suite.session)
	err = r.DB("db_joins").TableCreate("senders").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Table("senders").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_joins").TableDrop("table_test_joins").Exec(suite.session)
	err = r.DB("db_joins").TableCreate("table_test_joins").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Table("table_test_joins").Wait().Exec(suite.session)
	suite.Require().NoError(err)
	r.DB("db_joins").TableDrop("table_test_joins2").Exec(suite.session)
	err = r.DB("db_joins").TableCreate("table_test_joins2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_joins").Table("table_test_joins2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *JoinsSuite) TearDownSuite() {
	suite.T().Log("Tearing down JoinsSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("db_joins").TableDrop("messages").Exec(suite.session)
		r.DB("db_joins").TableDrop("otable_test_joins").Exec(suite.session)
		r.DB("db_joins").TableDrop("otable_test_joins2").Exec(suite.session)
		r.DB("db_joins").TableDrop("receivers").Exec(suite.session)
		r.DB("db_joins").TableDrop("senders").Exec(suite.session)
		r.DB("db_joins").TableDrop("table_test_joins").Exec(suite.session)
		r.DB("db_joins").TableDrop("table_test_joins2").Exec(suite.session)
		r.DBDrop("db_joins").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *JoinsSuite) TestCases() {
	suite.T().Log("Running JoinsSuite: Tests that manipulation data in tables")

	messages := r.DB("db_joins").Table("messages")
	_ = messages // Prevent any noused variable errors
	otable_test_joins := r.DB("db_joins").Table("otable_test_joins")
	_ = otable_test_joins // Prevent any noused variable errors
	otable_test_joins2 := r.DB("db_joins").Table("otable_test_joins2")
	_ = otable_test_joins2 // Prevent any noused variable errors
	receivers := r.DB("db_joins").Table("receivers")
	_ = receivers // Prevent any noused variable errors
	senders := r.DB("db_joins").Table("senders")
	_ = senders // Prevent any noused variable errors
	table_test_joins := r.DB("db_joins").Table("table_test_joins")
	_ = table_test_joins // Prevent any noused variable errors
	table_test_joins2 := r.DB("db_joins").Table("table_test_joins2")
	_ = table_test_joins2 // Prevent any noused variable errors

	{
		// joins.yaml line #7
		/* partial({'tables_created':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_created": 1})
		/* r.db('test').table_create('test3', primary_key='foo') */

		suite.T().Log("About to run line #7: r.DB('test').TableCreate('test3').OptArgs(r.TableCreateOpts{PrimaryKey: 'foo', })")

		runAndAssert(suite.Suite, expected_, r.DB("db_joins").TableCreate("test3").OptArgs(r.TableCreateOpts{PrimaryKey: "foo"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	// joins.yaml line #11
	// table_test_joins3 = r.db('test').table('test3')
	suite.T().Log("Possibly executing: var table_test_joins3 r.Term = r.DB('test').Table('test3')")

	table_test_joins3 := r.DB("db_joins").Table("test3")
	_ = table_test_joins3 // Prevent any noused variable errors

	{
		// joins.yaml line #13
		/* partial({'errors':0, 'inserted':100}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 100})
		/* table_test_joins.insert(r.range(0, 100).map({'id':r.row, 'a':r.row % 4})) */

		suite.T().Log("About to run line #13: table_test_joins.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{'id': r.Row, 'a': r.Row.Mod(4), }))")

		runAndAssert(suite.Suite, expected_, table_test_joins.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{"id": r.Row, "a": r.Row.Mod(4)})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #13")
	}

	{
		// joins.yaml line #18
		/* partial({'errors':0, 'inserted':100}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 100})
		/* table_test_joins2.insert(r.range(0, 100).map({'id':r.row, 'b':r.row % 4})) */

		suite.T().Log("About to run line #18: table_test_joins2.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{'id': r.Row, 'b': r.Row.Mod(4), }))")

		runAndAssert(suite.Suite, expected_, table_test_joins2.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{"id": r.Row, "b": r.Row.Mod(4)})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #18")
	}

	{
		// joins.yaml line #23
		/* partial({'errors':0, 'inserted':100}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 100})
		/* table_test_joins3.insert(r.range(0, 100).map({'foo':r.row, 'b':r.row % 4})) */

		suite.T().Log("About to run line #23: table_test_joins3.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{'foo': r.Row, 'b': r.Row.Mod(4), }))")

		runAndAssert(suite.Suite, expected_, table_test_joins3.Insert(r.Range(0, 100).Map(map[interface{}]interface{}{"foo": r.Row, "b": r.Row.Mod(4)})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// joins.yaml line #28
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* otable_test_joins.insert(r.range(1,100).map({'id': r.row, 'a': r.row})) */

		suite.T().Log("About to run line #28: otable_test_joins.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{'id': r.Row, 'a': r.Row, }))")

		runAndAssert(suite.Suite, expected_, otable_test_joins.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{"id": r.Row, "a": r.Row})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// joins.yaml line #29
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* otable_test_joins2.insert(r.range(1,100).map({'id': r.row, 'b': 2 * r.row})) */

		suite.T().Log("About to run line #29: otable_test_joins2.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{'id': r.Row, 'b': r.Mul(2, r.Row), }))")

		runAndAssert(suite.Suite, expected_, otable_test_joins2.Insert(r.Range(1, 100).Map(map[interface{}]interface{}{"id": r.Row, "b": r.Mul(2, r.Row)})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #29")
	}

	// joins.yaml line #34
	// ij = table_test_joins.inner_join(table_test_joins2, lambda x,y:x['a'] == y['b']).zip()
	suite.T().Log("Possibly executing: var ij r.Term = table_test_joins.InnerJoin(table_test_joins2, func(x r.Term, y r.Term) interface{} { return x.AtIndex('a').Eq(y.AtIndex('b'))}).Zip()")

	ij := table_test_joins.InnerJoin(table_test_joins2, func(x r.Term, y r.Term) interface{} { return x.AtIndex("a").Eq(y.AtIndex("b")) }).Zip()
	_ = ij // Prevent any noused variable errors

	{
		// joins.yaml line #37
		/* 2500 */
		var expected_ int = 2500
		/* ij.count() */

		suite.T().Log("About to run line #37: ij.Count()")

		runAndAssert(suite.Suite, expected_, ij.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// joins.yaml line #39
		/* 0 */
		var expected_ int = 0
		/* ij.filter(lambda row:row['a'] != row['b']).count() */

		suite.T().Log("About to run line #39: ij.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Ne(row.AtIndex('b'))}).Count()")

		runAndAssert(suite.Suite, expected_, ij.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Ne(row.AtIndex("b")) }).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #39")
	}

	// joins.yaml line #46
	// oj = table_test_joins.outer_join(table_test_joins2, lambda x,y:x['a'] == y['b']).zip()
	suite.T().Log("Possibly executing: var oj r.Term = table_test_joins.OuterJoin(table_test_joins2, func(x r.Term, y r.Term) interface{} { return x.AtIndex('a').Eq(y.AtIndex('b'))}).Zip()")

	oj := table_test_joins.OuterJoin(table_test_joins2, func(x r.Term, y r.Term) interface{} { return x.AtIndex("a").Eq(y.AtIndex("b")) }).Zip()
	_ = oj // Prevent any noused variable errors

	{
		// joins.yaml line #49
		/* 2500 */
		var expected_ int = 2500
		/* oj.count() */

		suite.T().Log("About to run line #49: oj.Count()")

		runAndAssert(suite.Suite, expected_, oj.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #49")
	}

	{
		// joins.yaml line #51
		/* 0 */
		var expected_ int = 0
		/* oj.filter(lambda row:row['a'] != row['b']).count() */

		suite.T().Log("About to run line #51: oj.Filter(func(row r.Term) interface{} { return row.AtIndex('a').Ne(row.AtIndex('b'))}).Count()")

		runAndAssert(suite.Suite, expected_, oj.Filter(func(row r.Term) interface{} { return row.AtIndex("a").Ne(row.AtIndex("b")) }).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #51")
	}

	// joins.yaml line #57
	// blah = otable_test_joins.order_by("id").eq_join(r.row['id'], otable_test_joins2, ordered=True).zip()
	suite.T().Log("Possibly executing: var blah r.Term = otable_test_joins.OrderBy('id').EqJoin(r.Row.AtIndex('id'), otable_test_joins2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip()")

	blah := maybeRun(otable_test_joins.OrderBy("id").EqJoin(r.Row.AtIndex("id"), otable_test_joins2).OptArgs(r.EqJoinOpts{Ordered: true}).Zip(), suite.session, r.RunOpts{})
	_ = blah // Prevent any noused variable errors

	// joins.yaml line #59
	// blah = otable_test_joins.order_by(r.desc("id")).eq_join(r.row['id'], otable_test_joins2, ordered=True).zip()
	suite.T().Log("Possibly executing: var blah r.Term = otable_test_joins.OrderBy(r.Desc('id')).EqJoin(r.Row.AtIndex('id'), otable_test_joins2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip()")

	blah = maybeRun(otable_test_joins.OrderBy(r.Desc("id")).EqJoin(r.Row.AtIndex("id"), otable_test_joins2).OptArgs(r.EqJoinOpts{Ordered: true}).Zip(), suite.session, r.RunOpts{})

	// joins.yaml line #61
	// blah = otable_test_joins.order_by("id").eq_join(r.row['a'], otable_test_joins2, ordered=True).zip()
	suite.T().Log("Possibly executing: var blah r.Term = otable_test_joins.OrderBy('id').EqJoin(r.Row.AtIndex('a'), otable_test_joins2).OptArgs(r.EqJoinOpts{Ordered: true, }).Zip()")

	blah = maybeRun(otable_test_joins.OrderBy("id").EqJoin(r.Row.AtIndex("a"), otable_test_joins2).OptArgs(r.EqJoinOpts{Ordered: true}).Zip(), suite.session, r.RunOpts{})

	{
		// joins.yaml line #65
		/* 100 */
		var expected_ int = 100
		/* table_test_joins.eq_join('a', table_test_joins2).zip().count() */

		suite.T().Log("About to run line #65: table_test_joins.EqJoin('a', table_test_joins2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin("a", table_test_joins2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #65")
	}

	{
		// joins.yaml line #68
		/* 0 */
		var expected_ int = 0
		/* table_test_joins.eq_join('fake', table_test_joins2).zip().count() */

		suite.T().Log("About to run line #68: table_test_joins.EqJoin('fake', table_test_joins2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin("fake", table_test_joins2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #68")
	}

	{
		// joins.yaml line #71
		/* 100 */
		var expected_ int = 100
		/* table_test_joins.eq_join(lambda x:x['a'], table_test_joins2).zip().count() */

		suite.T().Log("About to run line #71: table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex('a')}, table_test_joins2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex("a") }, table_test_joins2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #71")
	}

	{
		// joins.yaml line #76
		/* 0 */
		var expected_ int = 0
		/* table_test_joins.eq_join(lambda x:x['fake'], table_test_joins2).zip().count() */

		suite.T().Log("About to run line #76: table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex('fake')}, table_test_joins2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex("fake") }, table_test_joins2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #76")
	}

	{
		// joins.yaml line #81
		/* 0 */
		var expected_ int = 0
		/* table_test_joins.eq_join(lambda x:null, table_test_joins2).zip().count() */

		suite.T().Log("About to run line #81: table_test_joins.EqJoin(func(x r.Term) interface{} { return nil}, table_test_joins2).Zip().Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin(func(x r.Term) interface{} { return nil }, table_test_joins2).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #81")
	}

	{
		// joins.yaml line #86
		/* 100 */
		var expected_ int = 100
		/* table_test_joins.eq_join(lambda x:x['a'], table_test_joins2).count() */

		suite.T().Log("About to run line #86: table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex('a')}, table_test_joins2).Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex("a") }, table_test_joins2).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #86")
	}

	{
		// joins.yaml line #92
		/* 100 */
		var expected_ int = 100
		/* table_test_joins.eq_join('a', table_test_joins3).zip().count() */

		suite.T().Log("About to run line #92: table_test_joins.EqJoin('a', table_test_joins3).Zip().Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin("a", table_test_joins3).Zip().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #92")
	}

	{
		// joins.yaml line #95
		/* 100 */
		var expected_ int = 100
		/* table_test_joins.eq_join(lambda x:x['a'], table_test_joins3).count() */

		suite.T().Log("About to run line #95: table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex('a')}, table_test_joins3).Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin(func(x r.Term) interface{} { return x.AtIndex("a") }, table_test_joins3).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #95")
	}

	{
		// joins.yaml line #101
		/* 100 */
		var expected_ int = 100
		/* table_test_joins.eq_join(r.row['a'], table_test_joins2).count() */

		suite.T().Log("About to run line #101: table_test_joins.EqJoin(r.Row.AtIndex('a'), table_test_joins2).Count()")

		runAndAssert(suite.Suite, expected_, table_test_joins.EqJoin(r.Row.AtIndex("a"), table_test_joins2).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #101")
	}

	// joins.yaml line #106
	// left = r.expr([{'a':1},{'a':2},{'a':3}])
	suite.T().Log("Possibly executing: var left r.Term = r.Expr([]interface{}{map[interface{}]interface{}{'a': 1, }, map[interface{}]interface{}{'a': 2, }, map[interface{}]interface{}{'a': 3, }})")

	left := r.Expr([]interface{}{map[interface{}]interface{}{"a": 1}, map[interface{}]interface{}{"a": 2}, map[interface{}]interface{}{"a": 3}})
	_ = left // Prevent any noused variable errors

	// joins.yaml line #107
	// right = r.expr([{'b':2},{'b':3}])
	suite.T().Log("Possibly executing: var right r.Term = r.Expr([]interface{}{map[interface{}]interface{}{'b': 2, }, map[interface{}]interface{}{'b': 3, }})")

	right := r.Expr([]interface{}{map[interface{}]interface{}{"b": 2}, map[interface{}]interface{}{"b": 3}})
	_ = right // Prevent any noused variable errors

	{
		// joins.yaml line #109
		/* [{'a':2,'b':2},{'a':3,'b':3}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": 2, "b": 2}, map[interface{}]interface{}{"a": 3, "b": 3}}
		/* left.inner_join(right, lambda l, r:l['a'] == r['b']).zip() */

		suite.T().Log("About to run line #109: left.InnerJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex('a').Eq(r.AtIndex('b'))}).Zip()")

		runAndAssert(suite.Suite, expected_, left.InnerJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex("a").Eq(r.AtIndex("b")) }).Zip(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #109")
	}

	{
		// joins.yaml line #115
		/* [{'a':1},{'a':2,'b':2},{'a':3,'b':3}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": 1}, map[interface{}]interface{}{"a": 2, "b": 2}, map[interface{}]interface{}{"a": 3, "b": 3}}
		/* left.outer_join(right, lambda l, r:l['a'] == r['b']).zip() */

		suite.T().Log("About to run line #115: left.OuterJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex('a').Eq(r.AtIndex('b'))}).Zip()")

		runAndAssert(suite.Suite, expected_, left.OuterJoin(right, func(l r.Term, r r.Term) interface{} { return l.AtIndex("a").Eq(r.AtIndex("b")) }).Zip(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #115")
	}

	{
		// joins.yaml line #132
		/* partial({'tables_dropped':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_dropped": 1})
		/* r.db('test').table_drop('test3') */

		suite.T().Log("About to run line #132: r.DB('test').TableDrop('test3')")

		runAndAssert(suite.Suite, expected_, r.DB("db_joins").TableDrop("test3"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #132")
	}
}
