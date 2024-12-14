/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseReport(t *testing.T) {
	type testCase struct {
		line           string
		expectedError  bool
		expectedLevels []int
	}
	testCases := []testCase{
		{
			line:           "7 6 4 2 1",
			expectedError:  false,
			expectedLevels: []int{7, 6, 4, 2, 1},
		},
		{
			line:           "1 2 7 8 9",
			expectedError:  false,
			expectedLevels: []int{1, 2, 7, 8, 9},
		},
		{
			line:           "1 2",
			expectedError:  false,
			expectedLevels: []int{1, 2},
		},
		{
			line:           "1",
			expectedError:  true,
			expectedLevels: []int{},
		},
		{
			line:           "",
			expectedError:  true,
			expectedLevels: []int{},
		},
	}

	for _, test := range testCases {
		levels, err := ParseReport(test.line)
		if test.expectedError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expectedLevels, levels)
		}
	}
}

func TestGetReportDirection(t *testing.T) {
	type testCase struct {
		a                 int
		b                 int
		expectedDirection ReportDirection
	}
	testCases := []testCase{
		{
			a:                 1,
			b:                 2,
			expectedDirection: Increasing,
		},
		{
			a:                 2,
			b:                 1,
			expectedDirection: Decreasing,
		},
		{
			a:                 -1,
			b:                 1,
			expectedDirection: Increasing,
		},
		{
			a:                 1,
			b:                 -1,
			expectedDirection: Decreasing,
		},
		{
			a:                 5,
			b:                 5,
			expectedDirection: Steady,
		},
		{
			a:                 0,
			b:                 0,
			expectedDirection: Steady,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedDirection, GetReportDirection(test.a, test.b))
	}
}

func TestCheckReportSafety(t *testing.T) {
	type testCase struct {
		levels         []int
		expectedSafety bool
	}
	testCases := []testCase{
		{
			levels:         []int{7, 6, 4, 2, 1},
			expectedSafety: true,
		},
		{
			levels:         []int{1, 2, 7, 8, 9},
			expectedSafety: false,
		},
		{
			levels:         []int{9, 7, 6, 2, 1},
			expectedSafety: false,
		},
		{
			levels:         []int{1, 3, 2, 4, 5},
			expectedSafety: false,
		},
		{
			levels:         []int{8, 6, 4, 4, 1},
			expectedSafety: false,
		},
		{
			levels:         []int{1, 3, 6, 7, 9},
			expectedSafety: true,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedSafety, CheckReportSafety(test.levels))
	}
}

func TestCheckReportSafetyProblemDamper(t *testing.T) {
	type testCase struct {
		levels         []int
		expectedSafety bool
	}
	testCases := []testCase{
		{
			levels:         []int{7, 6, 4, 2, 1},
			expectedSafety: true,
		},
		{
			levels:         []int{1, 2, 7, 8, 9},
			expectedSafety: false,
		},
		{
			levels:         []int{9, 7, 6, 2, 1},
			expectedSafety: false,
		},
		{
			levels:         []int{1, 3, 2, 4, 5},
			expectedSafety: true,
		},
		{
			levels:         []int{8, 6, 4, 4, 1},
			expectedSafety: true,
		},
		{
			levels:         []int{1, 3, 6, 7, 9},
			expectedSafety: true,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedSafety, CheckReportSafetyProblemDamper(test.levels))
	}
}
