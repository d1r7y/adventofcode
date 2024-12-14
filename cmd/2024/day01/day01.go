/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day01

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

// Day01Cmd represents the day01 command
var Day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: `Historian Hysteria`,
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

func ParseLocationIDs(fileContents string) ([]int, []int, error) {
	left := []int{}
	right := []int{}

	for _, line := range strings.Split(fileContents, "\n") {
		ids := utilities.ParseIntList(line)
		if len(ids) != 2 {
			return left, right, fmt.Errorf("invalid line '%s'", line)
		}

		left = append(left, ids[0])
		right = append(right, ids[1])
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right, nil
}

func CalculateSimilarity(left []int, right []int) int {
	hist := utilities.MakeHistogram(right)

	totalSimilarity := 0

	for _, id := range left {
		if count, ok := hist[id]; ok {
			totalSimilarity += id * count
		}
	}

	return totalSimilarity
}

func day(_ *cobra.Command, fileContents string) error {
	// Part 1: Pair up the smallest number in the left list with the
	// smallest number in the right list, then the second-smallest left
	// number with the second-smallest right number, and so on.
	//
	// Find the total distance between all the numbers.
	left, right, err := ParseLocationIDs(fileContents)
	if err != nil {
		return err
	}

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := utilities.AbsoluteDifference(left[i], right[i])
		totalDistance += distance
	}

	fmt.Printf("Total distance: %d.\n", totalDistance)

	// Part 2: This time, you'll need to figure out exactly how often each
	// number from the left list appears in the right list. Calculate a total
	// similarity score by adding up each number in the left list after multiplying
	// it by the number of times that number appears in the right list.

	similarity := CalculateSimilarity(left, right)

	fmt.Printf("Similarity: %d.\n", similarity)

	return nil
}
