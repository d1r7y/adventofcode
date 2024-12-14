/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day04

import (
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseLetterGrid(t *testing.T) {
	type testCase struct {
		text           string
		expectedBounds utilities.Size2D
		expectedRows   [][]rune
	}
	testCases := []testCase{
		{
			text: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			expectedBounds: utilities.Size2D{Width: 10, Height: 10},
			expectedRows: [][]rune{
				{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
				{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
				{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
				{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
				{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
				{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
				{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
				{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
				{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
				{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
			},
		},
	}

	for _, test := range testCases {
		lg := ParseLetterGrid(test.text)
		assert.Equal(t, test.expectedBounds, lg.Bounds)
		assert.Equal(t, test.expectedRows, lg.Rows)
	}
}

func TestGetLetter(t *testing.T) {
	type testCase struct {
		text           string
		position       utilities.Point2D
		expectedLetter rune
	}
	testCases := []testCase{
		{
			text: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			position:       utilities.Point2D{X: 0, Y: 0},
			expectedLetter: 'M',
		},
		{
			text: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			position:       utilities.Point2D{X: 9, Y: 9},
			expectedLetter: 'X',
		},
	}

	for _, test := range testCases {
		lg := ParseLetterGrid(test.text)
		assert.Equal(t, test.expectedLetter, lg.GetLetter(test.position))
	}
}

func TestFindString(t *testing.T) {
	type testCase struct {
		text              string
		str               string
		expectedPositions []utilities.Point2D
	}
	testCases := []testCase{
		{
			text: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			str: "XMAS",
			expectedPositions: []utilities.Point2D{
				utilities.NewPoint2D(5, 0),
				utilities.NewPoint2D(4, 1),
				utilities.NewPoint2D(9, 3),
				utilities.NewPoint2D(9, 3),
				utilities.NewPoint2D(0, 4),
				utilities.NewPoint2D(6, 4),
				utilities.NewPoint2D(6, 4),
				utilities.NewPoint2D(5, 9),
				utilities.NewPoint2D(5, 9),
				utilities.NewPoint2D(5, 9),
				utilities.NewPoint2D(9, 9),
				utilities.NewPoint2D(9, 9),
				utilities.NewPoint2D(0, 5),
				utilities.NewPoint2D(1, 9),
				utilities.NewPoint2D(3, 9),
				utilities.NewPoint2D(3, 9),
				utilities.NewPoint2D(6, 5),
				utilities.NewPoint2D(4, 0),
			},
		},
	}

	for _, test := range testCases {
		lg := ParseLetterGrid(test.text)
		utilities.SortPoints(test.expectedPositions)
		assert.Equal(t, test.expectedPositions, lg.FindString(test.str))
	}
}

func TestFindXMAS(t *testing.T) {
	type testCase struct {
		text              string
		expectedPositions []utilities.Point2D
	}
	testCases := []testCase{
		{
			text: `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			expectedPositions: []utilities.Point2D{
				utilities.NewPoint2D(2, 1),
				utilities.NewPoint2D(6, 2),
				utilities.NewPoint2D(7, 2),
				utilities.NewPoint2D(2, 3),
				utilities.NewPoint2D(4, 3),
				utilities.NewPoint2D(1, 7),
				utilities.NewPoint2D(3, 7),
				utilities.NewPoint2D(5, 7),
				utilities.NewPoint2D(7, 7),
			},
		},
	}

	for _, test := range testCases {
		lg := ParseLetterGrid(test.text)
		utilities.SortPoints(test.expectedPositions)
		assert.Equal(t, test.expectedPositions, lg.FindXMAS())
	}
}
