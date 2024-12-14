/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyThree_day06

import (
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
	Short: `Wait For It`,
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

type Race struct {
	Time     int
	Distance int
}

func SpeedAtTime(pressTime int, raceTime int) int {
	if pressTime < raceTime {
		return pressTime
	} else {
		return 0
	}
}

func DistanceAtTime(speed int, time int) int {
	return speed * time
}

func DoesBeatRecord(pressTime int, raceTime int, recordDistance int) bool {
	return DistanceAtTime(SpeedAtTime(pressTime, raceTime), raceTime-pressTime) > recordDistance
}

func ParseRaces(fileContents string, part1 bool) []Race {
	races := make([]Race, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		switch {
		case strings.HasPrefix(strings.TrimSpace(line), "Time: "):
			var timeList []int
			if part1 {
				timeList = utilities.ParseIntList(strings.TrimPrefix(line, "Time: "))
			} else {
				timeList = utilities.ParseIntListRemovingAllWhitespace(strings.TrimPrefix(line, "Time: "))
			}

			for i, t := range timeList {
				if i >= len(races) {
					races = append(races, Race{})
				}

				races[i].Time = t
			}
		case strings.HasPrefix(strings.TrimSpace(line), "Distance: "):
			var distanceList []int
			if part1 {
				distanceList = utilities.ParseIntList(strings.TrimPrefix(line, "Distance: "))
			} else {
				distanceList = utilities.ParseIntListRemovingAllWhitespace(strings.TrimPrefix(line, "Distance: "))
			}

			for i, d := range distanceList {
				if i >= len(races) {
					races = append(races, Race{})
				}

				races[i].Distance = d
			}
		default:
			// Range line.
			log.Panicf("unexpected line '%s'\n", line)
		}
	}

	return races
}

func day(fileContents string) error {
	races1 := ParseRaces(fileContents, true)

	// Part 1: Determine the number of ways you could beat the record in each race.
	// What do you get if you multiply these numbers together?
	totalWinningWays := 1

	for _, r := range races1 {
		beatRecordCount := 0
		for pt := 0; pt < r.Time; pt++ {
			if DoesBeatRecord(pt, r.Time, r.Distance) {
				beatRecordCount++
			}
		}

		totalWinningWays *= beatRecordCount
	}

	log.Printf("Total winning ways: %d\n", totalWinningWays)

	races2 := ParseRaces(fileContents, false)

	// Part 1: How many ways can you beat the record in this one much longer race?
	totalWinningWays = 1

	for _, r := range races2 {
		beatRecordCount := 0
		for pt := 0; pt < r.Time; pt++ {
			if DoesBeatRecord(pt, r.Time, r.Distance) {
				beatRecordCount++
			}
		}

		totalWinningWays *= beatRecordCount
	}

	log.Printf("Total winning ways: %d\n", totalWinningWays)

	return nil
}
