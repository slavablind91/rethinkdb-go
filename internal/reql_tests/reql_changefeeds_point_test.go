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

// Test point changebasics
func TestChangefeedsPointSuite(t *testing.T) {
	suite.Run(t, new(ChangefeedsPointSuite))
}

type ChangefeedsPointSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *ChangefeedsPointSuite) SetupTest() {
	suite.T().Log("Setting up ChangefeedsPointSuite")
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

	r.DB("test").TableDrop("tbl").Exec(suite.session)
	err = r.DB("test").TableCreate("tbl").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Table("tbl").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *ChangefeedsPointSuite) TearDownSuite() {
	suite.T().Log("Tearing down ChangefeedsPointSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("test").TableDrop("tbl").Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *ChangefeedsPointSuite) TestCases() {
	suite.T().Log("Running ChangefeedsPointSuite: Test point changebasics")

	tbl := r.DB("test").Table("tbl")
	_ = tbl // Prevent any noused variable errors

	// changefeeds/point.yaml line #10
	// basic = tbl.get(1).changes(include_initial=True)
	suite.T().Log("Possibly executing: var basic r.Term = tbl.Get(1).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true, })")

	basic := maybeRun(tbl.Get(1).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true}), suite.session, r.RunOpts{})
	_ = basic // Prevent any noused variable errors

	{
		// changefeeds/point.yaml line #14
		/* [{'new_val':null}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"new_val": nil}}
		/* fetch(basic, 1) */

		suite.T().Log("About to run line #14: fetch(basic, 1)")

		fetchAndAssert(suite.Suite, expected_, basic, 1)
		suite.T().Log("Finished running line #14")
	}

	{
		// changefeeds/point.yaml line #19
		/* partial({'errors':0, 'inserted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 1})
		/* tbl.insert({'id':1}) */

		suite.T().Log("About to run line #19: tbl.Insert(map[interface{}]interface{}{'id': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Insert(map[interface{}]interface{}{"id": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// changefeeds/point.yaml line #22
		/* [{'old_val':null, 'new_val':{'id':1}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 1}}}
		/* fetch(basic, 1) */

		suite.T().Log("About to run line #22: fetch(basic, 1)")

		fetchAndAssert(suite.Suite, expected_, basic, 1)
		suite.T().Log("Finished running line #22")
	}

	{
		// changefeeds/point.yaml line #27
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* tbl.get(1).update({'update':1}) */

		suite.T().Log("About to run line #27: tbl.Get(1).Update(map[interface{}]interface{}{'update': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(1).Update(map[interface{}]interface{}{"update": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// changefeeds/point.yaml line #30
		/* [{'old_val':{'id':1}, 'new_val':{'id':1,'update':1}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1}, "new_val": map[interface{}]interface{}{"id": 1, "update": 1}}}
		/* fetch(basic, 1) */

		suite.T().Log("About to run line #30: fetch(basic, 1)")

		fetchAndAssert(suite.Suite, expected_, basic, 1)
		suite.T().Log("Finished running line #30")
	}

	{
		// changefeeds/point.yaml line #35
		/* partial({'errors':0, 'deleted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "deleted": 1})
		/* tbl.get(1).delete() */

		suite.T().Log("About to run line #35: tbl.Get(1).Delete()")

		runAndAssert(suite.Suite, expected_, tbl.Get(1).Delete(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #35")
	}

	{
		// changefeeds/point.yaml line #38
		/* [{'old_val':{'id':1,'update':1}, 'new_val':null}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1, "update": 1}, "new_val": nil}}
		/* fetch(basic, 1) */

		suite.T().Log("About to run line #38: fetch(basic, 1)")

		fetchAndAssert(suite.Suite, expected_, basic, 1)
		suite.T().Log("Finished running line #38")
	}

	// changefeeds/point.yaml line #49
	// filter = tbl.get(1).changes(squash=false,include_initial=True).filter(r.row['new_val']['update'].gt(2))['new_val']['update']
	suite.T().Log("Possibly executing: var filter r.Term = tbl.Get(1).Changes().OptArgs(r.ChangesOpts{Squash: false, IncludeInitial: true, }).Filter(r.Row.AtIndex('new_val').AtIndex('update').Gt(2)).AtIndex('new_val').AtIndex('update')")

	filter := maybeRun(tbl.Get(1).Changes().OptArgs(r.ChangesOpts{Squash: false, IncludeInitial: true}).Filter(r.Row.AtIndex("new_val").AtIndex("update").Gt(2)).AtIndex("new_val").AtIndex("update"), suite.session, r.RunOpts{})
	_ = filter // Prevent any noused variable errors

	{
		// changefeeds/point.yaml line #53
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* tbl.insert({'id':1, 'update':1}) */

		suite.T().Log("About to run line #53: tbl.Insert(map[interface{}]interface{}{'id': 1, 'update': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Insert(map[interface{}]interface{}{"id": 1, "update": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// changefeeds/point.yaml line #54
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* tbl.get(1).update({'update':4}) */

		suite.T().Log("About to run line #54: tbl.Get(1).Update(map[interface{}]interface{}{'update': 4, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(1).Update(map[interface{}]interface{}{"update": 4}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #54")
	}

	{
		// changefeeds/point.yaml line #55
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* tbl.get(1).update({'update':1}) */

		suite.T().Log("About to run line #55: tbl.Get(1).Update(map[interface{}]interface{}{'update': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(1).Update(map[interface{}]interface{}{"update": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #55")
	}

	{
		// changefeeds/point.yaml line #56
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* tbl.get(1).update({'update':7}) */

		suite.T().Log("About to run line #56: tbl.Get(1).Update(map[interface{}]interface{}{'update': 7, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(1).Update(map[interface{}]interface{}{"update": 7}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #56")
	}

	{
		// changefeeds/point.yaml line #58
		/* [4,7] */
		var expected_ []interface{} = []interface{}{4, 7}
		/* fetch(filter, 2) */

		suite.T().Log("About to run line #58: fetch(filter, 2)")

		fetchAndAssert(suite.Suite, expected_, filter, 2)
		suite.T().Log("Finished running line #58")
	}

	// changefeeds/point.yaml line #63
	// pluck = tbl.get(3).changes(squash=false,include_initial=True).pluck({'new_val':['red', 'blue']})['new_val']
	suite.T().Log("Possibly executing: var pluck r.Term = tbl.Get(3).Changes().OptArgs(r.ChangesOpts{Squash: false, IncludeInitial: true, }).Pluck(map[interface{}]interface{}{'new_val': []interface{}{'red', 'blue'}, }).AtIndex('new_val')")

	pluck := maybeRun(tbl.Get(3).Changes().OptArgs(r.ChangesOpts{Squash: false, IncludeInitial: true}).Pluck(map[interface{}]interface{}{"new_val": []interface{}{"red", "blue"}}).AtIndex("new_val"), suite.session, r.RunOpts{})
	_ = pluck // Prevent any noused variable errors

	{
		// changefeeds/point.yaml line #67
		/* partial({'errors':0, 'inserted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 1})
		/* tbl.insert({'id':3, 'red':1, 'green':1}) */

		suite.T().Log("About to run line #67: tbl.Insert(map[interface{}]interface{}{'id': 3, 'red': 1, 'green': 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Insert(map[interface{}]interface{}{"id": 3, "red": 1, "green": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #67")
	}

	{
		// changefeeds/point.yaml line #69
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* tbl.get(3).update({'blue':2, 'green':3}) */

		suite.T().Log("About to run line #69: tbl.Get(3).Update(map[interface{}]interface{}{'blue': 2, 'green': 3, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(3).Update(map[interface{}]interface{}{"blue": 2, "green": 3}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #69")
	}

	{
		// changefeeds/point.yaml line #71
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* tbl.get(3).update({'green':4}) */

		suite.T().Log("About to run line #71: tbl.Get(3).Update(map[interface{}]interface{}{'green': 4, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(3).Update(map[interface{}]interface{}{"green": 4}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #71")
	}

	{
		// changefeeds/point.yaml line #73
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* tbl.get(3).update({'blue':4}) */

		suite.T().Log("About to run line #73: tbl.Get(3).Update(map[interface{}]interface{}{'blue': 4, })")

		runAndAssert(suite.Suite, expected_, tbl.Get(3).Update(map[interface{}]interface{}{"blue": 4}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #73")
	}

	{
		// changefeeds/point.yaml line #76
		/* [{'red': 1}, {'blue': 2, 'red': 1}, {'blue': 2, 'red': 1}, {'blue': 4, 'red': 1}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"red": 1}, map[interface{}]interface{}{"blue": 2, "red": 1}, map[interface{}]interface{}{"blue": 2, "red": 1}, map[interface{}]interface{}{"blue": 4, "red": 1}}
		/* fetch(pluck, 4) */

		suite.T().Log("About to run line #76: fetch(pluck, 4)")

		fetchAndAssert(suite.Suite, expected_, pluck, 4)
		suite.T().Log("Finished running line #76")
	}

	// changefeeds/point.yaml line #83
	// dtbl = r.db('rethinkdb').table('_debug_scratch')
	suite.T().Log("Possibly executing: var dtbl r.Term = r.DB('rethinkdb').Table('_debug_scratch')")

	dtbl := r.DB("rethinkdb").Table("_debug_scratch")
	_ = dtbl // Prevent any noused variable errors

	// changefeeds/point.yaml line #86
	// debug = dtbl.get(1).changes(include_initial=True)
	suite.T().Log("Possibly executing: var debug r.Term = dtbl.Get(1).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true, })")

	debug := maybeRun(dtbl.Get(1).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true}), suite.session, r.RunOpts{})
	_ = debug // Prevent any noused variable errors

	{
		// changefeeds/point.yaml line #88
		/* [{'new_val':null}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"new_val": nil}}
		/* fetch(debug, 1) */

		suite.T().Log("About to run line #88: fetch(debug, 1)")

		fetchAndAssert(suite.Suite, expected_, debug, 1)
		suite.T().Log("Finished running line #88")
	}

	{
		// changefeeds/point.yaml line #91
		/* partial({'errors':0, 'inserted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "inserted": 1})
		/* dtbl.insert({'id':1}) */

		suite.T().Log("About to run line #91: dtbl.Insert(map[interface{}]interface{}{'id': 1, })")

		runAndAssert(suite.Suite, expected_, dtbl.Insert(map[interface{}]interface{}{"id": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #91")
	}

	{
		// changefeeds/point.yaml line #93
		/* [{'old_val':null, 'new_val':{'id':1}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": nil, "new_val": map[interface{}]interface{}{"id": 1}}}
		/* fetch(debug, 1) */

		suite.T().Log("About to run line #93: fetch(debug, 1)")

		fetchAndAssert(suite.Suite, expected_, debug, 1)
		suite.T().Log("Finished running line #93")
	}

	{
		// changefeeds/point.yaml line #96
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* dtbl.get(1).update({'update':1}) */

		suite.T().Log("About to run line #96: dtbl.Get(1).Update(map[interface{}]interface{}{'update': 1, })")

		runAndAssert(suite.Suite, expected_, dtbl.Get(1).Update(map[interface{}]interface{}{"update": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #96")
	}

	{
		// changefeeds/point.yaml line #98
		/* [{'old_val':{'id':1}, 'new_val':{'id':1,'update':1}}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1}, "new_val": map[interface{}]interface{}{"id": 1, "update": 1}}}
		/* fetch(debug, 1) */

		suite.T().Log("About to run line #98: fetch(debug, 1)")

		fetchAndAssert(suite.Suite, expected_, debug, 1)
		suite.T().Log("Finished running line #98")
	}

	{
		// changefeeds/point.yaml line #101
		/* partial({'errors':0, 'deleted':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "deleted": 1})
		/* dtbl.get(1).delete() */

		suite.T().Log("About to run line #101: dtbl.Get(1).Delete()")

		runAndAssert(suite.Suite, expected_, dtbl.Get(1).Delete(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #101")
	}

	{
		// changefeeds/point.yaml line #103
		/* [{'old_val':{'id':1,'update':1}, 'new_val':null}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"old_val": map[interface{}]interface{}{"id": 1, "update": 1}, "new_val": nil}}
		/* fetch(debug, 1) */

		suite.T().Log("About to run line #103: fetch(debug, 1)")

		fetchAndAssert(suite.Suite, expected_, debug, 1)
		suite.T().Log("Finished running line #103")
	}

	{
		// changefeeds/point.yaml line #106
		/* {'skipped':0, 'deleted':0, 'unchanged':0, 'errors':0, 'replaced':0, 'inserted':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"skipped": 0, "deleted": 0, "unchanged": 0, "errors": 0, "replaced": 0, "inserted": 1}
		/* dtbl.insert({'id':5, 'red':1, 'green':1}) */

		suite.T().Log("About to run line #106: dtbl.Insert(map[interface{}]interface{}{'id': 5, 'red': 1, 'green': 1, })")

		runAndAssert(suite.Suite, expected_, dtbl.Insert(map[interface{}]interface{}{"id": 5, "red": 1, "green": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #106")
	}

	// changefeeds/point.yaml line #108
	// dtblPluck = dtbl.get(5).changes(include_initial=True).pluck({'new_val':['red', 'blue']})['new_val']
	suite.T().Log("Possibly executing: var dtblPluck r.Term = dtbl.Get(5).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true, }).Pluck(map[interface{}]interface{}{'new_val': []interface{}{'red', 'blue'}, }).AtIndex('new_val')")

	dtblPluck := maybeRun(dtbl.Get(5).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true}).Pluck(map[interface{}]interface{}{"new_val": []interface{}{"red", "blue"}}).AtIndex("new_val"), suite.session, r.RunOpts{})
	_ = dtblPluck // Prevent any noused variable errors

	{
		// changefeeds/point.yaml line #113
		/* [{'red':1}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"red": 1}}
		/* fetch(dtblPluck, 1) */

		suite.T().Log("About to run line #113: fetch(dtblPluck, 1)")

		fetchAndAssert(suite.Suite, expected_, dtblPluck, 1)
		suite.T().Log("Finished running line #113")
	}

	{
		// changefeeds/point.yaml line #116
		/* partial({'errors':0, 'replaced':1}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"errors": 0, "replaced": 1})
		/* dtbl.get(5).update({'blue':2, 'green':3}) */

		suite.T().Log("About to run line #116: dtbl.Get(5).Update(map[interface{}]interface{}{'blue': 2, 'green': 3, })")

		runAndAssert(suite.Suite, expected_, dtbl.Get(5).Update(map[interface{}]interface{}{"blue": 2, "green": 3}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #116")
	}

	{
		// changefeeds/point.yaml line #119
		/* [{'blue':2, 'red':1}] */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"blue": 2, "red": 1}}
		/* fetch(dtblPluck, 1) */

		suite.T().Log("About to run line #119: fetch(dtblPluck, 1)")

		fetchAndAssert(suite.Suite, expected_, dtblPluck, 1)
		suite.T().Log("Finished running line #119")
	}

	// changefeeds/point.yaml line #132
	// tableId = tbl.info()['id']
	suite.T().Log("Possibly executing: var tableId r.Term = tbl.Info().AtIndex('id')")

	tableId := maybeRun(tbl.Info().AtIndex("id"), suite.session, r.RunOpts{})
	_ = tableId // Prevent any noused variable errors

	// changefeeds/point.yaml line #136
	// rtblPluck = r.db('rethinkdb').table('table_status').get(tableId).changes(include_initial=True)
	suite.T().Log("Possibly executing: var rtblPluck r.Term = r.DB('rethinkdb').Table('table_status').Get(tableId).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true, })")

	rtblPluck := maybeRun(r.DB("rethinkdb").Table("table_status").Get(tableId).Changes().OptArgs(r.ChangesOpts{IncludeInitial: true}), suite.session, r.RunOpts{})
	_ = rtblPluck // Prevent any noused variable errors

	{
		// changefeeds/point.yaml line #137
		/* partial([{'new_val':partial({'db':'test'})}]) */
		var expected_ compare.Expected = compare.PartialMatch([]interface{}{map[interface{}]interface{}{"new_val": compare.PartialMatch(map[interface{}]interface{}{"db": "test"})}})
		/* fetch(rtblPluck, 1) */

		suite.T().Log("About to run line #137: fetch(rtblPluck, 1)")

		fetchAndAssert(suite.Suite, expected_, rtblPluck, 1)
		suite.T().Log("Finished running line #137")
	}

	{
		// changefeeds/point.yaml line #140
		/* AnythingIsFine */
		var expected_ string = compare.AnythingIsFine
		/* tbl.reconfigure(shards=3, replicas=1) */

		suite.T().Log("About to run line #140: tbl.Reconfigure().OptArgs(r.ReconfigureOpts{Shards: 3, Replicas: 1, })")

		runAndAssert(suite.Suite, expected_, tbl.Reconfigure().OptArgs(r.ReconfigureOpts{Shards: 3, Replicas: 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #140")
	}

	{
		// changefeeds/point.yaml line #143
		/* partial([{'old_val':partial({'db':'test'}), 'new_val':partial({'db':'test'})}]) */
		var expected_ compare.Expected = compare.PartialMatch([]interface{}{map[interface{}]interface{}{"old_val": compare.PartialMatch(map[interface{}]interface{}{"db": "test"}), "new_val": compare.PartialMatch(map[interface{}]interface{}{"db": "test"})}})
		/* fetch(rtblPluck, 1, 2) */

		suite.T().Log("About to run line #143: fetch(rtblPluck, 1)")

		fetchAndAssert(suite.Suite, expected_, rtblPluck, 1)
		suite.T().Log("Finished running line #143")
	}
}
