/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day13

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
	"gonum.org/v1/gonum/mat"
)

// Day13Cmd represents the day13 command
var Day13Cmd = &cobra.Command{
	Use:   "day13",
	Short: `Claw Contraption`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := utilities.GetInputPath(cmd)
		var fileContents = ""

		if inputPath != "" {
			var err error
			df, err := os.Open(inputPath)
			if err != nil {
				log.Fatal(err)
			}

			defer df.Close()

			fileBytes, err := io.ReadAll(df)
			if err != nil {
				log.Fatal(err)
			}

			fileContents = string(fileBytes)
		}

		err := day(cmd, string(fileContents))
		if err != nil {
			log.Fatal(err)
		}
	},
}

type ClawMachine struct {
	StartingPosition utilities.Point2D
	PrizeLocation    utilities.Point2D
	MovementA        utilities.Point2D
	MovementB        utilities.Point2D
}

func (c ClawMachine) WinningPrizeCost() int64 {
	A := mat.NewDense(2, 2, []float64{float64(c.MovementA.X), float64(c.MovementB.X), float64(c.MovementA.Y), float64(c.MovementB.Y)})
	b := mat.NewVecDense(2, []float64{float64(c.PrizeLocation.X), float64(c.PrizeLocation.Y)})

	var x mat.VecDense
	if err := x.SolveVec(A, b); err != nil {
		return 0
	}

	isInteger := func(a float64) (int64, bool) {
		epsilon := 1e-4
		if _, fractional := math.Modf(math.Abs(a)); fractional < epsilon || fractional > 1-epsilon {
			return int64(math.Round(a)), true
		}

		return 0, false
	}

	Ap, valid := isInteger(x.AtVec(0))
	if !valid {
		return 0
	}

	Bp, valid := isInteger(x.AtVec(1))
	if !valid {
		return 0
	}

	gameCost := func(Ap, Bp int64) int64 {
		return 3*Ap + 1*Bp
	}

	return gameCost(Ap, Bp)
}

func ParseClawMachines(fileContents string, correctPrizePosition bool) []ClawMachine {
	machines := make([]ClawMachine, 0)

	machineRE := regexp.MustCompile(`(?m)Button A: X\+([0-9]+), Y\+([0-9]+)\nButton B: X\+([0-9]+), Y\+([0-9]+)\nPrize: X=([0-9]+), Y=([0-9]+)`)
	machineMatches := machineRE.FindAllStringSubmatch(fileContents, -1)

	for i := 0; i < len(machineMatches); i++ {
		prizeLocationX, err := strconv.Atoi(machineMatches[i][5])
		if err != nil {
			return machines
		}
		if correctPrizePosition {
			prizeLocationX += 10000000000000
		}

		prizeLocationY, err := strconv.Atoi(machineMatches[i][6])
		if err != nil {
			return machines
		}
		if correctPrizePosition {
			prizeLocationY += 10000000000000
		}

		buttonAHorizontalMovement, err := strconv.Atoi(machineMatches[i][1])
		if err != nil {
			return machines
		}

		buttonAVerticalMovement, err := strconv.Atoi(machineMatches[i][2])
		if err != nil {
			return machines
		}

		buttonBHorizontalMovement, err := strconv.Atoi(machineMatches[i][3])
		if err != nil {
			return machines
		}

		buttonBVerticalMovement, err := strconv.Atoi(machineMatches[i][4])
		if err != nil {
			return machines
		}

		machine := ClawMachine{
			StartingPosition: utilities.NewPoint2D(0, 0),
			PrizeLocation:    utilities.NewPoint2D(prizeLocationX, prizeLocationY),
			MovementA:        utilities.NewPoint2D(buttonAHorizontalMovement, buttonAVerticalMovement),
			MovementB:        utilities.NewPoint2D(buttonBHorizontalMovement, buttonBVerticalMovement),
		}

		machines = append(machines, machine)
	}

	return machines
}

func day(command *cobra.Command, fileContents string) error {
	// Part 1: Next up: the lobby of a resort on a tropical island. The Historians
	// take a moment to admire the hexagonal floor tiles before spreading out.
	//
	// Fortunately, it looks like the resort has a new arcade! Maybe you can win some
	// prizes from the claw machines?
	//
	// The claw machines here are a little unusual. Instead of a joystick or directional
	// buttons to control the claw, these machines have two buttons labeled A and B. Worse,
	// you can't just put in a token and play; it costs 3 tokens to push the A button and
	// 1 token to push the B button.
	//
	// With a little experimentation, you figure out that each machine's buttons are configured
	// to move the claw a specific amount to the right (along the X axis) and a specific
	// amount forward (along the Y axis) each time that button is pressed.
	//
	// Each machine contains one prize; to win the prize, the claw must be positioned exactly
	// above the prize on both the X and Y axes.
	//
	// You wonder: what is the smallest number of tokens you would have to spend to win as many
	// prizes as possible? You assemble a list of every machine's button behavior and prize
	// location (your puzzle input).
	//
	// You estimate that each button would need to be pressed no more than 100 times to win a
	// prize. How else would someone be expected to play?
	//
	// Figure out how to win as many prizes as possible. What is the fewest tokens you would
	// have to spend to win all possible prizes?

	machines := ParseClawMachines(fileContents, false)

	totalWinnablePrizeCost := int64(0)

	for _, m := range machines {
		totalWinnablePrizeCost += m.WinningPrizeCost()
	}

	fmt.Printf("Fewest number of tokens spent to win all possible prizes: %d\n", totalWinnablePrizeCost)

	// Part 2: As you go to win the first prize, you discover that the claw is nowhere near
	// where you expected it would be. Due to a unit conversion error in your measurements,
	// the position of every prize is actually 10000000000000 higher on both the X and Y axis!
	//
	// Add 10000000000000 to the X and Y position of every prize.
	//
	// Using the corrected prize coordinates, figure out how to win as many prizes as possible.
	//
	// What is the fewest tokens you would have to spend to win all possible prizes?

	machinesCorrected := ParseClawMachines(fileContents, true)

	totalWinnablePrizeCostCorrected := int64(0)

	totalWinnablePrizeCost = 0
	totalUnsolvable := 0

	for _, m := range machinesCorrected {
		cost := m.WinningPrizeCost()

		var solvability string

		if cost == 0 {
			solvability = "Unsolvable"
			totalUnsolvable++
		} else {
			solvability = "Solvable"
		}

		if utilities.GetVerbosity(command) > 0 {
			fmt.Printf("%s: %dx + %dy = %d; %dx + %dy = %d\n", solvability, m.MovementA.X, m.MovementB.X, m.PrizeLocation.X, m.MovementA.Y, m.MovementB.Y, m.PrizeLocation.Y)
		}

		totalWinnablePrizeCostCorrected += cost
	}

	if utilities.GetVerbosity(command) > 0 {
		fmt.Printf("%d unsolvable with corrected prize coordinates\n", totalUnsolvable)
	}

	fmt.Printf("Fewest number of tokens spent to win all possible prizes with corrected prize coordinates: %d\n", totalWinnablePrizeCostCorrected)

	return nil
}
