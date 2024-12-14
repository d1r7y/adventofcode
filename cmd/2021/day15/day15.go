/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyOne_day15

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day15Cmd represents the day15 command
var Day15Cmd = &cobra.Command{
	Use:   "day15",
	Short: `Chiton`,
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

type Point struct {
	X int
	Y int
}

type Sensor struct {
	Position       Point
	Beacon         Point
	BeaconDistance int
}

type SensorList []*Sensor

func (s SensorList) Len() int {
	return len(s)
}

func (s SensorList) Less(i, j int) bool {
	if s[i].Position.X < s[j].Position.X {
		return true
	}
	if s[i].Position.X > s[j].Position.X {
		return false
	}

	// X values are equal.

	if s[i].Position.Y < s[j].Position.Y {
		return true
	}

	return false
}

func (s SensorList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Network struct {
	Min     Point
	Max     Point
	Sensors SensorList
}

func (n *Network) ClosestSensor(p Point) *Sensor {
	closestDistance := math.MaxInt
	var closestSensor *Sensor

	for _, s := range n.Sensors {
		d := ManhattanDistance(s.Position, p)

		if d < closestDistance {
			closestDistance = d
			closestSensor = s
		}
	}

	return closestSensor
}

func (n *Network) SensorIntersection(p Point) SensorList {
	sensors := make(SensorList, 0)

	for _, s := range n.Sensors {
		d := ManhattanDistance(s.Position, p)

		if d <= s.BeaconDistance {
			sensors = append(sensors, s)
		}
	}

	return sensors
}

func (n *Network) InvalidBeaconLocations(row int) []Point {
	// Since sensors can overlap, duplicate points can appear.  Create a map
	// with the points as keys to keep them unique.
	locationMap := make(map[Point]bool)

	for i := n.Min.X; i <= n.Max.X; i++ {
		p := Point{i, row}

		sensors := n.SensorIntersection(p)
		for _, s := range sensors {
			if p != s.Beacon {
				d := ManhattanDistance(p, s.Position)
				if d <= s.BeaconDistance {
					locationMap[p] = true
				}
			}
		}
	}

	locations := make([]Point, 0)

	// Extract the keys.
	for p := range locationMap {
		locations = append(locations, p)
	}

	return locations
}

func (n *Network) PossibleBeaconLocations() []Point {
	locations := make([]Point, 0)

	for i := n.Min.Y; i <= n.Max.Y; i++ {
	NextLocation:
		for j := n.Min.X; j <= n.Max.X; j++ {
			p := Point{j, i}

			for _, s := range n.Sensors {
				d := ManhattanDistance(s.Position, p)

				if d <= s.BeaconDistance {
					j = s.Position.X + s.BeaconDistance - AbsoluteDifference(i, s.Position.Y)
					continue NextLocation
				}
			}

			locations = append(locations, p)
		}
	}

	return locations
}

func AbsoluteDifference(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

func ManhattanDistance(p1, p2 Point) int {
	return AbsoluteDifference(p1.X, p2.X) + AbsoluteDifference(p1.Y, p2.Y)
}

func NewSensor(position Point, beacon Point) *Sensor {
	s := &Sensor{Position: position, Beacon: beacon}
	s.BeaconDistance = ManhattanDistance(s.Position, s.Beacon)
	return s
}

func ParseSensorLine(line string) *Sensor {
	var sensorX, sensorY int
	var beaconX, beaconY int

	count, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
	if err != nil {
		log.Panic(err)
	}

	if count != 4 {
		log.Panic("invalid sensor line")
	}

	return NewSensor(Point{sensorX, sensorY}, Point{beaconX, beaconY})
}

func ParseSensors(fileContents string) SensorList {
	sensors := make(SensorList, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		sensor := ParseSensorLine(line)
		sensors = append(sensors, sensor)
	}

	return sensors
}

func ParseNetwork(fileContents string, tightBounds bool) *Network {
	sensors := ParseSensors(fileContents)
	n := &Network{Sensors: sensors}

	n.Min = Point{math.MaxInt, math.MaxInt}
	n.Max = Point{math.MinInt, math.MinInt}

	for _, sensor := range sensors {
		if tightBounds {
			if sensor.Position.X < n.Min.X {
				n.Min.X = sensor.Position.X
			}
			if sensor.Position.Y < n.Min.Y {
				n.Min.Y = sensor.Position.Y
			}

			if sensor.Position.X > n.Max.X {
				n.Max.X = sensor.Position.X
			}
			if sensor.Position.Y > n.Max.Y {
				n.Max.Y = sensor.Position.Y
			}
		} else {
			if sensor.Position.X-sensor.BeaconDistance < n.Min.X {
				n.Min.X = sensor.Position.X - sensor.BeaconDistance
			}
			if sensor.Position.Y-sensor.BeaconDistance < n.Min.Y {
				n.Min.Y = sensor.Position.Y - sensor.BeaconDistance
			}

			if sensor.Position.X+sensor.BeaconDistance > n.Max.X {
				n.Max.X = sensor.Position.X + sensor.BeaconDistance
			}
			if sensor.Position.Y+sensor.BeaconDistance > n.Max.Y {
				n.Max.Y = sensor.Position.Y + sensor.BeaconDistance
			}
		}
	}

	return n
}

type RiskMap struct {
	Bounds     utilities.Size2D
	Map        [][]int
	Cumulative [][]int
}

type VisitedMap struct {
	Bounds utilities.Size2D
	Map    [][]bool
}

func ParseRiskMap(fileContents string) *RiskMap {
	rm := &RiskMap{}

	rm.Map = make([][]int, 0)
	rm.Cumulative = make([][]int, 0)
	rm.Bounds.Height = 0

	for _, line := range strings.Split(fileContents, "\n") {
		rm.Bounds.Width = 0

		row := make([]int, 0)
		cumulativeRow := make([]int, 0)

		for _, risk := range line {
			row = append(row, int(risk-'0'))
			cumulativeRow = append(cumulativeRow, int(risk-'0'))
			rm.Bounds.Width++
		}

		rm.Map = append(rm.Map, row)
		rm.Cumulative = append(rm.Cumulative, cumulativeRow)
		rm.Bounds.Height++
	}

	return rm
}

func NewVisitedMap(bounds utilities.Size2D) *VisitedMap {
	vm := &VisitedMap{}

	vm.Bounds = bounds
	vm.Map = make([][]bool, bounds.Height)
	for i := 0; i < bounds.Height; i++ {
		vm.Map[i] = make([]bool, bounds.Width)
	}

	return vm
}

func (vm *VisitedMap) SetVisited(position utilities.Point2D) {
	vm.Map[position.Y][position.X] = true
}

func (vm *VisitedMap) GetVisited(position utilities.Point2D) bool {
	return vm.Map[position.Y][position.X]
}

func (vm *VisitedMap) Clone() *VisitedMap {
	newVisitedMap := &VisitedMap{}

	newVisitedMap.Bounds = vm.Bounds
	newVisitedMap.Map = make([][]bool, newVisitedMap.Bounds.Height)

	for i := 0; i < newVisitedMap.Bounds.Height; i++ {
		newVisitedMap.Map[i] = make([]bool, newVisitedMap.Bounds.Width)
		copy(newVisitedMap.Map[i], vm.Map[i])
	}

	return newVisitedMap
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
	Invalid
)

const MaxRisk = 99

func (rm *RiskMap) Risk(position utilities.Point2D) int {
	return rm.Map[position.Y][position.X]
}

func (rm *RiskMap) CumulativeRisk(position utilities.Point2D) int {
	return rm.Cumulative[position.Y][position.X]
}

func (rm *RiskMap) SetCumulativeRisk(position utilities.Point2D, cumulativeRisk int) {
	rm.Cumulative[position.Y][position.X] = cumulativeRisk
}

func (rm *RiskMap) RecurseBestDirection(currentPosition utilities.Point2D) (int, Direction) {

	if currentPosition.X >= rm.Bounds.Width-1 || currentPosition.Y >= rm.Bounds.Height-1 {
		return MaxRisk, Invalid
	}

	downRisk := MaxRisk

	if currentPosition.Y < rm.Bounds.Height-1 {
		// We can go down.
		downRisk = rm.Risk(currentPosition.Down())
	}

	rightRisk := MaxRisk

	if currentPosition.X < rm.Bounds.Width-1 {
		// We can go right.
		rightRisk = rm.Risk(currentPosition.Right())
	}

	if downRisk < rightRisk {
		return downRisk, Down
	} else if rightRisk < downRisk {
		return rightRisk, Right
	} else {
		candidateRightRisk, _ := rm.RecurseBestDirection(currentPosition.Right())
		candidateDownRisk, _ := rm.RecurseBestDirection(currentPosition.Down())
		if candidateRightRisk < candidateDownRisk {
			return candidateRightRisk, Right
		} else {
			return candidateDownRisk, Down
		}
	}
}

func (rm *RiskMap) WalkLeastRiskReversed(currentPosition utilities.Point2D) int {
	visitedMap := NewVisitedMap(rm.Bounds)

	validPosition := func(np utilities.Point2D) bool {
		if np.X < 0 {
			return false
		}
		if np.X >= rm.Bounds.Width {
			return false
		}
		if np.Y < 0 {
			return false
		}
		if np.Y >= rm.Bounds.Height {
			return false
		}

		if visitedMap.GetVisited(np) {
			return false
		}

		return true
	}

	positionsToProcess := &utilities.FIFO[utilities.Point2D]{}
	positionsToProcess.Push(currentPosition)
	visitedMap.SetVisited(currentPosition)

	for {
		if positionsToProcess.IsEmpty() {
			break
		}
		candidatePosition := positionsToProcess.Pop()

		fmt.Printf("[%d,%d]\n", candidatePosition.X, candidatePosition.Y)

		leastRisk := math.MaxInt

		if candidatePosition.X < rm.Bounds.Width-1 {
			rightRisk := rm.CumulativeRisk((candidatePosition.Right()))
			if rightRisk < leastRisk {
				leastRisk = rightRisk
			}
		}

		if candidatePosition.Y < rm.Bounds.Height-1 {
			downRisk := rm.CumulativeRisk((candidatePosition.Down()))
			if downRisk < leastRisk {
				leastRisk = downRisk
			}
		}

		if leastRisk < math.MaxInt {
			if candidatePosition.X == 0 && candidatePosition.Y == 0 {
				rm.SetCumulativeRisk(candidatePosition, leastRisk)
			} else {
				rm.SetCumulativeRisk(candidatePosition, leastRisk+rm.Risk(candidatePosition))
			}
		}

		if validPosition(candidatePosition.Left()) {
			positionsToProcess.Push(candidatePosition.Left())
			visitedMap.SetVisited(candidatePosition.Left())
		}
		if validPosition(candidatePosition.Up()) {
			positionsToProcess.Push(candidatePosition.Up())
			visitedMap.SetVisited(candidatePosition.Up())
		}
	}

	return rm.CumulativeRisk(utilities.NewPoint2D(0, 0))
}

func day(fileContents string) error {
	// Part 1: Your goal is to find a path with the lowest total risk
	rm := ParseRiskMap(fileContents)

	totalRisk := rm.WalkLeastRiskReversed(utilities.NewPoint2D(rm.Bounds.Width-1, rm.Bounds.Height-1))

	fmt.Printf("Least amount of risk: %d.\n", totalRisk)

	return nil
}
