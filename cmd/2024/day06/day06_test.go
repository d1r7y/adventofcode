/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day06

import (
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseMap(t *testing.T) {
	type testCase struct {
		text             string
		expectedBounds   utilities.Size2D
		expectedPosition utilities.Point2D
		expectedFacing   Direction
		expectedColumns  []Row
	}
	testCases := []testCase{
		{
			text: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			expectedBounds:   utilities.NewSize2D(10, 10),
			expectedPosition: utilities.NewPoint2D(4, 6),
			expectedFacing:   North,
			expectedColumns: []Row{
				{Empty, Empty, Empty, Empty, Obstruction, Empty, Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Obstruction},
				{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				{Empty, Empty, Obstruction, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Obstruction, Empty, Empty},
				{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				{Empty, Obstruction, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Obstruction, Empty},
				{Obstruction, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty, Empty, Empty, Obstruction, Empty, Empty, Empty},
			},
		},
	}

	for _, test := range testCases {
		roomMap := ParseMap(test.text)
		assert.Equal(t, test.expectedBounds, roomMap.Bounds)
		assert.Equal(t, test.expectedPosition, roomMap.Position)
		assert.Equal(t, test.expectedFacing, roomMap.Facing)
		assert.Equal(t, test.expectedColumns, roomMap.Columns)
	}
}

func TestWalk(t *testing.T) {
	type testCase struct {
		position         utilities.Point2D
		facing           Direction
		expectedDone     bool
		expectedPosition utilities.Point2D
		expectedFacing   Direction
	}

	text := `....#.....
.........#
..........
..#.......
.......#..
..........
.#........
........#.
#.........
......#...`
	testCases := []testCase{
		{
			position:         utilities.NewPoint2D(0, 0),
			facing:           North,
			expectedDone:     true,
			expectedPosition: utilities.NewPoint2D(0, 0),
			expectedFacing:   North,
		},
		{
			position:         utilities.NewPoint2D(0, 0),
			facing:           West,
			expectedDone:     true,
			expectedPosition: utilities.NewPoint2D(0, 0),
			expectedFacing:   North,
		},
		{
			position:         utilities.NewPoint2D(0, 0),
			facing:           East,
			expectedDone:     false,
			expectedPosition: utilities.NewPoint2D(1, 0),
			expectedFacing:   East,
		},
		{
			position:         utilities.NewPoint2D(0, 0),
			facing:           South,
			expectedDone:     false,
			expectedPosition: utilities.NewPoint2D(0, 1),
			expectedFacing:   South,
		},
		{
			position:         utilities.NewPoint2D(2, 4),
			facing:           North,
			expectedDone:     false,
			expectedPosition: utilities.NewPoint2D(2, 4),
			expectedFacing:   East,
		},
		{
			position:         utilities.NewPoint2D(1, 3),
			facing:           East,
			expectedDone:     false,
			expectedPosition: utilities.NewPoint2D(1, 3),
			expectedFacing:   South,
		},
	}

	for _, test := range testCases {
		roomMap := ParseMap(text)
		roomMap.Position = test.position
		roomMap.Facing = test.facing

		done := roomMap.Walk()

		assert.Equal(t, test.expectedDone, done, test)
		if !done {
			assert.Equal(t, test.expectedPosition, roomMap.Position, test)
			assert.Equal(t, test.expectedFacing, roomMap.Facing, test)
		}
	}
}

func TestTotalVisited(t *testing.T) {
	type testCase struct {
		text                 string
		expectedVisitedCells int
	}
	testCases := []testCase{
		{
			text: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			expectedVisitedCells: 41,
		},
	}

	for _, test := range testCases {
		roomMap := ParseMap(test.text)
		for {
			if roomMap.Walk() {
				break
			}
		}
		assert.Equal(t, test.expectedVisitedCells, roomMap.VisitedCells)
	}
}

func TestLoopingObstructionCount(t *testing.T) {
	type testCase struct {
		text                            string
		expectedLoopingObstructionCount int
	}
	testCases := []testCase{
		{
			text: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			expectedLoopingObstructionCount: 6,
		},
	}

	for _, test := range testCases {
		roomMap := ParseMap(test.text)

		loopingObstructionCount := 0

		for y := 0; y < roomMap.Bounds.Height; y++ {
			for x := 0; x < roomMap.Bounds.Width; x++ {
				obstructionLocation := utilities.NewPoint2D(x, y)
				if roomMap.Position == obstructionLocation {
					continue
				}

				obstructedRoomMap := ParseMap(test.text)
				obstructedRoomMap.AddObstruction(obstructionLocation)

				for {
					if obstructedRoomMap.Walk() {
						break
					}

					if obstructedRoomMap.AreLooping() {
						loopingObstructionCount++
						break
					}
				}
			}
		}

		assert.Equal(t, test.expectedLoopingObstructionCount, loopingObstructionCount)
	}
}
