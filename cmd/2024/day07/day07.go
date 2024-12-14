/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day07

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day07Cmd represents the day07 command
var Day07Cmd = &cobra.Command{
	Use:   "day07",
	Short: `Bridge Repair`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileContent []byte
		var err error

		if utilities.GetInputPath(cmd) != "" {
			df, err := os.Open(utilities.GetInputPath(cmd))
			if err != nil {
				log.Fatal(err)
			}

			defer df.Close()

			fileContent, err = io.ReadAll(df)
			if err != nil {
				log.Fatal(err)
			}
		} else if equation != "" {
			fileContent = []byte(equation)
		}

		if fileContent != nil {
			err = day(cmd, string(fileContent))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

var equation string

func init() {
	Day07Cmd.Flags().StringVarP(&equation, "equation", "e", "", "Equation")
}

type Equation struct {
	TestValue int64
	Numbers   []int64
}

type Operator struct {
	n string
	f func(a, b int64) int64
}

var cmd *cobra.Command

var concatOp = Operator{"||", func(a, b int64) int64 {
	return utilities.Concatenate(a, b)
}}

var addOp = Operator{"+", func(a, b int64) int64 {
	return a + b
}}

var multOp = Operator{"*", func(a, b int64) int64 {
	return a * b
}}

func (e *Equation) EvaluateValidity(operators []Operator) bool {

	var evaluate func(equation string, requiredValue int64, currentValue int64, currentIndex int) bool

	evaluate = func(equation string, requiredValue int64, currentValue int64, currentIndex int) bool {
		// Reached the end of the numbers?
		if currentIndex == len(e.Numbers) {
			if utilities.GetVerbosity(cmd) > 2 {
				if currentValue == requiredValue {
					fmt.Printf("%d == %s\n", requiredValue, equation)
				} else {
					fmt.Printf("%d != %s (%d)\n", requiredValue, equation, currentValue)
				}
			}
			return currentValue == requiredValue
		}

		// Because we don't have subtraction, division, or negative numbers, if currentValue > requiredValue,
		// there's no point in continuing.
		if currentValue > requiredValue {
			return false
		}

		for _, o := range operators {
			if evaluate(fmt.Sprintf("%s %s %d", equation, o.n, e.Numbers[currentIndex]), requiredValue, o.f(currentValue, e.Numbers[currentIndex]), currentIndex+1) {
				return true
			}
		}

		return false
	}

	return evaluate(fmt.Sprintf("%d", e.Numbers[0]), e.TestValue, e.Numbers[0], 1)
}

func ParseEquation(line string) *Equation {
	equation := &Equation{}
	equation.Numbers = make([]int64, 0)

	parsedInts := utilities.ParseIntList(line)

	equation.TestValue = int64(parsedInts[0])

	for i := 1; i < len(parsedInts); i++ {
		equation.Numbers = append(equation.Numbers, int64(parsedInts[i]))
	}

	return equation
}

func SprintEquation(e *Equation) string {
	str := fmt.Sprintf("%d: ", e.TestValue)

	for _, n := range e.Numbers {
		str += fmt.Sprintf("%d ", n)
	}

	return str
}

func day(command *cobra.Command, fileContents string) error {
	cmd = command

	// Part 1: You ask how long it'll take; the engineers tell you that it only needs final calibrations,
	// but some young elephants were playing nearby and stole all the operators from their calibration
	// equations! They could finish the calibrations if only someone could determine which test values
	// could possibly be produced by placing any combination of operators into their calibration equations
	// (your puzzle input).
	//
	// Each line represents a single equation. The test value appears before the colon on each line; it is
	// your job to determine whether the remaining numbers can be combined with operators to produce the test value.
	//
	// Operators are always evaluated left-to-right, not according to precedence rules. Furthermore, numbers
	// in the equations cannot be rearranged. Glancing into the jungle, you can see elephants holding two
	// different types of operators: add (+) and multiply (*).
	//
	// The engineers just need the total calibration result, which is the sum of the test values from just
	// the equations that could possibly be true.
	//
	// Determine which equations could possibly be true. What is their total calibration result?

	equations := make([]*Equation, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		equation := ParseEquation(line)

		equations = append(equations, equation)
	}

	validEquations := make([]bool, len(equations))

	totalCalibrationResult := int64(0)

	for i, equation := range equations {
		if equation.EvaluateValidity([]Operator{addOp, multOp}) {
			totalCalibrationResult += equation.TestValue
			validEquations[i] = true
		}
	}

	fmt.Printf("Total calibration result: %d\n", totalCalibrationResult)

	// Part 2: The engineers seem concerned; the total calibration result you gave them is nowhere close to
	// being within safety tolerances. Just then, you spot your mistake: some well-hidden elephants are holding
	// a third type of operator.
	//
	// The concatenation operator (||) combines the digits from its left and right inputs into a single number.
	// For example, 12 || 345 would become 12345. All operators are still evaluated left-to-right.
	//
	// Using your new knowledge of elephant hiding spots, determine which equations could possibly be true.
	// What is their total calibration result?

	totalCalibrationConcatResult := int64(0)

	for i, equation := range equations {
		if equation.EvaluateValidity([]Operator{concatOp, addOp, multOp}) {
			if utilities.GetVerbosity(cmd) > 1 {
				if !validEquations[i] {
					// This equation wasn't valid before but it is now with the additional concatenation operator,
					// log it.
					fmt.Printf("%snow valid\n", SprintEquation(equation))
				}
			}

			totalCalibrationConcatResult += int64(equation.TestValue)
		} else {
			if utilities.GetVerbosity(cmd) > 0 {
				fmt.Printf("%snot valid\n", SprintEquation(equation))
			}
		}
	}

	fmt.Printf("Total calibration result with concatenation operator: %d\n", totalCalibrationConcatResult)

	return nil
}
