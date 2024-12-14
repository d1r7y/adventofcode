/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day04

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day04Cmd represents the day04 command
var Day04Cmd = &cobra.Command{
	Use:   "day04",
	Short: `Ceres Search`,
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

type LetterGrid struct {
	Bounds utilities.Size2D
	Rows   [][]rune
}

func ParseLetterGrid(fileContents string) *LetterGrid {
	lg := &LetterGrid{}
	lg.Rows = make([][]rune, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		row := make([]rune, 0)

		for _, r := range line {
			row = append(row, r)
		}

		lg.Rows = append(lg.Rows, row)
	}

	lg.Bounds.Height = len(lg.Rows)
	lg.Bounds.Width = len(lg.Rows[0])

	return lg
}

func (lg *LetterGrid) GetLetter(position utilities.Point2D) rune {
	return lg.Rows[position.Y][position.X]
}

func (lg *LetterGrid) FindXMAS() []utilities.Point2D {
	locations := make([]utilities.Point2D, 0)

	for y := 0; y < lg.Bounds.Height; y++ {
		for x := 0; x < lg.Bounds.Width; x++ {

			if x < 1 {
				continue
			}
			if x >= lg.Bounds.Width-1 {
				continue
			}
			if y < 1 {
				continue
			}
			if y >= lg.Bounds.Height-1 {
				continue
			}

			currentLocation := utilities.NewPoint2D(x, y)

			if lg.GetLetter(currentLocation) != 'A' {
				continue
			}

			type Check struct {
				UpLeft    rune
				UpRight   rune
				DownLeft  rune
				DownRight rune
			}

			checkList := []Check{
				{'M', 'M', 'S', 'S'}, // MAS/MAS
				{'S', 'M', 'S', 'M'}, // SAM/MAS
				{'M', 'S', 'M', 'S'}, // MAS/SAM
				{'S', 'S', 'M', 'M'}, // SAM/SAM
			}

			for _, i := range checkList {
				if lg.GetLetter(currentLocation.UpLeft()) == i.UpLeft && lg.GetLetter(currentLocation.UpRight()) == i.UpRight &&
					lg.GetLetter(currentLocation.DownLeft()) == i.DownLeft && lg.GetLetter(currentLocation.DownRight()) == i.DownRight {
					locations = append(locations, currentLocation)
				}

			}
		}
	}

	utilities.SortPoints(locations)

	return locations
}

func (lg *LetterGrid) FindString(str string) []utilities.Point2D {
	locations := make([]utilities.Point2D, 0)

	for y := 0; y < lg.Bounds.Height; y++ {
		for x := 0; x < lg.Bounds.Width; x++ {

			currentLocation := utilities.NewPoint2D(x, y)

			checkString := func(loc utilities.Point2D, move func(utilities.Point2D) utilities.Point2D) {
				cl := loc
				for _, l := range str {
					if lg.GetLetter(cl) != l {
						return
					}
					cl = move(cl)
					// checkLocation = checkLocation.Right()
				}
				locations = append(locations, loc)
			}

			if x+len(str) <= lg.Bounds.Width {
				// Check for horizontal forward
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.Right()
				})
			}

			if (x+1)-len(str) >= 0 {
				// Check for horizontal backward
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.Left()
				})
			}

			if y+len(str) <= lg.Bounds.Height {
				// Check for vertical forward
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.Down()
				})
			}

			if (y+1)-len(str) >= 0 {
				// Check for vertical backward
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.Up()
				})
			}

			if x+len(str) <= lg.Bounds.Width && (y+1)-len(str) >= 0 {
				// Check for diagonal heading upper right
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.UpRight()
				})
			}

			if (x+1)-len(str) >= 0 && (y+1)-len(str) >= 0 {
				// Check for diagonal heading upper left
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.UpLeft()
				})
			}

			if x+len(str) <= lg.Bounds.Width && y+len(str) <= lg.Bounds.Height {
				// Check for diagonal heading lower right
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.DownRight()
				})
			}

			if (x+1)-len(str) >= 0 && y+len(str) <= lg.Bounds.Height {
				// Check for diagonal heading lower left
				checkString(currentLocation, func(p utilities.Point2D) utilities.Point2D {
					return p.DownLeft()
				})
			}
		}
	}

	utilities.SortPoints(locations)

	return locations
}

func day(_ *cobra.Command, fileContents string) error {
	// Part 1: "Looks like the Chief's not here. Next!" One of The Historians pulls out a
	// device and pushes the only button on it. After a brief flash, you recognize the interior
	// of the Ceres monitoring station!
	//
	// As the search for the Chief continues, a small Elf who lives on the station tugs on your
	// shirt; she'd like to know if you could help her with her word search (your puzzle input).
	// She only has to find one word: XMAS.
	//
	// This word search allows words to be horizontal, vertical, diagonal, written backwards, or
	// even overlapping other words. It's a little unusual, though, as you don't merely need to find
	// one instance of XMAS - you need to find all of them.
	//
	// Take a look at the little Elf's word search. How many times does XMAS appear?

	lg := ParseLetterGrid(fileContents)
	locations := lg.FindString("XMAS")

	fmt.Printf("'XMAS' appears: %d times\n", len(locations))

	// Part 2: Looking for the instructions, you flip over the word search to find that this isn't
	// actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the
	// shape of an X.
	//
	// Flip the word search from the instructions back over to the word search side and try again.
	// How many times does an X-MAS appear?

	locations = lg.FindXMAS()

	fmt.Printf("X-MAS appears: %d times\n", len(locations))

	return nil
}
