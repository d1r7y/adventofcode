/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day03

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Day03Cmd represents the day03 command
var Day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: `Mull It Over`,
	Run: func(cmd *cobra.Command, args []string) {
		df, err := os.Open(utilities.GetInputPath(cmd))
		if err != nil {
			log.Fatal(err)
		}

		defer df.Close()

		fileContent, err := io.ReadAll(df)
		if err != nil {
			log.Fatal(err)
		}
		err = day(cmd, string(fileContent))
		if err != nil {
			log.Fatal(err)
		}
	},
}

var cmd *cobra.Command

type MulInstruction struct {
	factorA int
	factorB int
}

func ScanMulInstructions(fileContents string) []MulInstruction {
	instructions := make([]MulInstruction, 0)
	mulRE := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	for _, line := range strings.Split(fileContents, "\n") {
		mulMatches := mulRE.FindAllStringSubmatch(line, -1)

		for _, m := range mulMatches {
			factorA, err := strconv.Atoi(m[1])
			if err != nil {
				return instructions
			}

			factorB, err := strconv.Atoi(m[2])
			if err != nil {
				return instructions
			}

			instructions = append(instructions, MulInstruction{factorA: factorA, factorB: factorB})
		}
	}

	return instructions
}

func ScanEnabledMulInstructions(fileContents string) []MulInstruction {
	instructions := make([]MulInstruction, 0)
	mulRE := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)

	enabled := true
	for _, line := range strings.Split(fileContents, "\n") {
		mulMatches := mulRE.FindAllStringSubmatch(line, -1)
		for _, m := range mulMatches {
			if m[0] == "do()" {
				enabled = true
			} else if m[0] == "don't()" {
				enabled = false
			} else {
				factorA, err := strconv.Atoi(m[1])
				if err != nil {
					return instructions
				}

				factorB, err := strconv.Atoi(m[2])
				if err != nil {
					return instructions
				}

				if utilities.GetVerbosity(cmd) > 0 {
					if enabled {
						color.Set(color.FgGreen)
					} else {
						color.Set(color.FgRed)
					}

					fmt.Printf("%d*%d ", factorA, factorB)

					color.Unset()
				}

				if enabled {
					instructions = append(instructions, MulInstruction{factorA: factorA, factorB: factorB})
				}
			}
		}
	}

	if utilities.GetVerbosity(cmd) > 0 {
		fmt.Println()
	}

	return instructions
}

func SumMultiplicationInstructions(inst []MulInstruction) int {
	total := 0
	for _, i := range inst {
		total += i.factorA * i.factorB
	}

	return total
}

func day(command *cobra.Command, fileContents string) error {
	cmd = command

	// Part 1: It seems like the goal of the program is just to multiply some numbers. It
	// does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers.
	// For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly,
	// mul(123,4) would multiply 123 by 4.
	//
	// However, because the program's memory has been corrupted, there are also many invalid
	// characters that should be ignored, even if they look like part of a mul instruction.
	// Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.
	//
	// Scan the corrupted memory for uncorrupted mul instructions. What do you get if you add
	// up all of the results of the multiplications?

	totalSum := 0

	instructions := ScanMulInstructions(fileContents)
	totalSum += SumMultiplicationInstructions(instructions)

	fmt.Printf("Multiplications sum: %d\n", totalSum)

	// Part 2: As you scan through the corrupted memory, you notice that some of the conditional
	// statements are also still intact. If you handle some of the uncorrupted conditional
	// statements in the program, you might be able to get an even more accurate result.
	//
	// There are two new instructions you'll need to handle:
	//	- The do() instruction enables future mul instructions.
	//	- The don't() instruction disables future mul instructions.
	//
	// Only the most recent do() or don't() instruction applies. At the beginning of the program,
	// mul instructions are enabled.
	//
	// Handle the new instructions; what do you get if you add up all of the results of just the
	// enabled multiplications?

	totalSum = 0

	instructions = ScanEnabledMulInstructions(fileContents)
	totalSum += SumMultiplicationInstructions(instructions)

	fmt.Printf("Multiplications sum of enabled instructions: %d\n", totalSum)

	return nil
}
