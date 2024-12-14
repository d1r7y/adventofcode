/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day01

import (
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseLocationIDs(t *testing.T) {
	type testCase struct {
		str           string
		expectedError error
		expectedLeft  []int
		expectedRight []int
	}
	testCases := []testCase{
		{
			str: `1 2
3 4
5 6
7 8
100 200
400 500
10000 100001
-1 -3
0 0`,
			expectedError: nil,
			expectedLeft:  []int{-1, 0, 1, 3, 5, 7, 100, 400, 10000},
			expectedRight: []int{-3, 0, 2, 4, 6, 8, 200, 500, 100001},
		},
	}

	for _, test := range testCases {
		left, right, err := ParseLocationIDs(test.str)
		assert.Equal(t, test.expectedError, err)
		assert.Equal(t, test.expectedLeft, left)
		assert.Equal(t, test.expectedRight, right)
	}
}

func TestTotalDistance(t *testing.T) {
	str := `3   4
4   3
2   5
1   3
3   9
3   3`

	left, right, err := ParseLocationIDs(str)
	assert.NoError(t, err)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := utilities.AbsoluteDifference(left[i], right[i])
		totalDistance += distance
	}

	assert.Equal(t, 11, totalDistance)
}

func TestCalculateSimilarity(t *testing.T) {
	str := `3   4
4   3
2   5
1   3
3   9
3   3`

	left, right, err := ParseLocationIDs(str)
	assert.NoError(t, err)

	similarity := CalculateSimilarity(left, right)
	assert.Equal(t, 31, similarity)
}
