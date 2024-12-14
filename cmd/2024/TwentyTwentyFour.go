/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyFour

import (
	TwentyTwentyFour_day01 "github.com/d1r7y/adventofcode/cmd/2024/day01"
	TwentyTwentyFour_day02 "github.com/d1r7y/adventofcode/cmd/2024/day02"
	TwentyTwentyFour_day03 "github.com/d1r7y/adventofcode/cmd/2024/day03"
	TwentyTwentyFour_day04 "github.com/d1r7y/adventofcode/cmd/2024/day04"
	TwentyTwentyFour_day05 "github.com/d1r7y/adventofcode/cmd/2024/day05"
	TwentyTwentyFour_day06 "github.com/d1r7y/adventofcode/cmd/2024/day06"
	TwentyTwentyFour_day07 "github.com/d1r7y/adventofcode/cmd/2024/day07"
	TwentyTwentyFour_day08 "github.com/d1r7y/adventofcode/cmd/2024/day08"
	TwentyTwentyFour_day09 "github.com/d1r7y/adventofcode/cmd/2024/day09"
	TwentyTwentyFour_day10 "github.com/d1r7y/adventofcode/cmd/2024/day10"
	TwentyTwentyFour_day11 "github.com/d1r7y/adventofcode/cmd/2024/day11"
	TwentyTwentyFour_day12 "github.com/d1r7y/adventofcode/cmd/2024/day12"
	"github.com/spf13/cobra"
)

// TwentyTwentyFour represents the base commands for 2024
var TwentyTwentyFourCmd = &cobra.Command{
	Use:   "2024",
	Short: "2024 solutions for Advent Of Code (www.adventofcode.com)",
	Long:  ``,
}

func init() {
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day01.Day01Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day02.Day02Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day03.Day03Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day04.Day04Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day05.Day05Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day06.Day06Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day07.Day07Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day08.Day08Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day09.Day09Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day10.Day10Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day11.Day11Cmd)
	TwentyTwentyFourCmd.AddCommand(TwentyTwentyFour_day12.Day12Cmd)
}
