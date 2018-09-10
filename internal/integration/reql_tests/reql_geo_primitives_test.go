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

// Test geometric primitive constructors
func TestGeoPrimitivesSuite(t *testing.T) {
	suite.Run(t, new(GeoPrimitivesSuite))
}

type GeoPrimitivesSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *GeoPrimitivesSuite) SetupTest() {
	suite.T().Log("Setting up GeoPrimitivesSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_geopr").Exec(suite.session)
	err = r.DBCreate("db_geopr").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_geopr").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *GeoPrimitivesSuite) TearDownSuite() {
	suite.T().Log("Tearing down GeoPrimitivesSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_geopr").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *GeoPrimitivesSuite) TestCases() {
	suite.T().Log("Running GeoPrimitivesSuite: Test geometric primitive constructors")

	{
		// geo/primitives.yaml line #5
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0, -9.04369477050382e-06], [-7.779638566553426e-06, 4.5218473852518965e-06], [7.779638566553426e-06, 4.5218473852518965e-06], [0, -9.04369477050382e-06]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, -9.04369477050382e-06}, []interface{}{-7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{0, -9.04369477050382e-06}}}, "type": "Polygon"}
		/* r.circle([0,0], 1, num_vertices=3) */

		suite.T().Log("About to run line #5: r.Circle([]interface{}{0, 0}, 1).OptArgs(r.CircleOpts{NumVertices: 3, })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, 1).OptArgs(r.CircleOpts{NumVertices: 3}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// geo/primitives.yaml line #10
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0, -9.04369477050382e-06], [-7.779638566553426e-06, 4.5218473852518965e-06], [7.779638566553426e-06, 4.5218473852518965e-06], [0, -9.04369477050382e-06]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, -9.04369477050382e-06}, []interface{}{-7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{0, -9.04369477050382e-06}}}, "type": "Polygon"}
		/* r.circle(r.point(0,0), 1, num_vertices=3) */

		suite.T().Log("About to run line #10: r.Circle(r.Point(0, 0), 1).OptArgs(r.CircleOpts{NumVertices: 3, })")

		runAndAssert(suite.Suite, expected_, r.Circle(r.Point(0, 0), 1).OptArgs(r.CircleOpts{NumVertices: 3}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #10")
	}

	{
		// geo/primitives.yaml line #15
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[0, -9.04369477050382e-06], [-7.779638566553426e-06, 4.5218473852518965e-06], [7.779638566553426e-06, 4.5218473852518965e-06], [0, -9.04369477050382e-06]], 'type':'LineString'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{0, -9.04369477050382e-06}, []interface{}{-7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{0, -9.04369477050382e-06}}, "type": "LineString"}
		/* r.circle([0,0], 1, num_vertices=3, fill=false) */

		suite.T().Log("About to run line #15: r.Circle([]interface{}{0, 0}, 1).OptArgs(r.CircleOpts{NumVertices: 3, Fill: false, })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, 1).OptArgs(r.CircleOpts{NumVertices: 3, Fill: false}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// geo/primitives.yaml line #20
		/* err('ReqlQueryLogicError', 'Radius must be smaller than a quarter of the circumference along the minor axis of the reference ellipsoid.  Got 14000000m, but must be smaller than 9985163.1855612862855m.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Radius must be smaller than a quarter of the circumference along the minor axis of the reference ellipsoid.  Got 14000000m, but must be smaller than 9985163.1855612862855m.")
		/* r.circle([0,0], 14000000, num_vertices=3) */

		suite.T().Log("About to run line #20: r.Circle([]interface{}{0, 0}, 14000000).OptArgs(r.CircleOpts{NumVertices: 3, })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, 14000000).OptArgs(r.CircleOpts{NumVertices: 3}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #20")
	}

	{
		// geo/primitives.yaml line #25
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0, -9.04369477050382e-06], [-7.779638566553426e-06, 4.5218473852518965e-06], [7.779638566553426e-06, 4.5218473852518965e-06], [0, -9.04369477050382e-06]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, -9.04369477050382e-06}, []interface{}{-7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{0, -9.04369477050382e-06}}}, "type": "Polygon"}
		/* r.circle([0,0], 1, num_vertices=3, geo_system='WGS84') */

		suite.T().Log("About to run line #25: r.Circle([]interface{}{0, 0}, 1).OptArgs(r.CircleOpts{NumVertices: 3, GeoSystem: 'WGS84', })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, 1).OptArgs(r.CircleOpts{NumVertices: 3, GeoSystem: "WGS84"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #25")
	}

	{
		// geo/primitives.yaml line #30
		/* err('ReqlQueryLogicError', 'Radius must be smaller than a quarter of the circumference along the minor axis of the reference ellipsoid.  Got 2m, but must be smaller than 1.570796326794896558m.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Radius must be smaller than a quarter of the circumference along the minor axis of the reference ellipsoid.  Got 2m, but must be smaller than 1.570796326794896558m.")
		/* r.circle([0,0], 2, num_vertices=3, geo_system='unit_sphere') */

		suite.T().Log("About to run line #30: r.Circle([]interface{}{0, 0}, 2).OptArgs(r.CircleOpts{NumVertices: 3, GeoSystem: 'unit_sphere', })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, 2).OptArgs(r.CircleOpts{NumVertices: 3, GeoSystem: "unit_sphere"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// geo/primitives.yaml line #35
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0, -5.729577951308232], [-4.966092947444857, 2.861205754495701], [4.966092947444857, 2.861205754495701], [0, -5.729577951308232]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, -5.729577951308232}, []interface{}{-4.966092947444857, 2.861205754495701}, []interface{}{4.966092947444857, 2.861205754495701}, []interface{}{0, -5.729577951308232}}}, "type": "Polygon"}
		/* r.circle([0,0], 0.1, num_vertices=3, geo_system='unit_sphere') */

		suite.T().Log("About to run line #35: r.Circle([]interface{}{0, 0}, 0.1).OptArgs(r.CircleOpts{NumVertices: 3, GeoSystem: 'unit_sphere', })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, 0.1).OptArgs(r.CircleOpts{NumVertices: 3, GeoSystem: "unit_sphere"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #35")
	}

	{
		// geo/primitives.yaml line #42
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0, -9.04369477050382e-06], [-7.779638566553426e-06, 4.5218473852518965e-06], [7.779638566553426e-06, 4.5218473852518965e-06], [0, -9.04369477050382e-06]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, -9.04369477050382e-06}, []interface{}{-7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{0, -9.04369477050382e-06}}}, "type": "Polygon"}
		/* r.circle([0,0], 1.0/1000.0, num_vertices=3, unit='km') */

		suite.T().Log("About to run line #42: r.Circle([]interface{}{0, 0}, r.Div(1.0, 1000.0)).OptArgs(r.CircleOpts{NumVertices: 3, Unit: 'km', })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, r.Div(1.0, 1000.0)).OptArgs(r.CircleOpts{NumVertices: 3, Unit: "km"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #42")
	}

	{
		// geo/primitives.yaml line #47
		/* ({'$reql_type$':'GEOMETRY', 'coordinates':[[[0, -9.04369477050382e-06], [-7.779638566553426e-06, 4.5218473852518965e-06], [7.779638566553426e-06, 4.5218473852518965e-06], [0, -9.04369477050382e-06]]], 'type':'Polygon'}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"$reql_type$": "GEOMETRY", "coordinates": []interface{}{[]interface{}{[]interface{}{0, -9.04369477050382e-06}, []interface{}{-7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{7.779638566553426e-06, 4.5218473852518965e-06}, []interface{}{0, -9.04369477050382e-06}}}, "type": "Polygon"}
		/* r.circle([0,0], 1.0/1609.344, num_vertices=3, unit='mi') */

		suite.T().Log("About to run line #47: r.Circle([]interface{}{0, 0}, r.Div(1.0, 1609.344)).OptArgs(r.CircleOpts{NumVertices: 3, Unit: 'mi', })")

		runAndAssert(suite.Suite, expected_, r.Circle([]interface{}{0, 0}, r.Div(1.0, 1609.344)).OptArgs(r.CircleOpts{NumVertices: 3, Unit: "mi"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #47")
	}
}
