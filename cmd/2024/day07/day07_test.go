/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day07

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEquation(t *testing.T) {
	type testCase struct {
		line             string
		expectedEquation *Equation
	}
	testCases := []testCase{
		{
			line:             "190: 10 19",
			expectedEquation: &Equation{190, []int64{10, 19}},
		},
		{
			line:             "3267: 81 40 27",
			expectedEquation: &Equation{3267, []int64{81, 40, 27}},
		},
		{
			line:             "83: 17 5",
			expectedEquation: &Equation{83, []int64{17, 5}},
		},
		{
			line:             "156: 15 6",
			expectedEquation: &Equation{156, []int64{15, 6}},
		},
		{
			line:             "7290: 6 8 6 15",
			expectedEquation: &Equation{7290, []int64{6, 8, 6, 15}},
		},
		{
			line:             "161011: 16 10 13",
			expectedEquation: &Equation{161011, []int64{16, 10, 13}},
		},
		{
			line:             "192: 17 8 14",
			expectedEquation: &Equation{192, []int64{17, 8, 14}},
		},
		{
			line:             "21037: 9 7 18 13",
			expectedEquation: &Equation{21037, []int64{9, 7, 18, 13}},
		},
		{
			line:             "292: 11 6 16 20",
			expectedEquation: &Equation{292, []int64{11, 6, 16, 20}},
		},
		{
			line:             "1: 2 3 4 5 6 7 8 9 10",
			expectedEquation: &Equation{1, []int64{2, 3, 4, 5, 6, 7, 8, 9, 10}},
		},
	}

	for _, test := range testCases {
		assert.Equal(t, test.expectedEquation, ParseEquation(test.line))
	}
}

func TestEvaluateValidity(t *testing.T) {
	type testCase struct {
		line             string
		expectedValidity bool
	}
	testCases := []testCase{
		{
			line:             "190: 10 19",
			expectedValidity: true,
		},
		{
			line:             "3267: 81 40 27",
			expectedValidity: true,
		},
		{
			line:             "83: 17 5",
			expectedValidity: false,
		},
		{
			line:             "156: 15 6",
			expectedValidity: false,
		},
		{
			line:             "7290: 6 8 6 15",
			expectedValidity: false,
		},
		{
			line:             "161011: 16 10 13",
			expectedValidity: false,
		},
		{
			line:             "192: 17 8 14",
			expectedValidity: false,
		},
		{
			line:             "21037: 9 7 18 13",
			expectedValidity: false,
		},
		{
			line:             "292: 11 6 16 20",
			expectedValidity: true,
		},
	}

	for _, test := range testCases {
		e := ParseEquation(test.line)
		assert.Equal(t, test.expectedValidity, e.EvaluateValidity([]Operator{addOp, multOp}))
	}
}

func TestEvaluateValidityConcatenate(t *testing.T) {
	type testCase struct {
		line             string
		expectedValidity bool
	}
	testCases := []testCase{
		{
			line:             "62080: 86 718 329",
			expectedValidity: false,
		},
		{
			line:             "190: 10 19",
			expectedValidity: true,
		},
		{
			line:             "3267: 81 40 27",
			expectedValidity: true,
		},
		{
			line:             "83: 17 5",
			expectedValidity: false,
		},
		{
			line:             "156: 15 6",
			expectedValidity: true,
		},
		{
			line:             "7290: 6 8 6 15",
			expectedValidity: true,
		},
		{
			line:             "161011: 16 10 13",
			expectedValidity: false,
		},
		{
			line:             "192: 17 8 14",
			expectedValidity: true,
		},
		{
			line:             "21037: 9 7 18 13",
			expectedValidity: false,
		},
		{
			line:             "292: 11 6 16 20",
			expectedValidity: true,
		},
	}

	for _, test := range testCases {
		e := ParseEquation(test.line)
		assert.Equal(t, test.expectedValidity, e.EvaluateValidity([]Operator{addOp, multOp, concatOp}))
	}
}

func TestTotalCalibration(t *testing.T) {
	type testCase struct {
		text                     string
		expectedTotalCalibration int64
	}
	testCases := []testCase{
		{
			text: `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			expectedTotalCalibration: 3749,
		},
	}

	for _, test := range testCases {
		equations := make([]*Equation, 0)

		for _, line := range strings.Split(test.text, "\n") {
			equation := ParseEquation(line)

			equations = append(equations, equation)
		}

		totalCalibrationResult := int64(0)

		for _, equation := range equations {
			if equation.EvaluateValidity([]Operator{addOp, multOp}) {
				totalCalibrationResult += equation.TestValue
			}
		}

		assert.Equal(t, test.expectedTotalCalibration, totalCalibrationResult)
	}
}

func TestTotalCalibrationConcat(t *testing.T) {
	type testCase struct {
		text                     string
		expectedTotalCalibration int64
	}
	testCases := []testCase{
		{
			text: `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
			expectedTotalCalibration: 11387,
		},
	}

	for _, test := range testCases {
		equations := make([]*Equation, 0)

		for _, line := range strings.Split(test.text, "\n") {
			equation := ParseEquation(line)

			equations = append(equations, equation)
		}

		totalCalibrationResult := int64(0)

		for _, equation := range equations {
			if equation.EvaluateValidity([]Operator{addOp, multOp, concatOp}) {
				totalCalibrationResult += equation.TestValue
			}
		}

		assert.Equal(t, test.expectedTotalCalibration, totalCalibrationResult)
	}
}
