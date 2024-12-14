/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day08

import (
	"reflect"
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseAntennaMap(t *testing.T) {
	type testCase struct {
		text                            string
		expectedBounds                  utilities.Size2D
		expectedAntennas                AntennaLocations
		expectedAntinodes               AntinodeLocations
		expectedUniqueAntinodeLocations int
	}
	testCases := []testCase{
		{
			text: `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
			expectedBounds: utilities.NewSize2D(12, 12),
			expectedAntennas: AntennaLocations{
				'0': []utilities.Point2D{utilities.NewPoint2D(8, 1), utilities.NewPoint2D(5, 2), utilities.NewPoint2D(7, 3), utilities.NewPoint2D(4, 4)},
				'A': []utilities.Point2D{utilities.NewPoint2D(6, 5), utilities.NewPoint2D(8, 8), utilities.NewPoint2D(9, 9)},
			},
			expectedAntinodes: AntinodeLocations{
				utilities.NewPoint2D(6, 0):  true,
				utilities.NewPoint2D(11, 0): true,
				utilities.NewPoint2D(3, 1):  true,
				utilities.NewPoint2D(4, 2):  true,
				utilities.NewPoint2D(10, 2): true,
				utilities.NewPoint2D(2, 3):  true,

				utilities.NewPoint2D(9, 4):   true,
				utilities.NewPoint2D(1, 5):   true,
				utilities.NewPoint2D(6, 5):   true,
				utilities.NewPoint2D(3, 6):   true,
				utilities.NewPoint2D(0, 7):   true,
				utilities.NewPoint2D(7, 7):   true,
				utilities.NewPoint2D(10, 10): true,
				utilities.NewPoint2D(10, 11): true,
			},
			expectedUniqueAntinodeLocations: 14,
		},
	}

	for _, test := range testCases {
		antennaMap := ParseAntennaMap(test.text, false)
		assert.Equal(t, test.expectedBounds, antennaMap.Bounds)
		assert.True(t, reflect.DeepEqual(test.expectedAntennas, antennaMap.Antennas))
		assert.True(t, reflect.DeepEqual(test.expectedAntinodes, antennaMap.Antinodes))
		assert.Equal(t, test.expectedUniqueAntinodeLocations, len(antennaMap.Antinodes))
	}
}

func TestParseAntennaMapHarmonics(t *testing.T) {
	type testCase struct {
		text                            string
		expectedBounds                  utilities.Size2D
		expectedAntennas                AntennaLocations
		expectedAntinodes               AntinodeLocations
		expectedUniqueAntinodeLocations int
	}
	testCases := []testCase{
		{
			text: `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
			expectedBounds: utilities.NewSize2D(12, 12),
			expectedAntennas: AntennaLocations{
				'0': []utilities.Point2D{utilities.NewPoint2D(8, 1), utilities.NewPoint2D(5, 2), utilities.NewPoint2D(7, 3), utilities.NewPoint2D(4, 4)},
				'A': []utilities.Point2D{utilities.NewPoint2D(6, 5), utilities.NewPoint2D(8, 8), utilities.NewPoint2D(9, 9)},
			},
			expectedAntinodes: AntinodeLocations{
				utilities.NewPoint2D(5, 7):   true,
				utilities.NewPoint2D(7, 7):   true,
				utilities.NewPoint2D(10, 10): true,
				utilities.NewPoint2D(4, 4):   true,
				utilities.NewPoint2D(5, 2):   true,
				utilities.NewPoint2D(1, 5):   true,
				utilities.NewPoint2D(3, 11):  true,
				utilities.NewPoint2D(0, 0):   true,
				utilities.NewPoint2D(0, 7):   true,
				utilities.NewPoint2D(11, 5):  true,
				utilities.NewPoint2D(1, 10):  true,
				utilities.NewPoint2D(2, 8):   true,
				utilities.NewPoint2D(4, 9):   true,
				utilities.NewPoint2D(1, 1):   true,
				utilities.NewPoint2D(5, 5):   true,
				utilities.NewPoint2D(4, 2):   true,
				utilities.NewPoint2D(3, 3):   true,
				utilities.NewPoint2D(7, 3):   true,
				utilities.NewPoint2D(3, 6):   true,
				utilities.NewPoint2D(10, 2):  true,
				utilities.NewPoint2D(2, 3):   true,
				utilities.NewPoint2D(6, 6):   true,
				utilities.NewPoint2D(8, 1):   true,
				utilities.NewPoint2D(6, 0):   true,
				utilities.NewPoint2D(11, 0):  true,
				utilities.NewPoint2D(8, 8):   true,
				utilities.NewPoint2D(11, 11): true,
				utilities.NewPoint2D(9, 4):   true,
				utilities.NewPoint2D(6, 5):   true,
				utilities.NewPoint2D(2, 2):   true,
				utilities.NewPoint2D(1, 0):   true,
				utilities.NewPoint2D(3, 1):   true,
				utilities.NewPoint2D(10, 11): true,
				utilities.NewPoint2D(9, 9):   true,
			},
			expectedUniqueAntinodeLocations: 34,
		},
	}

	for _, test := range testCases {
		antennaMap := ParseAntennaMap(test.text, true)
		assert.Equal(t, test.expectedBounds, antennaMap.Bounds)
		assert.True(t, reflect.DeepEqual(test.expectedAntennas, antennaMap.Antennas))
		assert.True(t, reflect.DeepEqual(test.expectedAntinodes, antennaMap.Antinodes))
		assert.Equal(t, test.expectedUniqueAntinodeLocations, len(antennaMap.Antinodes))
	}
}

func TestGenerateCollinearLocations(t *testing.T) {
	type testCase struct {
		pointA            utilities.Point2D
		pointB            utilities.Point2D
		bounds            utilities.Size2D
		expectedLocations []utilities.Point2D
	}
	testCases := []testCase{
		{
			pointA: utilities.NewPoint2D(5, 2),
			pointB: utilities.NewPoint2D(11, 9),
			bounds: utilities.NewSize2D(50, 50),
			expectedLocations: []utilities.Point2D{
				utilities.NewPoint2D(5, 2),
				utilities.NewPoint2D(11, 9),
				utilities.NewPoint2D(17, 16),
				utilities.NewPoint2D(23, 23),
				utilities.NewPoint2D(29, 30),
				utilities.NewPoint2D(35, 37),
				utilities.NewPoint2D(41, 44),
			},
		},
	}

	for _, test := range testCases {
		assert.True(t, reflect.DeepEqual(test.expectedLocations, GenerateCollinearLocations(test.pointA, test.pointB, test.bounds)))
	}
}
