/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day12

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day12Cmd represents the day12 command
var Day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: `Garden Groups`,
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

type PlantType rune
type Row []PlantType
type Region struct {
	Plant PlantType
	Plots *utilities.SetPoint2D
}

type Map struct {
	Bounds  utilities.Size2D
	Plants  map[PlantType][]utilities.Point2D
	Regions map[int]*Region
	Columns []Row
}

func (m *Map) GetPlant(location utilities.Point2D) PlantType {
	return m.Columns[location.Y][location.X]
}

func (m *Map) NumRegions() int {
	return len(m.Regions)
}

func (m *Map) RegionArea(regionID int) int {
	if region, ok := m.Regions[regionID]; ok {
		return region.Plots.Size()
	}

	return 0
}

func (m *Map) RegionFencingPrice(regionID int) int {
	return m.RegionArea(regionID) * m.RegionPerimeter(regionID)
}

func (m *Map) RegionFencingPriceBulkDiscount(regionID int) int {
	return m.RegionArea(regionID) * m.RegionPerimeterSides(regionID)
}

func (m *Map) RegionPerimeter(regionID int) int {
	perimeter := 0

	if region, ok := m.Regions[regionID]; ok {
		for p := range region.Plots.All() {
			if p.Y == 0 || m.GetPlant(p.Up()) != region.Plant {
				perimeter++
			}

			if p.Y == m.Bounds.Height-1 || m.GetPlant(p.Down()) != region.Plant {
				perimeter++
			}

			if p.X == 0 || m.GetPlant(p.Left()) != region.Plant {
				perimeter++
			}

			if p.X == m.Bounds.Width-1 || m.GetPlant(p.Right()) != region.Plant {
				perimeter++
			}
		}
	}

	return perimeter
}

func (m *Map) RegionPerimeterSides(regionID int) int {
	perimeter := 0

	if region, ok := m.Regions[regionID]; ok {
		const (
			BottomLeft   = 0x01
			TopLeft      = 0b11
			TopRight     = 0b10
			BottomRight  = 0b00
			PositionMask = 0b11
		)

		sharedCorners := make(map[utilities.Point2D][]int)

		// Add all the plot edges of this region.
		for p := range region.Plots.All() {
			if cornerList, ok := sharedCorners[p]; ok {
				sharedCorners[p] = append(cornerList, TopLeft)
			} else {
				sharedCorners[p] = []int{TopLeft}
			}

			if cornerList, ok := sharedCorners[p.Right()]; ok {
				sharedCorners[p.Right()] = append(cornerList, TopRight)
			} else {
				sharedCorners[p.Right()] = []int{TopRight}
			}

			if cornerList, ok := sharedCorners[p.Down()]; ok {
				sharedCorners[p.Down()] = append(cornerList, BottomLeft)
			} else {
				sharedCorners[p.Down()] = []int{BottomLeft}
			}

			if cornerList, ok := sharedCorners[p.DownRight()]; ok {
				sharedCorners[p.DownRight()] = append(cornerList, BottomRight)
			} else {
				sharedCorners[p.DownRight()] = []int{BottomRight}
			}
		}

		for _, countList := range sharedCorners {
			if len(countList) == 1 {
				perimeter++
			} else if len(countList) == 2 {
				// The only case we care about with 2 shared corners is if the plots are
				// diagonal to each other: two corners are shared, but they aren't neighbors.
				// So we count them as two perimeters.
				if (countList[0]^countList[1])&PositionMask == PositionMask {
					perimeter += 2
				}
			} else if len(countList) == 3 {
				perimeter++
			}
		}
	}

	return perimeter
}

func ParseMap(fileContents string) *Map {
	gardenMap := &Map{}

	gardenMap.Columns = make([]Row, 0)
	gardenMap.Plants = make(map[PlantType][]utilities.Point2D)
	gardenMap.Regions = make(map[int]*Region)

	for y, line := range strings.Split(fileContents, "\n") {
		gardenMap.Bounds.Width = 0
		row := make(Row, 0)

		for x, c := range line {
			currentLocation := utilities.NewPoint2D(x, y)

			row = append(row, PlantType(c))
			gardenMap.Plants[PlantType(c)] = append(gardenMap.Plants[PlantType(c)], currentLocation)

			gardenMap.Bounds.Width++
		}

		gardenMap.Columns = append(gardenMap.Columns, row)
		gardenMap.Bounds.Height++
	}

	unclaimedPlots := gardenMap.Bounds.Height * gardenMap.Bounds.Width
	claimedLocations := utilities.NewSetPoint2D()

	regionID := 0

	for unclaimedPlots > 0 {
		for y := 0; y < gardenMap.Bounds.Height; y++ {
			for x := 0; x < gardenMap.Bounds.Width; x++ {
				currentPosition := utilities.NewPoint2D(x, y)

				if claimedLocations.Exists(currentPosition) {
					// Already claimed.
					continue
				}

				visitedLocations := utilities.NewSetPoint2D()

				claimedLocations.Add(currentPosition)
				unclaimedPlots--
				visitedLocations.Add(currentPosition)

				plant := gardenMap.GetPlant(currentPosition)

				regionLocations := utilities.NewSetPoint2D()
				regionLocations.Add(currentPosition)

				var claimLocations func(m *Map, regionLocations *utilities.SetPoint2D, plant PlantType, cp utilities.Point2D) int

				checkLocation := func(m *Map, regionLocations *utilities.SetPoint2D, plant PlantType, pos utilities.Point2D) int {
					claims := 0
					if !visitedLocations.Exists(pos) {
						visitedLocations.Add(pos)
						if m.GetPlant(pos) == plant {
							claims++
							claimedLocations.Add(pos)
							regionLocations.Add(pos)
							claims += claimLocations(m, regionLocations, plant, pos)
						}
					}
					return claims
				}

				claimLocations = func(m *Map, regionLocations *utilities.SetPoint2D, plant PlantType, cp utilities.Point2D) int {
					claims := 0
					if cp.Y > 0 {
						np := cp.Up()
						claims += checkLocation(m, regionLocations, plant, np)
					}
					if cp.Y < gardenMap.Bounds.Height-1 {
						np := cp.Down()
						claims += checkLocation(m, regionLocations, plant, np)
					}
					if cp.X > 0 {
						np := cp.Left()
						claims += checkLocation(m, regionLocations, plant, np)
					}
					if cp.X < gardenMap.Bounds.Width-1 {
						np := cp.Right()
						claims += checkLocation(m, regionLocations, plant, np)
					}

					return claims
				}

				claims := claimLocations(gardenMap, regionLocations, plant, currentPosition)
				unclaimedPlots -= claims

				region := &Region{
					Plant: plant,
					Plots: regionLocations,
				}

				gardenMap.Regions[regionID] = region

				regionID++
			}
		}
	}

	return gardenMap
}

func day(fileContents string) error {
	// Part 1: You're about to settle near a complex arrangement of garden plots when some
	// Elves ask if you can lend a hand. They'd like to set up fences around each region of
	// garden plots, but they can't figure out how much fence they need to order or how much
	// it will cost. They hand you a map (your puzzle input) of the garden plots.
	//
	// Each garden plot grows only a single type of plant and is indicated by a single letter
	// on your map. When multiple garden plots are growing the same type of plant and are
	// touching (horizontally or vertically), they form a region.
	//
	// In order to accurately calculate the cost of the fence around a single region, you
	// need to know that region's area and perimeter.
	//
	// Due to "modern" business practices, the price of fence required for a region is found
	// by multiplying that region's area by its perimeter. The total price of fencing all
	// regions on a map is found by adding together the price of fence for every region on the map.
	//
	// What is the total price of fencing all regions on your map?

	gardenMap := ParseMap(fileContents)

	totalFencingPrice := 0

	for i := 0; i < gardenMap.NumRegions(); i++ {
		totalFencingPrice += gardenMap.RegionFencingPrice(i)
	}

	fmt.Printf("Total cost of fencing: %d\n", totalFencingPrice)

	// Part 2: Fortunately, the Elves are trying to order so much fence that they qualify for a
	// bulk discount!
	//
	// Under the bulk discount, instead of using the perimeter to calculate the price, you need to
	// use the number of sides each region has. Each straight section of fence counts as a side,
	// regardless of how long it is.

	totalFencingPriceBulkDiscount := 0

	for i := 0; i < gardenMap.NumRegions(); i++ {
		totalFencingPriceBulkDiscount += gardenMap.RegionFencingPriceBulkDiscount(i)
	}

	fmt.Printf("Total cost of fencing with bulk discount: %d\n", totalFencingPriceBulkDiscount)

	return nil
}
