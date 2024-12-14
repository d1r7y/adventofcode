/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyOne_day15

import (
	"sort"
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseSensorLine(t *testing.T) {
	type testCase struct {
		line           string
		expectedSensor *Sensor
	}
	testCases := []testCase{
		{
			line:           "Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
			expectedSensor: NewSensor(Point{X: 2, Y: 18}, Point{-2, 15}),
		},
		{
			line:           "Sensor at x=3008012, y=993590: closest beacon is at x=2971569, y=2563051",
			expectedSensor: NewSensor(Point{X: 3008012, Y: 993590}, Point{X: 2971569, Y: 2563051}),
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedSensor, ParseSensorLine(test.line))
	}
}

func TestParseSensors(t *testing.T) {
	type testCase struct {
		str             string
		expectedSensors SensorList
	}
	testCases := []testCase{
		{
			str: `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`,
			expectedSensors: SensorList{
				NewSensor(Point{X: 2, Y: 18}, Point{-2, 15}),
				NewSensor(Point{X: 9, Y: 16}, Point{10, 16}),
				NewSensor(Point{X: 13, Y: 2}, Point{15, 3}),
				NewSensor(Point{X: 12, Y: 14}, Point{10, 16}),
				NewSensor(Point{X: 10, Y: 20}, Point{10, 16}),
				NewSensor(Point{X: 14, Y: 17}, Point{10, 16}),
				NewSensor(Point{X: 8, Y: 7}, Point{2, 10}),
				NewSensor(Point{X: 2, Y: 0}, Point{2, 10}),
				NewSensor(Point{X: 0, Y: 11}, Point{2, 10}),
				NewSensor(Point{X: 20, Y: 14}, Point{25, 17}),
				NewSensor(Point{X: 17, Y: 20}, Point{21, 22}),
				NewSensor(Point{X: 16, Y: 7}, Point{15, 3}),
				NewSensor(Point{X: 14, Y: 3}, Point{15, 3}),
				NewSensor(Point{X: 20, Y: 1}, Point{15, 3}),
			},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedSensors, ParseSensors(test.str))
	}
}

func TestParseNetwork(t *testing.T) {
	type testCase struct {
		str             string
		expectedNetwork *Network
	}
	testCases := []testCase{
		{
			str: `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`,
			expectedNetwork: &Network{
				Min: Point{-8, -10},
				Max: Point{28, 26},
				Sensors: SensorList{
					NewSensor(Point{X: 2, Y: 18}, Point{-2, 15}),
					NewSensor(Point{X: 9, Y: 16}, Point{10, 16}),
					NewSensor(Point{X: 13, Y: 2}, Point{15, 3}),
					NewSensor(Point{X: 12, Y: 14}, Point{10, 16}),
					NewSensor(Point{X: 10, Y: 20}, Point{10, 16}),
					NewSensor(Point{X: 14, Y: 17}, Point{10, 16}),
					NewSensor(Point{X: 8, Y: 7}, Point{2, 10}),
					NewSensor(Point{X: 2, Y: 0}, Point{2, 10}),
					NewSensor(Point{X: 0, Y: 11}, Point{2, 10}),
					NewSensor(Point{X: 20, Y: 14}, Point{25, 17}),
					NewSensor(Point{X: 17, Y: 20}, Point{21, 22}),
					NewSensor(Point{X: 16, Y: 7}, Point{15, 3}),
					NewSensor(Point{X: 14, Y: 3}, Point{15, 3}),
					NewSensor(Point{X: 20, Y: 1}, Point{15, 3}),
				}},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedNetwork, ParseNetwork(test.str, false))
	}
}

func TestAbsoluteDifference(t *testing.T) {
	type testCase struct {
		a, b               int
		expectedDifference int
	}

	testCases := []testCase{
		{5, 0, 5},
		{0, 5, 5},
		{-5, -10, 5},
		{-10, -5, 5},
		{5, -10, 15},
		{-10, 5, 15},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedDifference, AbsoluteDifference(test.a, test.b))
	}
}

func TestManhattanDistance(t *testing.T) {
	type testCase struct {
		p1, p2           Point
		expectedDistance int
	}

	testCases := []testCase{
		{Point{0, 0}, Point{6, 6}, 12},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedDistance, ManhattanDistance(test.p1, test.p2))
	}
}

func TestClosestSensor(t *testing.T) {
	type testCase struct {
		point          Point
		expectedSensor *Sensor
	}
	testCases := []testCase{
		{
			point:          Point{5, 1},
			expectedSensor: NewSensor(Point{2, 0}, Point{2, 10}),
		},
		{
			point:          Point{0, 11},
			expectedSensor: NewSensor(Point{0, 11}, Point{2, 10}),
		},
	}

	str := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

	n := ParseNetwork(str, false)
	for _, test := range testCases {
		assert.Equal(t, test.expectedSensor, n.ClosestSensor(test.point))
	}
}

func TestSensorIntersect(t *testing.T) {
	type testCase struct {
		point           Point
		expectedSensors SensorList
	}
	testCases := []testCase{
		{
			point: Point{3, 11},
			expectedSensors: SensorList{
				NewSensor(Point{X: 0, Y: 11}, Point{2, 10}),
				NewSensor(Point{X: 8, Y: 7}, Point{2, 10}),
			},
		},
	}

	str := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

	n := ParseNetwork(str, false)
	for _, test := range testCases {
		sensors := n.SensorIntersection(test.point)
		sort.Sort(sensors)
		assert.Equal(t, test.expectedSensors, sensors)
	}
}

func TestInvalidBeaconLocationsCount(t *testing.T) {
	type testCase struct {
		row           int
		expectedCount int
	}
	testCases := []testCase{
		{
			row:           10,
			expectedCount: 26,
		},
	}

	str := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

	n := ParseNetwork(str, false)
	for _, test := range testCases {
		locations := n.InvalidBeaconLocations(test.row)
		assert.Equal(t, test.expectedCount, len(locations))
	}
}

func TestPossibleBeaconLocationCount(t *testing.T) {
	str := `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

	n := ParseNetwork(str, true)
	locations := n.PossibleBeaconLocations()
	assert.Equal(t, 1, len(locations))
	assert.Equal(t, Point{14, 11}, locations[0])
}

func TestParseRiskMap(t *testing.T) {
	str := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	parsedMap := [][]int{
		{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
		{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
		{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
		{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
		{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
		{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
		{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
		{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
		{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
		{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
	}
	rm := ParseRiskMap(str)
	assert.Equal(t, len(parsedMap[0]), rm.Bounds.Width)
	assert.Equal(t, len(parsedMap), rm.Bounds.Height)
	assert.Equal(t, parsedMap, rm.Map)
}

func TestWalkLeastRisk(t *testing.T) {
	str := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	rm := ParseRiskMap(str)

	leastRisk := rm.WalkLeastRiskReversed(utilities.NewPoint2D(rm.Bounds.Width-1, rm.Bounds.Height-1))
	assert.Equal(t, 40, leastRisk)
}
