/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day11

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStones(t *testing.T) {
	type testCase struct {
		text              string
		expectedStoneList *StoneList
	}
	testCases := []testCase{
		{
			text: "0 1 10 99 999",
			expectedStoneList: &StoneList{
				Stones: []Stone{
					{0},
					{1},
					{10},
					{99},
					{999},
				},
			},
		},
	}

	for _, test := range testCases {
		stoneList := ParseStones(test.text)
		assert.True(t, reflect.DeepEqual(test.expectedStoneList, stoneList))
	}
}

func TestSplitDigits(t *testing.T) {
	type testCase struct {
		number              int
		expectedLeftDigits  int
		expectedRightDigits int
	}
	testCases := []testCase{
		{
			number:              12,
			expectedLeftDigits:  1,
			expectedRightDigits: 2,
		},
		{
			number:              1212,
			expectedLeftDigits:  12,
			expectedRightDigits: 12,
		},
		{
			number:              12345678,
			expectedLeftDigits:  1234,
			expectedRightDigits: 5678,
		},
		{
			number:              1001,
			expectedLeftDigits:  10,
			expectedRightDigits: 1,
		},
	}

	for _, test := range testCases {
		left, right := SplitDigits(test.number)
		assert.Equal(t, test.expectedLeftDigits, left)
		assert.Equal(t, test.expectedRightDigits, right)
	}
}

func TestBlink(t *testing.T) {
	type testCase struct {
		text              string
		numBlinks         int
		expectedStoneList *StoneList
	}
	testCases := []testCase{
		{
			text:      "0 1 10 99 999",
			numBlinks: 1,
			expectedStoneList: &StoneList{
				Stones: []Stone{
					{1}, {2024}, {1}, {0}, {9}, {9}, {2021976},
				},
			},
		},
		{
			text:      "125 17",
			numBlinks: 6,
			expectedStoneList: &StoneList{
				//2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
				Stones: []Stone{
					{2097446912}, {14168}, {4048}, {2}, {0}, {2}, {4}, {40}, {48}, {2024}, {40}, {48}, {80}, {96}, {2}, {8}, {6}, {7}, {6}, {0}, {3}, {2},
				},
			},
		},
	}

	for _, test := range testCases {
		stoneList := ParseStones(test.text)
		for i := 0; i < test.numBlinks; i++ {
			stoneList.Blink()
		}
		assert.True(t, reflect.DeepEqual(test.expectedStoneList, stoneList))
	}
}
