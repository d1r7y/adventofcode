/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour_day10

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/d1r7y/adventofcode/utilities"
	"github.com/spf13/cobra"
)

// Day10Cmd represents the day10 command
var Day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: `Hoof It`,
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

		if fileContent != nil {
			err = day(string(fileContent))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

const (
	TrailPeakHeight = 9
)

type Scores map[utilities.Point2D]int

type Ratings map[utilities.Point2D]int

type Row []int

type TopoMap struct {
	Bounds     utilities.Size2D
	Trailheads []utilities.Point2D
	Columns    []Row
}

func (t *TopoMap) GetHeight(location utilities.Point2D) int {
	return t.Columns[location.Y][location.X]
}

func (t *TopoMap) HikeScores() Scores {
	scores := make(Scores)

	for _, th := range t.Trailheads {
		positions := utilities.NewFIFO[utilities.Point2D]()
		visitedPositions := make(map[utilities.Point2D]bool)

		neighborCheck := func(ch int, np utilities.Point2D) int {
			s := 0
			if _, ok := visitedPositions[np]; !ok {

				neighborHeight := t.GetHeight(np)
				if neighborHeight == ch+1 {
					if neighborHeight == TrailPeakHeight {
						s++
					} else {
						positions.Push(np)
					}
					visitedPositions[np] = true
				}
			}

			return s
		}

		score := 0

		positions.Push(th)

		for !positions.IsEmpty() {
			position := positions.Pop()

			currentHeight := t.GetHeight(position)
			if position.Y > 0 {
				neighborPosition := position.Up()

				score += neighborCheck(currentHeight, neighborPosition)
			}

			if position.Y < t.Bounds.Height-1 {
				neighborPosition := position.Down()

				score += neighborCheck(currentHeight, neighborPosition)
			}

			if position.X > 0 {
				neighborPosition := position.Left()

				score += neighborCheck(currentHeight, neighborPosition)
			}

			if position.X < t.Bounds.Width-1 {
				neighborPosition := position.Right()

				score += neighborCheck(currentHeight, neighborPosition)
			}

			visitedPositions[position] = true
		}

		scores[th] = score
	}

	return scores
}

func (t *TopoMap) HikeRatings() Ratings {
	ratings := make(Ratings)

	for _, th := range t.Trailheads {

		var neighborCheck func(ch int, pos utilities.Point2D) int
		var recursive func(position utilities.Point2D) int

		neighborCheck = func(ch int, pos utilities.Point2D) int {
			neighborHeight := t.GetHeight(pos)
			if neighborHeight == ch+1 {
				if neighborHeight == TrailPeakHeight {
					return 1
				} else {
					return recursive(pos)
				}
			}

			return 0
		}

		recursive = func(position utilities.Point2D) int {
			currentRating := 0
			currentHeight := t.GetHeight(position)

			if position.Y > 0 {
				neighborPosition := position.Up()

				currentRating += neighborCheck(currentHeight, neighborPosition)
			}

			if position.Y < t.Bounds.Height-1 {
				neighborPosition := position.Down()

				currentRating += neighborCheck(currentHeight, neighborPosition)
			}

			if position.X > 0 {
				neighborPosition := position.Left()

				currentRating += neighborCheck(currentHeight, neighborPosition)
			}

			if position.X < t.Bounds.Width-1 {
				neighborPosition := position.Right()

				currentRating += neighborCheck(currentHeight, neighborPosition)
			}

			return currentRating
		}

		rating := recursive(th)

		ratings[th] = rating
	}

	return ratings
}

func ParseTopoMap(fileContents string) *TopoMap {
	topoMap := &TopoMap{}
	topoMap.Trailheads = make([]utilities.Point2D, 0)
	topoMap.Columns = make([]Row, 0)

	for y, line := range strings.Split(fileContents, "\n") {
		topoMap.Bounds.Width = 0
		row := make(Row, 0)

		for x, c := range line {
			currentLocation := utilities.NewPoint2D(x, y)

			var height int

			if c == '.' {
				height = TrailPeakHeight + 1 // Impeneterable
			} else {
				height = int(c) - '0'
			}

			if height == 0 {
				topoMap.Trailheads = append(topoMap.Trailheads, currentLocation)
			}

			row = append(row, height)

			topoMap.Bounds.Width++
		}

		topoMap.Columns = append(topoMap.Columns, row)
		topoMap.Bounds.Height++
	}

	return topoMap
}

func day(fileContents string) error {
	// Part 1: The topographic map indicates the height at each position using a scale from
	// 0 (lowest) to 9 (highest).
	//
	// Based on un-scorched scraps of the book, you determine that a good hiking trail is as long
	// as possible and has an even, gradual, uphill slope. For all practical purposes, this means
	// that a hiking trail is any path that starts at height 0, ends at height 9, and always
	// increases by a height of exactly 1 at each step. Hiking trails never include diagonal
	// steps - only up, down, left, or right (from the perspective of the map).
	//
	// You look up from the map and notice that the reindeer has helpfully begun to construct a small
	// pile of pencils, markers, rulers, compasses, stickers, and other equipment you might need to
	// update the map with hiking trails.
	//
	// A trailhead is any position that starts one or more hiking trails - here, these positions will
	// always have height 0. Assembling more fragments of pages, you establish that a trailhead's score
	// is the number of 9-height positions reachable from that trailhead via a hiking trail.
	//
	// What is the sum of the scores of all trailheads on your topographic map?

	topoMap := ParseTopoMap(fileContents)

	trailHeadScores := topoMap.HikeScores()

	totalTrailheadScores := 0

	for _, score := range trailHeadScores {
		totalTrailheadScores += score
	}

	fmt.Printf("Sum of scores of all trailheads: %d\n", totalTrailheadScores)

	// Part 2: The reindeer spends a few minutes reviewing your hiking trail map before realizing something,
	// disappearing for a few minutes, and finally returning with yet another slightly-charred piece of paper.
	//
	// The paper describes a second way to measure a trailhead called its rating. A trailhead's rating is the
	// number of distinct hiking trails which begin at that trailhead.
	//
	// You're not sure how, but the reindeer seems to have crafted some tiny flags out of toothpicks and bits of
	// paper and is using them to mark trailheads on your topographic map. What is the sum of the ratings of all
	// trailheads?

	trailHeadRatings := topoMap.HikeRatings()

	totalTrailheadRatings := 0

	for _, rating := range trailHeadRatings {
		totalTrailheadRatings += rating
	}

	fmt.Printf("Sum of ratings of all trailheads: %d\n", totalTrailheadRatings)

	return nil
}
