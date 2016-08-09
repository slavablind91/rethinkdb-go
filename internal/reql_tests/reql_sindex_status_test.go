// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../../gen_tests/template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/dancannon/gorethink.v2"
	"gopkg.in/dancannon/gorethink.v2/internal/compare"
)

// sindex status
func TestSindexStatusSuite(t *testing.T) {
	suite.Run(t, new(SindexStatusSuite))
}

type SindexStatusSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *SindexStatusSuite) SetupTest() {
	suite.T().Log("Setting up SindexStatusSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

	r.DB("test").TableDrop("tbl2").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl2").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl2").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *SindexStatusSuite) TearDownSuite() {
	suite.T().Log("Tearing down SindexStatusSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("test").TableDrop("tbl2").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *SindexStatusSuite) TestCases() {
	suite.T().Log("Running SindexStatusSuite: sindex status")

	tbl2 := r.DB("test").Table("tbl2")
	_ = tbl2 // Prevent any noused variable errors

	{
		// sindex/status.yaml line #7
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* tbl2.index_create("a") */

		suite.T().Log("About to run line #7: tbl2.IndexCreate('a')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexCreate("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// sindex/status.yaml line #9
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* tbl2.index_create("b") */

		suite.T().Log("About to run line #9: tbl2.IndexCreate('b')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexCreate("b"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #9")
	}

	{
		// sindex/status.yaml line #12
		/* 2 */
		var expected_ int = 2
		/* tbl2.index_status().count() */

		suite.T().Log("About to run line #12: tbl2.IndexStatus().Count()")

		runAndAssert(suite.Suite, expected_, tbl2.IndexStatus().Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #12")
	}

	{
		// sindex/status.yaml line #14
		/* 1 */
		var expected_ int = 1
		/* tbl2.index_status("a").count() */

		suite.T().Log("About to run line #14: tbl2.IndexStatus('a').Count()")

		runAndAssert(suite.Suite, expected_, tbl2.IndexStatus("a").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #14")
	}

	{
		// sindex/status.yaml line #16
		/* 1 */
		var expected_ int = 1
		/* tbl2.index_status("b").count() */

		suite.T().Log("About to run line #16: tbl2.IndexStatus('b').Count()")

		runAndAssert(suite.Suite, expected_, tbl2.IndexStatus("b").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #16")
	}

	{
		// sindex/status.yaml line #18
		/* 2 */
		var expected_ int = 2
		/* tbl2.index_status("a", "b").count() */

		suite.T().Log("About to run line #18: tbl2.IndexStatus('a', 'b').Count()")

		runAndAssert(suite.Suite, expected_, tbl2.IndexStatus("a", "b").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #18")
	}

	{
		// sindex/status.yaml line #21
		/* ({'dropped':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"dropped": 1}
		/* tbl2.index_drop("a") */

		suite.T().Log("About to run line #21: tbl2.IndexDrop('a')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexDrop("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #21")
	}

	{
		// sindex/status.yaml line #23
		/* ({'dropped':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"dropped": 1}
		/* tbl2.index_drop("b") */

		suite.T().Log("About to run line #23: tbl2.IndexDrop('b')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexDrop("b"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// sindex/status.yaml line #28
		/* partial({'inserted':5000}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"inserted": 5000})
		/* tbl2.insert(r.range(0, 5000).map({'a':r.row})) */

		suite.T().Log("About to run line #28: tbl2.Insert(r.Range(0, 5000).Map(map[interface{}]interface{}{'a': r.Row, }))")

		runAndAssert(suite.Suite, expected_, tbl2.Insert(r.Range(0, 5000).Map(map[interface{}]interface{}{"a": r.Row})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// sindex/status.yaml line #33
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* tbl2.index_create("foo") */

		suite.T().Log("About to run line #33: tbl2.IndexCreate('foo')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexCreate("foo"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #33")
	}

	{
		// sindex/status.yaml line #36
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* tbl2.index_create("bar", multi=True) */

		suite.T().Log("About to run line #36: tbl2.IndexCreate('bar').OptArgs(r.IndexCreateOpts{Multi: true, })")

		runAndAssert(suite.Suite, expected_, tbl2.IndexCreate("bar").OptArgs(r.IndexCreateOpts{Multi: true}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// sindex/status.yaml line #44
		/* [true, true] */
		var expected_ []interface{} = []interface{}{true, true}
		/* tbl2.index_status().map(lambda x:x["progress"] < 1) */

		suite.T().Log("About to run line #44: tbl2.IndexStatus().Map(func(x r.Term) interface{} { return x.AtIndex('progress').Lt(1)})")

		runAndAssert(suite.Suite, expected_, tbl2.IndexStatus().Map(func(x r.Term) interface{} { return x.AtIndex("progress").Lt(1) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #44")
	}

	{
		// sindex/status.yaml line #49
		/* ([true, true]) */
		var expected_ []interface{} = []interface{}{true, true}
		/* tbl2.index_wait()['ready'] */

		suite.T().Log("About to run line #49: tbl2.IndexWait().AtIndex('ready')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexWait().AtIndex("ready"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #49")
	}

	{
		// sindex/status.yaml line #54
		/* bag([false, false]) */
		var expected_ compare.Expected = compare.UnorderedMatch([]interface{}{false, false})
		/* tbl2.index_wait()['geo'] */

		suite.T().Log("About to run line #54: tbl2.IndexWait().AtIndex('geo')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexWait().AtIndex("geo"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #54")
	}

	{
		// sindex/status.yaml line #57
		/* bag([false, true]) */
		var expected_ compare.Expected = compare.UnorderedMatch([]interface{}{false, true})
		/* tbl2.index_wait()['multi'] */

		suite.T().Log("About to run line #57: tbl2.IndexWait().AtIndex('multi')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexWait().AtIndex("multi"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #57")
	}

	{
		// sindex/status.yaml line #60
		/* ([false, false]) */
		var expected_ []interface{} = []interface{}{false, false}
		/* tbl2.index_wait()['outdated'] */

		suite.T().Log("About to run line #60: tbl2.IndexWait().AtIndex('outdated')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexWait().AtIndex("outdated"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #60")
	}

	{
		// sindex/status.yaml line #63
		/* ({'created':1}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"created": 1}
		/* tbl2.index_create("quux") */

		suite.T().Log("About to run line #63: tbl2.IndexCreate('quux')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexCreate("quux"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #63")
	}

	{
		// sindex/status.yaml line #66
		/* true */
		var expected_ bool = true
		/* tbl2.index_status("quux").do(lambda x:(x[0]["index"] == "quux") & (x[0]["progress"] < 1)) */

		suite.T().Log("About to run line #66: tbl2.IndexStatus('quux').Do(func(x r.Term) interface{} { return x.AtIndex(0).AtIndex('index').Eq('quux').And(x.AtIndex(0).AtIndex('progress').Lt(1))})")

		runAndAssert(suite.Suite, expected_, tbl2.IndexStatus("quux").Do(func(x r.Term) interface{} {
			return x.AtIndex(0).AtIndex("index").Eq("quux").And(x.AtIndex(0).AtIndex("progress").Lt(1))
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #66")
	}

	{
		// sindex/status.yaml line #71
		/* ([{'index':'quux', 'ready':true}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"index": "quux", "ready": true}}
		/* tbl2.index_wait("quux").pluck('index', 'ready') */

		suite.T().Log("About to run line #71: tbl2.IndexWait('quux').Pluck('index', 'ready')")

		runAndAssert(suite.Suite, expected_, tbl2.IndexWait("quux").Pluck("index", "ready"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #71")
	}

	{
		// sindex/status.yaml line #74
		/* ("PTYPE<BINARY>") */
		var expected_ string = "PTYPE<BINARY>"
		/* tbl2.index_wait("quux").nth(0).get_field('function').type_of() */

		suite.T().Log("About to run line #74: tbl2.IndexWait('quux').Nth(0).Field('function').TypeOf()")

		runAndAssert(suite.Suite, expected_, tbl2.IndexWait("quux").Nth(0).Field("function").TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #74")
	}
}
