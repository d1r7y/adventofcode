/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day10

import (
	"reflect"
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseTopoMap(t *testing.T) {
	type testCase struct {
		text            string
		expectedTopoMap *TopoMap
	}
	testCases := []testCase{
		{
			text: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			expectedTopoMap: &TopoMap{
				Bounds: utilities.NewSize2D(8, 8),
				Trailheads: []utilities.Point2D{
					utilities.NewPoint2D(2, 0),
					utilities.NewPoint2D(4, 0),
					utilities.NewPoint2D(4, 2),
					utilities.NewPoint2D(6, 4),
					utilities.NewPoint2D(2, 5),
					utilities.NewPoint2D(5, 5),
					utilities.NewPoint2D(0, 6),
					utilities.NewPoint2D(6, 6),
					utilities.NewPoint2D(1, 7),
				},
				Columns: []Row{
					{8, 9, 0, 1, 0, 1, 2, 3},
					{7, 8, 1, 2, 1, 8, 7, 4},
					{8, 7, 4, 3, 0, 9, 6, 5},
					{9, 6, 5, 4, 9, 8, 7, 4},
					{4, 5, 6, 7, 8, 9, 0, 3},
					{3, 2, 0, 1, 9, 0, 1, 2},
					{0, 1, 3, 2, 9, 8, 0, 1},
					{1, 0, 4, 5, 6, 7, 3, 2},
				},
			},
		},
	}

	for _, test := range testCases {
		topoMap := ParseTopoMap(test.text)
		assert.True(t, reflect.DeepEqual(test.expectedTopoMap, topoMap))
	}
}

func TestHikeScores(t *testing.T) {
	type testCase struct {
		text                string
		expectedScores      Scores
		expectedTotalScores int
	}
	testCases := []testCase{
		{
			text: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			expectedScores: Scores{
				utilities.NewPoint2D(2, 0): 5,
				utilities.NewPoint2D(4, 0): 6,
				utilities.NewPoint2D(4, 2): 5,
				utilities.NewPoint2D(6, 4): 3,
				utilities.NewPoint2D(2, 5): 1,
				utilities.NewPoint2D(5, 5): 3,
				utilities.NewPoint2D(0, 6): 5,
				utilities.NewPoint2D(6, 6): 3,
				utilities.NewPoint2D(1, 7): 5,
			},
			expectedTotalScores: 36,
		},
	}

	for _, test := range testCases {
		topoMap := ParseTopoMap(test.text)

		scores := topoMap.HikeScores()
		assert.True(t, reflect.DeepEqual(test.expectedScores, scores))

		totalTrailheadScores := 0

		for _, score := range scores {
			totalTrailheadScores += score
		}

		assert.Equal(t, test.expectedTotalScores, totalTrailheadScores)
	}
}

func TestHikeRatings(t *testing.T) {
	type testCase struct {
		text                 string
		expectedRatings      Ratings
		expectedTotalRatings int
	}
	testCases := []testCase{
		{
			text: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
			expectedRatings: Ratings{
				utilities.NewPoint2D(2, 0): 20,
				utilities.NewPoint2D(4, 0): 24,
				utilities.NewPoint2D(4, 2): 10,
				utilities.NewPoint2D(6, 4): 4,
				utilities.NewPoint2D(2, 5): 1,
				utilities.NewPoint2D(5, 5): 4,
				utilities.NewPoint2D(0, 6): 5,
				utilities.NewPoint2D(6, 6): 8,
				utilities.NewPoint2D(1, 7): 5,
			},
			expectedTotalRatings: 81,
		},
		{
			text: `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`,
			expectedRatings: Ratings{
				utilities.NewPoint2D(5, 0): 3,
			},
			expectedTotalRatings: 3,
		},
	}

	for _, test := range testCases {
		topoMap := ParseTopoMap(test.text)

		ratings := topoMap.HikeRatings()
		assert.True(t, reflect.DeepEqual(test.expectedRatings, ratings))

		totalTrailheadRatings := 0

		for _, rating := range ratings {
			totalTrailheadRatings += rating
		}

		assert.Equal(t, test.expectedTotalRatings, totalTrailheadRatings)
	}
}
