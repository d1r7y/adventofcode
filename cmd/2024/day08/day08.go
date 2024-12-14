/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day08

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
)

// Day08Cmd represents the day08 command
var Day08Cmd = &cobra.Command{
	Use:   "day08",
	Short: `Resonant Collinearity`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileContent []byte
		var err error

		df, err := os.Open(utilities.GetInputPath(cmd))
		if err != nil {
			log.Fatal(err)
		}

		defer df.Close()

		fileContent, err = io.ReadAll(df)
		if err != nil {
			log.Fatal(err)
		}

		if fileContent != nil {
			err = day(cmd, string(fileContent))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

type AntennaLocations map[rune][]utilities.Point2D
type AntinodeLocations map[utilities.Point2D]bool

type AntennaMap struct {
	Bounds    utilities.Size2D
	Antennas  AntennaLocations
	Antinodes AntinodeLocations
}

var cmd *cobra.Command

func GenerateCollinearLocations(pointA utilities.Point2D, pointB utilities.Point2D, bounds utilities.Size2D) []utilities.Point2D {

	m, b := utilities.CalculateSlopeIntercept(pointA, pointB)

	collinearLocations := make([]utilities.Point2D, 0)

	for x := 0; x < bounds.Width; x++ {
		y := m*float64(x) + b

		// Only handle whole number coordinates.
		if _, fractional := math.Modf(math.Abs(y)); fractional < 1e-6 || fractional > 1-1e-6 {
			intY := int(math.Round(y))
			// And values of Y within bounds.
			if intY >= 0 && intY < bounds.Height {

				newLocation := utilities.NewPoint2D(x, intY)
				collinearLocations = append(collinearLocations, newLocation)
			}
		}
	}

	return collinearLocations
}

func ParseAntennaMap(fileContents string, harmonics bool) *AntennaMap {
	antennaMap := &AntennaMap{}

	antennaMap.Antennas = make(AntennaLocations)
	antennaMap.Antinodes = make(AntinodeLocations)

	for y, line := range strings.Split(fileContents, "\n") {
		antennaMap.Bounds.Width = 0
		for x, c := range line {
			currentLocation := utilities.NewPoint2D(x, y)

			if c != '.' {
				if loc, ok := antennaMap.Antennas[c]; ok {
					loc = append(loc, currentLocation)
					antennaMap.Antennas[c] = loc
				} else {
					loc := []utilities.Point2D{currentLocation}
					antennaMap.Antennas[c] = loc
				}
			}
			antennaMap.Bounds.Width++
		}
		antennaMap.Bounds.Height++
	}

	for _, locations := range antennaMap.Antennas {
		pairs := utilities.GenerateUniquePointPairs(locations)
		for _, pair := range pairs {
			// Generate collinear points.
			collinearPoints := GenerateCollinearLocations(pair.One, pair.Two, antennaMap.Bounds)

			if utilities.GetVerbosity(cmd) > 0 {
				fmt.Printf("Collinear points for (%d,%d) (%d,%d): %# v\n", pair.One.X, pair.One.Y, pair.Two.X, pair.Two.Y, pretty.Formatter(collinearPoints))
			}

			// Find antinode locations
			for i := 0; i < len(collinearPoints); i++ {
				point := collinearPoints[i]
				if harmonics {
					antennaMap.Antinodes[point] = true
				} else {
					if point == pair.One {
						// This is the first point of pair.
						// First antinode is the previous collinear point.
						if i > 0 {
							if utilities.GetVerbosity(cmd) > 0 {
								fmt.Printf("Antinode location: (%d,%d)\n", collinearPoints[i-1].X, collinearPoints[i-1].Y)
							}
							antennaMap.Antinodes[collinearPoints[i-1]] = true
						}
					} else if point == pair.Two {
						// This is the second point of pair.
						// Second antinode is the subsequent collinear point.
						if i < len(collinearPoints)-1 {
							if utilities.GetVerbosity(cmd) > 0 {
								fmt.Printf("Antinode location: (%d,%d)\n", collinearPoints[i+1].X, collinearPoints[i+1].Y)
							}
							antennaMap.Antinodes[collinearPoints[i+1]] = true
						}
						break
					}
				}
			}
		}
	}

	return antennaMap
}

func day(command *cobra.Command, fileContents string) error {
	cmd = command

	// Part 1: While The Historians do their thing, you take a look at the familiar huge
	// antenna. Much to your surprise, it seems to have been reconfigured to emit a signal
	// that makes people 0.1% more likely to buy Easter Bunny brand Imitation Mediocre
	// Chocolate as a Christmas gift! Unthinkable!
	//
	// Scanning across the city, you find that there are actually many such antennas. Each
	// antenna is tuned to a specific frequency indicated by a single lowercase letter,
	// uppercase letter, or digit.
	//
	// The signal only applies its nefarious effect at specific antinodes based on the resonant
	// frequencies of the antennas. In particular, an antinode occurs at any point that is
	// perfectly in line with two antennas of the same frequency - but only when one of the
	// antennas is twice as far away as the other. This means that for any pair of antennas with
	// the same frequency, there are two antinodes, one on either side of them.
	//
	// Calculate the impact of the signal. How many unique locations within the bounds of the map contain an antinode?

	antennaMap := ParseAntennaMap(fileContents, false)

	if utilities.GetVerbosity(cmd) > 0 {
		fmt.Printf("%# v\n", pretty.Formatter(antennaMap))
	}

	fmt.Printf("Number of unique antinode locations: %d\n", len(antennaMap.Antinodes))

	// Part 2: Watching over your shoulder as you work, one of The Historians asks if you took
	// the effects of resonant harmonics into your calculations.
	//
	// Whoops!
	//
	// After updating your model, it turns out that an antinode occurs at any grid position exactly in
	// line with at least two antennas of the same frequency, regardless of distance. This means that
	// some of the new antinodes will occur at the position of each antenna (unless that antenna is the
	// only one of its frequency).
	//
	// Calculate the impact of the signal using this updated model. How many unique locations within the
	// bounds of the map contain an antinode?

	antennaMapHarmonics := ParseAntennaMap(fileContents, true)

	if utilities.GetVerbosity(cmd) > 0 {
		fmt.Printf("%# v\n", pretty.Formatter(antennaMapHarmonics))
	}

	fmt.Printf("Number of unique antinode locations accounting for harmonics: %d\n", len(antennaMapHarmonics.Antinodes))

	return nil
}
