/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanMulInstructions(t *testing.T) {
	type testCase struct {
		line                 string
		expectedInstructions []MulInstruction
	}
	testCases := []testCase{
		{
			line:                 "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expectedInstructions: []MulInstruction{{2, 4}, {5, 5}, {11, 8}, {8, 5}},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedInstructions, ScanMulInstructions(test.line))
	}
}

func TestScanEnabledMulInstructions(t *testing.T) {
	type testCase struct {
		line                 string
		expectedInstructions []MulInstruction
	}
	testCases := []testCase{
		{
			line:                 "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			expectedInstructions: []MulInstruction{{2, 4}, {8, 5}},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedInstructions, ScanEnabledMulInstructions(test.line))
	}
}

func TestSumMultiplicationInstructions(t *testing.T) {
	type testCase struct {
		instructions []MulInstruction
		expectedSum  int
	}
	testCases := []testCase{
		{
			instructions: []MulInstruction{{2, 4}, {5, 5}, {11, 8}, {8, 5}},
			expectedSum:  161,
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedSum, SumMultiplicationInstructions(test.instructions))
	}
}
