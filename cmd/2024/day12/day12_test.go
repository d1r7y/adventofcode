/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day12

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseMap(t *testing.T) {
	type testCase struct {
		text        string
		expectedMap *Map
	}
	testCases := []testCase{
		{
			text: `AAAA
BBCD
BBCC
EEEC
AAAA`,
			expectedMap: &Map{
				Bounds: utilities.NewSize2D(4, 5),
				Plants: map[PlantType][]utilities.Point2D{
					'A': {utilities.NewPoint2D(0, 0), utilities.NewPoint2D(1, 0), utilities.NewPoint2D(2, 0), utilities.NewPoint2D(3, 0),
						utilities.NewPoint2D(0, 4), utilities.NewPoint2D(1, 4), utilities.NewPoint2D(2, 4), utilities.NewPoint2D(3, 4)},
					'B': {utilities.NewPoint2D(0, 1), utilities.NewPoint2D(1, 1), utilities.NewPoint2D(0, 2), utilities.NewPoint2D(1, 2)},
					'C': {utilities.NewPoint2D(2, 1), utilities.NewPoint2D(2, 2), utilities.NewPoint2D(3, 2), utilities.NewPoint2D(3, 3)},
					'D': {utilities.NewPoint2D(3, 1)},
					'E': {utilities.NewPoint2D(0, 3), utilities.NewPoint2D(1, 3), utilities.NewPoint2D(2, 3)},
				},
				Regions: map[int]*Region{
					0: {Plant: 'A', Plots: &utilities.SetPoint2D{Points: map[utilities.Point2D]bool{utilities.NewPoint2D(0, 0): true, utilities.NewPoint2D(1, 0): true, utilities.NewPoint2D(2, 0): true, utilities.NewPoint2D(3, 0): true}}},
					1: {Plant: 'B', Plots: &utilities.SetPoint2D{Points: map[utilities.Point2D]bool{utilities.NewPoint2D(0, 1): true, utilities.NewPoint2D(0, 2): true, utilities.NewPoint2D(1, 2): true, utilities.NewPoint2D(1, 1): true}}},
					2: {Plant: 'C', Plots: &utilities.SetPoint2D{Points: map[utilities.Point2D]bool{utilities.NewPoint2D(2, 1): true, utilities.NewPoint2D(2, 2): true, utilities.NewPoint2D(3, 2): true, utilities.NewPoint2D(3, 3): true}}},
					3: {Plant: 'D', Plots: &utilities.SetPoint2D{Points: map[utilities.Point2D]bool{utilities.NewPoint2D(3, 1): true}}},
					4: {Plant: 'E', Plots: &utilities.SetPoint2D{Points: map[utilities.Point2D]bool{utilities.NewPoint2D(0, 3): true, utilities.NewPoint2D(1, 3): true, utilities.NewPoint2D(2, 3): true}}},
					5: {Plant: 'A', Plots: &utilities.SetPoint2D{Points: map[utilities.Point2D]bool{utilities.NewPoint2D(0, 4): true, utilities.NewPoint2D(1, 4): true, utilities.NewPoint2D(2, 4): true, utilities.NewPoint2D(3, 4): true}}},
				},

				Columns: []Row{
					{'A', 'A', 'A', 'A'},
					{'B', 'B', 'C', 'D'},
					{'B', 'B', 'C', 'C'},
					{'E', 'E', 'E', 'C'},
					{'A', 'A', 'A', 'A'},
				},
			},
		},
	}

	for _, test := range testCases {
		gardenMap := ParseMap(test.text)
		assert.True(t, reflect.DeepEqual(test.expectedMap, gardenMap))
	}
}

func TestRegionArea(t *testing.T) {
	type testCase struct {
		text         string
		regionID     int
		expectedArea int
	}
	testCases := []testCase{
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:     0,
			expectedArea: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:     1,
			expectedArea: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:     2,
			expectedArea: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:     3,
			expectedArea: 1,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:     4,
			expectedArea: 3,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:     5,
			expectedArea: 0,
		},
	}

	for _, test := range testCases {
		gardenMap := ParseMap(test.text)
		assert.Equal(t, test.expectedArea, gardenMap.RegionArea(test.regionID))
	}
}

func TestRegionPerimeter(t *testing.T) {
	type testCase struct {
		text              string
		regionID          int
		expectedPerimeter int
	}
	testCases := []testCase{
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          0,
			expectedPerimeter: 10,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          1,
			expectedPerimeter: 8,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          2,
			expectedPerimeter: 10,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          3,
			expectedPerimeter: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          4,
			expectedPerimeter: 8,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          5,
			expectedPerimeter: 0,
		},
	}

	for _, test := range testCases {
		gardenMap := ParseMap(test.text)
		assert.Equal(t, test.expectedPerimeter, gardenMap.RegionPerimeter(test.regionID))
	}
}

func TestRegionPerimeter2(t *testing.T) {
	type testCase struct {
		text              string
		regionID          int
		expectedPerimeter int
	}
	testCases := []testCase{
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          0,
			expectedPerimeter: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          1,
			expectedPerimeter: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          2,
			expectedPerimeter: 8,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          3,
			expectedPerimeter: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          4,
			expectedPerimeter: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:          5,
			expectedPerimeter: 0,
		},
		{
			text: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			regionID:          0,
			expectedPerimeter: 12,
		},
		{
			text: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			regionID:          1,
			expectedPerimeter: 4,
		},
		{
			text: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			regionID:          2,
			expectedPerimeter: 4,
		},
	}

	for _, test := range testCases {
		gardenMap := ParseMap(test.text)
		assert.Equal(t, test.expectedPerimeter, gardenMap.RegionPerimeterSides(test.regionID))
	}
}

func TestRegionFencingPrice(t *testing.T) {
	type testCase struct {
		text                 string
		regionID             int
		expectedFencingPrice int
	}
	testCases := []testCase{
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             0,
			expectedFencingPrice: 40,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             1,
			expectedFencingPrice: 32,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             2,
			expectedFencingPrice: 40,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             3,
			expectedFencingPrice: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             4,
			expectedFencingPrice: 24,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             5,
			expectedFencingPrice: 0,
		},
	}

	for _, test := range testCases {
		gardenMap := ParseMap(test.text)
		assert.Equal(t, test.expectedFencingPrice, gardenMap.RegionFencingPrice(test.regionID))
	}
}

func TestRegionFencingPriceBulkDiscount(t *testing.T) {
	type testCase struct {
		text                 string
		regionID             int
		expectedFencingPrice int
	}
	testCases := []testCase{
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             0,
			expectedFencingPrice: 16,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             1,
			expectedFencingPrice: 16,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             2,
			expectedFencingPrice: 32,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             3,
			expectedFencingPrice: 4,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             4,
			expectedFencingPrice: 12,
		},
		{
			text: `AAAA
BBCD
BBCC
EEEC`,
			regionID:             5,
			expectedFencingPrice: 0,
		},
		{
			text: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			regionID:             0,
			expectedFencingPrice: 204,
		},
		{
			text: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			regionID:             1,
			expectedFencingPrice: 16,
		},
		{
			text: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			regionID:             2,
			expectedFencingPrice: 16,
		},
	}

	for _, test := range testCases {
		gardenMap := ParseMap(test.text)
		assert.Equal(t, test.expectedFencingPrice, gardenMap.RegionFencingPriceBulkDiscount(test.regionID))
	}
}

func TestTotalMapFencingPrice(t *testing.T) {
	gardenMap := ParseMap(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)

	totalFencingPrice := 0

	for i := 0; i < gardenMap.NumRegions(); i++ {
		totalFencingPrice += gardenMap.RegionFencingPrice(i)
	}

	assert.Equal(t, 1930, totalFencingPrice)
}

func TestTotalMapFencingPriceBulkDiscount(t *testing.T) {
	type testCase struct {
		text                      string
		expectedTotalFencingPrice int
	}

	testCases := []testCase{
		// 		{
		// 			text: `RRRRIICCFF
		// RRRRIICCCF
		// VVRRRCCFFF
		// VVRCCCJFFF
		// VVVVCJJCFE
		// VVIVCCJJEE
		// VVIIICJJEE
		// MIIIIIJJEE
		// MIIISIJEEE
		// MMMISSJEEE`,
		// 			expectedTotalFencingPrice: 1206,
		// 		},
		{
			text: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			expectedTotalFencingPrice: 368,
		},
	}

	for _, test := range testCases {
		gardenMap := ParseMap(test.text)
		totalFencingPrice := 0

		for i := 0; i < gardenMap.NumRegions(); i++ {
			fencePrice := gardenMap.RegionFencingPriceBulkDiscount(i)
			fmt.Printf("Region: %s price %d\n", string(gardenMap.Regions[i].Plant), fencePrice)
			totalFencingPrice += fencePrice
		}

		assert.Equal(t, test.expectedTotalFencingPrice, totalFencingPrice)
	}
}
