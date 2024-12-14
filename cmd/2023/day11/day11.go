/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyThree_day11

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day11Cmd represents the day11 command
var Day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: `Cosmic Expansion`,
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
		err = day(string(fileContent))
		if err != nil {
			log.Fatal(err)
		}
	},
}

type Universe struct {
	Bounds   utilities.Size2D
	Galaxies []utilities.Point2D
}

func NewUniverse() *Universe {
	return &Universe{Galaxies: make([]utilities.Point2D, 0)}
}

func ParseUniverse(lines []string) *Universe {
	universe := NewUniverse()

	for y, line := range lines {
		for x, s := range line {
			if s == '#' {
				// Galaxy.
				universe.Galaxies = append(universe.Galaxies, utilities.NewPoint2D(x, y))
			}

			if y == 0 {
				universe.Bounds.Width++
			}
		}

		universe.Bounds.Height++
	}

	return universe
}

func (u *Universe) Describe() string {
	description := ""
	for y := 0; y < u.Bounds.Height; y++ {
		str := make([]rune, u.Bounds.Width)

		// Prefill string with empty space
		for x := 0; x < u.Bounds.Width; x++ {
			str[x] = '.'
		}

		for _, g := range u.Galaxies {
			if g.Y == y {
				str[g.X] = '#'
			}
		}

		if description != "" {
			description += "\n"
		}

		description += string(str)
	}

	return description
}

func (u *Universe) UnpopulatedRows() []int {
	populatedRows := make(map[int]bool, 0)

	for _, g := range u.Galaxies {
		populatedRows[g.Y] = true
	}

	unpopulatedRows := make([]int, 0)

	for r := 0; r < u.Bounds.Height; r++ {
		if !populatedRows[r] {
			unpopulatedRows = append(unpopulatedRows, r)
		}
	}

	return unpopulatedRows
}

func (u *Universe) UnpopulatedColumns() []int {
	populatedColumns := make(map[int]bool, 0)

	for _, g := range u.Galaxies {
		populatedColumns[g.X] = true
	}

	unpopulatedColumns := make([]int, 0)

	for c := 0; c < u.Bounds.Width; c++ {
		if !populatedColumns[c] {
			unpopulatedColumns = append(unpopulatedColumns, c)
		}
	}

	return unpopulatedColumns
}

func (u *Universe) Expand(emptyValue int) {
	unpopulatedColums := u.UnpopulatedColumns()
	unpopulatedRows := u.UnpopulatedRows()

	// Need to expand the bounds.
	u.Bounds.Width += (emptyValue - 1) * len(unpopulatedColums)
	u.Bounds.Height += (emptyValue - 1) * len(unpopulatedRows)

	// Sort the unpopulated values in decreasing order.
	sort.Sort(sort.Reverse(sort.IntSlice(unpopulatedColums)))
	sort.Sort(sort.Reverse(sort.IntSlice(unpopulatedRows)))

	for _, uc := range unpopulatedColums {
		for gi := 0; gi < len(u.Galaxies); gi++ {
			if u.Galaxies[gi].X > uc {
				u.Galaxies[gi].X += (emptyValue - 1)
			}
		}
	}

	for _, ur := range unpopulatedRows {
		for gi := 0; gi < len(u.Galaxies); gi++ {
			if u.Galaxies[gi].Y > ur {
				u.Galaxies[gi].Y += (emptyValue - 1)
			}
		}
	}
}

func (u *Universe) GetNumGalaxies() int {
	return len(u.Galaxies)
}

func (u *Universe) GetGalaxy(id int) utilities.Point2D {
	return u.Galaxies[id]
}

func (u *Universe) GalaxyDistance(id1 int, id2 int) int {
	return utilities.ManhattanDistance(u.GetGalaxy(id1), u.GetGalaxy(id2))
}

func GetPartners(id int) []int {
	partners := make([]int, 0)

	for i := 0; i < id; i++ {
		partners = append(partners, i)
	}

	return partners
}

func day(fileContents string) error {
	universe := ParseUniverse(strings.Split(strings.TrimSpace(fileContents), "\n"))

	// Part 1: Expand the universe, then find the length of the shortest path between every pair
	// of galaxies. What is the sum of these lengths?
	universe.Expand(2)

	sumGalaxyDistances := 0

	for id := 0; id < universe.GetNumGalaxies(); id++ {
		for _, p := range GetPartners(id) {
			sumGalaxyDistances += universe.GalaxyDistance(id, p)
		}
	}

	log.Printf("Sum of shortest paths between all pairs of galaxies: %d\n", sumGalaxyDistances)

	olderUniverse := ParseUniverse(strings.Split(strings.TrimSpace(fileContents), "\n"))

	// Part 2: Starting with the same initial image, expand the universe according to these new rules,
	// then find the length of the shortest path between every pair of galaxies. What is the sum of these lengths?
	olderUniverse.Expand(1000000)

	sumOlderGalaxyDistances := 0

	for id := 0; id < olderUniverse.GetNumGalaxies(); id++ {
		for _, p := range GetPartners(id) {
			sumOlderGalaxyDistances += olderUniverse.GalaxyDistance(id, p)
		}
	}

	log.Printf("Sum of shortest paths between all pairs of galaxies: %d\n", sumOlderGalaxyDistances)

	return nil
}
