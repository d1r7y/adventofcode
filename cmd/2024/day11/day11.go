/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day11

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day11Cmd represents the day11 command
var Day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: `Plutonian Pebbles`,
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

		err := day(string(fileContents))
		if err != nil {
			log.Fatal(err)
		}
	},
}

var analytics bool
var startingStones string
var numBlinks int

func init() {
	Day11Cmd.Flags().BoolVarP(&analytics, "analytics", "a", false, "display analytics instead of solving challenge")
	Day11Cmd.Flags().StringVarP(&startingStones, "stones", "s", "", "starting stone list")
	Day11Cmd.Flags().IntVarP(&numBlinks, "blinks", "b", 20, "number of blinks")
}

type Stone struct {
	Value int
}

type StoneList struct {
	Stones []Stone
}

func SplitDigits(number int) (int, int) {
	divisor := 10

	for {
		if number/divisor > divisor {
			divisor *= 10
		} else {
			return number / divisor, number % divisor
		}
	}
}

func (sl *StoneList) Blink() {
	for i := 0; i < len(sl.Stones); i++ {
		if sl.Stones[i].Value == 0 {
			sl.Stones[i].Value = 1
		} else if utilities.DigitCount(sl.Stones[i].Value)%2 == 0 {
			left, right := SplitDigits(sl.Stones[i].Value)
			sl.Stones[i].Value = left

			// Insert new stone
			newStone := Stone{right}
			sl.Stones = slices.Insert(sl.Stones, i+1, newStone)
			i++
		} else {
			sl.Stones[i].Value = sl.Stones[i].Value * 2024
		}
	}
}

func ParseStones(fileContents string) *StoneList {
	stoneList := &StoneList{}
	stoneList.Stones = make([]Stone, 0)

	for _, number := range utilities.ParseIntList(fileContents) {
		stone := Stone{
			Value: number,
		}

		stoneList.Stones = append(stoneList.Stones, stone)
	}

	return stoneList
}

func day(fileContents string) error {
	// Part 1: The ancient civilization on Pluto was known for its ability to manipulate
	// spacetime, and while The Historians explore their infinite corridors, you've
	// noticed a strange set of physics-defying stones.
	//
	// At first glance, they seem like normal stones: they're arranged in a perfectly
	// straight line, and each stone has a number engraved on it.
	//
	// The strange part is that every time you blink, the stones change.
	//
	// Sometimes, the number engraved on a stone changes. Other times, a stone might split
	// in two, causing all the other stones to shift over a bit to make room in their perfectly
	// straight line.
	//
	// As you observe them for a while, you find that the stones have a consistent behavior.
	// Every time you blink, the stones each simultaneously change according to the first
	// applicable rule in this list:
	//
	//	- If the stone is engraved with the number 0, it is replaced by a stone engraved with the
	//		number 1.
	//	- If the stone is engraved with a number that has an even number of digits, it is replaced
	//		by two stones. The left half of the digits are engraved on the new left stone, and the
	//		right half of the digits are engraved on the new right stone. (The new numbers don't keep
	//		extra leading zeroes: 1000 would become stones 10 and 0.)
	//	- If none of the other rules apply, the stone is replaced by a new stone; the old stone's
	//		number multiplied by 2024 is engraved on the new stone.
	//
	// No matter how the stones change, their order is preserved, and they stay on their perfectly
	// straight line.
	//
	// How will the stones evolve if you keep blinking at them? You take a note of the number engraved
	// on each stone in the line (your puzzle input).
	//
	// Consider the arrangement of stones in front of you. How many stones will you have after
	// blinking 25 times?

	if analytics {
		if startingStones != "" {
			stoneList := ParseStones(startingStones)

			fmt.Print("Stone list size after blinks: ")

			for blinks := 1; blinks <= numBlinks; blinks++ {
				stoneList.Blink()
				fmt.Printf("%5d ", len(stoneList.Stones))
			}

			fmt.Println()
		} else {
			for i := 0; i < 10; i++ {
				stoneList := ParseStones(fmt.Sprintf("%d", i))

				fmt.Printf("[%d] Stone list size after blinks: ", i)

				for blinks := 1; blinks <= numBlinks; blinks++ {
					stoneList.Blink()
					fmt.Printf("%5d ", len(stoneList.Stones))
				}

				fmt.Println()
			}
		}
	} else {
		stoneList := ParseStones(fileContents)

		for i := 0; i < 25; i++ {
			stoneList.Blink()
		}

		fmt.Printf("Stone list size after 25 blinks: %d\n", len(stoneList.Stones))

		stoneList2 := ParseStones(fileContents)

		for i := 0; i < 75; i++ {
			stoneList2.Blink()
		}

		fmt.Printf("Stone list size after 75 blinks: %d\n", len(stoneList2.Stones))
	}

	return nil
}
