/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day06

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day06Cmd represents the day06 command
var Day06Cmd = &cobra.Command{
	Use:   "day06",
	Short: `Guard Gallivant`,
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

type Cell int

const (
	Empty Cell = iota
	Obstruction
)

type Row []Cell

type VisitedRow []int

type Direction int

const (
	North Direction = 1
	East  Direction = 2
	South Direction = 4
	West  Direction = 8
)

type Map struct {
	Bounds       utilities.Size2D
	Position     utilities.Point2D
	Facing       Direction
	Columns      []Row
	VisitedCells int
	Visited      []VisitedRow
	Looping      bool
}

func (m *Map) GetCell(location utilities.Point2D) Cell {
	return m.Columns[location.Y][location.X]
}

func (m *Map) GetVisited(location utilities.Point2D) bool {
	return m.Visited[location.Y][location.X] != 0
}

func (m *Map) SetVisited(location utilities.Point2D, facing Direction) {
	if !m.GetVisited(location) {
		m.VisitedCells++
	}

	v := m.Visited[location.Y][location.X]

	if (v & int(facing)) != 0 {
		m.Looping = true
	}

	m.Visited[location.Y][location.X] = v | int(facing)
}

func (m *Map) AddObstruction(location utilities.Point2D) {
	m.Columns[location.Y][location.X] = Obstruction
}

func (m *Map) AreLooping() bool {
	return m.Looping
}

func (m *Map) Walk() bool {
	switch m.Facing {
	case North:
		if m.Position.Y == 0 {
			return true
		}

		c := m.GetCell(m.Position.Up())
		if c == Obstruction {
			m.Facing = East
		} else {
			m.Position = m.Position.Up()
			m.SetVisited(m.Position, m.Facing)
		}
	case East:
		if m.Position.X == m.Bounds.Width-1 {
			return true
		}

		c := m.GetCell(m.Position.Right())
		if c == Obstruction {
			m.Facing = South
		} else {
			m.Position = m.Position.Right()
			m.SetVisited(m.Position, m.Facing)
		}
	case South:
		if m.Position.Y == m.Bounds.Height-1 {
			return true
		}

		c := m.GetCell(m.Position.Down())
		if c == Obstruction {
			m.Facing = West
		} else {
			m.Position = m.Position.Down()
			m.SetVisited(m.Position, m.Facing)
		}
	case West:
		if m.Position.X == 0 {
			return true
		}

		c := m.GetCell(m.Position.Left())
		if c == Obstruction {
			m.Facing = North
		} else {
			m.Position = m.Position.Left()
			m.SetVisited(m.Position, m.Facing)
		}
	}

	return false
}

func ParseMap(fileContents string) *Map {
	m := &Map{}
	m.Columns = make([]Row, 0)

	m.Looping = false
	m.VisitedCells = 0

	for y, line := range strings.Split(fileContents, "\n") {
		row := make(Row, 0)
		visitedRow := make(VisitedRow, 0)
		for x, r := range line {
			if r == '.' {
				row = append(row, Empty)
			} else if r == '#' {
				row = append(row, Obstruction)
			} else if r == '^' {
				m.Facing = North
				m.Position = utilities.NewPoint2D(x, y)
				row = append(row, Empty)
			}
			visitedRow = append(visitedRow, 0)
		}

		m.Columns = append(m.Columns, row)
		m.Visited = append(m.Visited, visitedRow)
	}

	m.Bounds.Height = len(m.Columns)
	m.Bounds.Width = len(m.Columns[0])

	m.SetVisited(m.Position, m.Facing)

	return m
}

func day(cmd *cobra.Command, fileContents string) error {
	// Part 1: The map shows the current position of the guard with ^ (to indicate
	// the guard is currently facing up from the perspective of the map). Any
	// obstructions - crates, desks, alchemical reactors, etc. - are shown as #.
	//
	// Lab guards in 1518 follow a very strict patrol protocol which involves repeatedly
	// following these steps:
	//
	//	If there is something directly in front of you, turn right 90 degrees.
	//	Otherwise, take a step forward.
	//
	// By predicting the guard's route, you can determine which specific positions in the
	// lab will be in the patrol path. Including the guard's starting position, the positions
	// visited by the guard before leaving the area are marked with an X.
	//
	// Predict the path of the guard. How many distinct positions will the guard visit
	// before leaving the mapped area?

	roomMap := ParseMap(fileContents)

	guardStartingLocation := roomMap.Position

	for {
		if roomMap.Walk() {
			break
		}
	}

	fmt.Printf("Total distinct positions visited by guard: %d\n", roomMap.VisitedCells)

	// Part 2: Returning after what seems like only a few seconds to The Historians, they
	// explain that the guard's patrol area is simply too large for them to safely search
	// the lab without getting caught.
	//
	// Fortunately, they are pretty sure that adding a single new obstruction won't cause a
	// time paradox. They'd like to place the new obstruction in such a way that the guard
	// will get stuck in a loop, making the rest of the lab safe to search.
	//
	// To have the lowest chance of creating a time paradox, The Historians would like to know
	// all of the possible positions for such an obstruction. The new obstruction can't be placed
	// at the guard's starting position - the guard is there right now and would notice.
	//
	// You need to get the guard stuck in a loop by adding a single new obstruction. How many
	// different positions could you choose for this obstruction?

	loopingObstructionCount := 0

	for y := 0; y < roomMap.Bounds.Height; y++ {
		for x := 0; x < roomMap.Bounds.Width; x++ {
			obstructionLocation := utilities.NewPoint2D(x, y)

			if guardStartingLocation == obstructionLocation {
				continue
			}

			obstructedRoomMap := ParseMap(fileContents)
			obstructedRoomMap.AddObstruction(obstructionLocation)

			for {
				if obstructedRoomMap.Walk() {
					break
				}

				if obstructedRoomMap.AreLooping() {
					if utilities.GetVerbosity(cmd) > 0 {
						fmt.Printf("Obstruction @ %dx%d loops guard\n", x, y)
					}
					loopingObstructionCount++
					break
				}
			}
		}
	}

	fmt.Printf("Number of different positions to place obstruction to loop guard: %d\n", loopingObstructionCount)

	return nil
}
