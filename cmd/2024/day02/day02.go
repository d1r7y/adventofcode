/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day02

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day02Cmd represents the day02 command
var Day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: `Red-Nosed Reports`,
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

type ReportDirection int

const (
	Increasing ReportDirection = iota
	Decreasing
	Steady
)

func ParseReport(line string) ([]int, error) {
	levels := utilities.ParseIntList(line)

	if len(levels) < 2 {
		return levels, fmt.Errorf("too few levels in report")
	}

	return levels, nil
}

func GetReportDirection(a int, b int) ReportDirection {
	if a > b {
		return Decreasing
	} else if a < b {
		return Increasing
	} else {
		return Steady
	}
}

func CheckReportSafety(levels []int) bool {
	// Report must have a minimum of two levels.
	if len(levels) < 2 {
		return false
	}

	l1 := levels[0]
	l2 := levels[1]

	checkLevelPair := func(a int, b int, dir ReportDirection) bool {
		if dir == Steady {
			return false
		}
		if dir == Increasing {
			if b-a > 3 {
				return false
			}
		} else {
			if a-b > 3 {
				return false
			}
		}

		return true
	}

	startingDirection := GetReportDirection(l1, l2)
	if !checkLevelPair(l1, l2, startingDirection) {
		return false
	}

	for i := 2; i < len(levels); i++ {
		l1 = l2
		l2 = levels[i]

		// Check we're moving in the same initial direction.
		currentDirection := GetReportDirection(l1, l2)
		if currentDirection != startingDirection {
			return false
		}

		if !checkLevelPair(l1, l2, startingDirection) {
			return false
		}
	}

	return true
}

func CheckReportSafetyProblemDamper(levels []int) bool {
	if CheckReportSafety(levels) {
		return true
	}

	// Report isn't safe.  Try removing one of the levels and see if it becomes safe.

	for index := 0; index < len(levels); index++ {
		newLevels := slices.Clone(levels)
		newLevels = slices.Delete(newLevels, index, index+1)

		if CheckReportSafety(newLevels) {
			return true
		}
	}

	return false
}

func day(_ *cobra.Command, fileContents string) error {
	// Part 1: The engineers are trying to figure out which reports are safe.
	// The Red-Nosed reactor safety systems can only tolerate levels that are
	// either gradually increasing or gradually decreasing. So, a report only
	// counts as safe if both of the following are true:
	//
	//	- The levels are either all increasing or all decreasing.
	//	- Any two adjacent levels differ by at least one and at most three
	//
	// Analyze the unusual data from the engineers. How many reports are safe?

	numSafeReports := 0
	numDampenedSafeReports := 0
	for _, line := range strings.Split(fileContents, "\n") {
		levels, err := ParseReport(line)
		if err != nil {
			return err
		}

		if CheckReportSafety(levels) {
			numSafeReports++
		}

		if CheckReportSafetyProblemDamper(levels) {
			numDampenedSafeReports++
		}
	}

	fmt.Printf("Number of safe reports: %d.\n", numSafeReports)

	// Part 2: The Problem Dampener is a reactor-mounted module that lets the reactor
	// safety systems tolerate a single bad level in what would otherwise be a safe
	// report. It's like the bad level never happened!
	//
	// Now, the same rules apply as before, except if removing a single level from an
	// unsafe report would make it safe, the report instead counts as safe.
	//
	// Update your analysis by handling situations where the Problem Dampener can remove
	// a single level from unsafe reports. How many reports are now safe?

	fmt.Printf("Number of dampened safe reports: %d.\n", numDampenedSafeReports)

	return nil
}
