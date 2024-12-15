/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day13

import (
	"reflect"
	"testing"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/stretchr/testify/assert"
)

func TestParseClawMachines(t *testing.T) {
	type testCase struct {
		text             string
		expectedMachines []ClawMachine
	}
	testCases := []testCase{
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			expectedMachines: []ClawMachine{
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(8400, 5400),
					MovementA:        utilities.NewPoint2D(94, 34),
					MovementB:        utilities.NewPoint2D(22, 67),
				},
			},
		},
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,
			expectedMachines: []ClawMachine{
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(8400, 5400),
					MovementA:        utilities.NewPoint2D(94, 34),
					MovementB:        utilities.NewPoint2D(22, 67),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(12748, 12176),
					MovementA:        utilities.NewPoint2D(26, 66),
					MovementB:        utilities.NewPoint2D(67, 21),
				},
			},
		},
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			expectedMachines: []ClawMachine{
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(8400, 5400),
					MovementA:        utilities.NewPoint2D(94, 34),
					MovementB:        utilities.NewPoint2D(22, 67),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(12748, 12176),
					MovementA:        utilities.NewPoint2D(26, 66),
					MovementB:        utilities.NewPoint2D(67, 21),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(7870, 6450),
					MovementA:        utilities.NewPoint2D(17, 86),
					MovementB:        utilities.NewPoint2D(84, 37),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(18641, 10279),
					MovementA:        utilities.NewPoint2D(69, 23),
					MovementB:        utilities.NewPoint2D(27, 71),
				},
			},
		},
	}

	for _, test := range testCases {
		machines := ParseClawMachines(test.text, false)
		assert.True(t, reflect.DeepEqual(test.expectedMachines, machines))
	}
}

func TestParseClawMachinesCorrection(t *testing.T) {
	type testCase struct {
		text             string
		expectedMachines []ClawMachine
	}
	testCases := []testCase{
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			expectedMachines: []ClawMachine{
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(10000000000000+8400, 10000000000000+5400),
					MovementA:        utilities.NewPoint2D(94, 34),
					MovementB:        utilities.NewPoint2D(22, 67),
				},
			},
		},
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,
			expectedMachines: []ClawMachine{
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(10000000000000+8400, 10000000000000+5400),
					MovementA:        utilities.NewPoint2D(94, 34),
					MovementB:        utilities.NewPoint2D(22, 67),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(10000000000000+12748, 10000000000000+12176),
					MovementA:        utilities.NewPoint2D(26, 66),
					MovementB:        utilities.NewPoint2D(67, 21),
				},
			},
		},
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			expectedMachines: []ClawMachine{
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(10000000000000+8400, 10000000000000+5400),
					MovementA:        utilities.NewPoint2D(94, 34),
					MovementB:        utilities.NewPoint2D(22, 67),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(10000000000000+12748, 10000000000000+12176),
					MovementA:        utilities.NewPoint2D(26, 66),
					MovementB:        utilities.NewPoint2D(67, 21),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(10000000000000+7870, 10000000000000+6450),
					MovementA:        utilities.NewPoint2D(17, 86),
					MovementB:        utilities.NewPoint2D(84, 37),
				},
				{
					StartingPosition: utilities.NewPoint2D(0, 0),
					PrizeLocation:    utilities.NewPoint2D(10000000000000+18641, 10000000000000+10279),
					MovementA:        utilities.NewPoint2D(69, 23),
					MovementB:        utilities.NewPoint2D(27, 71),
				},
			},
		},
	}

	for _, test := range testCases {
		machines := ParseClawMachines(test.text, true)
		assert.True(t, reflect.DeepEqual(test.expectedMachines, machines))
	}
}

func TestWinningPrizeCost(t *testing.T) {
	type testCase struct {
		text                     string
		expectedWinningPrizeCost int64
	}
	testCases := []testCase{
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			expectedWinningPrizeCost: 280,
		},
		{
			text: `Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,
			expectedWinningPrizeCost: 0,
		},
		{
			text: `Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`,
			expectedWinningPrizeCost: 200,
		},
		{
			text: `Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			expectedWinningPrizeCost: 0,
		},
	}

	for _, test := range testCases {
		machines := ParseClawMachines(test.text, false)
		assert.Equal(t, test.expectedWinningPrizeCost, machines[0].WinningPrizeCost())
	}
}

func TestTotalWinnablePrizeCost(t *testing.T) {
	type testCase struct {
		text                           string
		expectedTotalWinnablePrizeCost int64
	}
	testCases := []testCase{
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			expectedTotalWinnablePrizeCost: 480,
		},
	}

	for _, test := range testCases {
		machines := ParseClawMachines(test.text, false)
		totalWinnablePrizeCost := int64(0)

		for _, m := range machines {
			totalWinnablePrizeCost += m.WinningPrizeCost()
		}

		assert.Equal(t, test.expectedTotalWinnablePrizeCost, totalWinnablePrizeCost)
	}
}

func TestWinningPrizeCostCorrected(t *testing.T) {
	type testCase struct {
		text                     string
		expectedWinningPrizeCost int64
	}
	testCases := []testCase{
		{
			text: `Button A: X+17, Y+37
Button B: X+47, Y+12
Prize: X=12898, Y=9663`,
			expectedWinningPrizeCost: 814_332_248_346,
		},
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			expectedWinningPrizeCost: 0,
		},
		{
			text: `Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,
			expectedWinningPrizeCost: 459_236_326_669,
		},
		{
			text: `Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`,
			expectedWinningPrizeCost: 0,
		},
		{
			text: `Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			expectedWinningPrizeCost: 416_082_282_239,
		},
	}

	for _, test := range testCases {
		machines := ParseClawMachines(test.text, true)
		assert.Equal(t, test.expectedWinningPrizeCost, machines[0].WinningPrizeCost())
	}
}

func TestTotalWinnablePrizeCostCorrected(t *testing.T) {
	type testCase struct {
		text                           string
		expectedTotalWinnablePrizeCost int64
	}
	testCases := []testCase{
		{
			text: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			expectedTotalWinnablePrizeCost: 875_318_608_908,
		},
	}

	for _, test := range testCases {
		machines := ParseClawMachines(test.text, true)
		totalWinnablePrizeCost := int64(0)

		for _, m := range machines {
			totalWinnablePrizeCost += m.WinningPrizeCost()
		}

		assert.Equal(t, test.expectedTotalWinnablePrizeCost, totalWinnablePrizeCost)
	}
}
