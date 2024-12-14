/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyTwo

import (
	TwentyTwentyTwo_day01 "github.com/d1r7y/adventofcode/cmd/2022/day01"
	TwentyTwentyTwo_day02 "github.com/d1r7y/adventofcode/cmd/2022/day02"
	TwentyTwentyTwo_day03 "github.com/d1r7y/adventofcode/cmd/2022/day03"
	TwentyTwentyTwo_day04 "github.com/d1r7y/adventofcode/cmd/2022/day04"
	TwentyTwentyTwo_day05 "github.com/d1r7y/adventofcode/cmd/2022/day05"
	TwentyTwentyTwo_day06 "github.com/d1r7y/adventofcode/cmd/2022/day06"
	TwentyTwentyTwo_day07 "github.com/d1r7y/adventofcode/cmd/2022/day07"
	TwentyTwentyTwo_day08 "github.com/d1r7y/adventofcode/cmd/2022/day08"
	TwentyTwentyTwo_day09 "github.com/d1r7y/adventofcode/cmd/2022/day09"
	TwentyTwentyTwo_day10 "github.com/d1r7y/adventofcode/cmd/2022/day10"
	TwentyTwentyTwo_day11 "github.com/d1r7y/adventofcode/cmd/2022/day11"
	TwentyTwentyTwo_day12 "github.com/d1r7y/adventofcode/cmd/2022/day12"
	TwentyTwentyTwo_day13 "github.com/d1r7y/adventofcode/cmd/2022/day13"
	TwentyTwentyTwo_day14 "github.com/d1r7y/adventofcode/cmd/2022/day14"
	TwentyTwentyTwo_day15 "github.com/d1r7y/adventofcode/cmd/2022/day15"
	TwentyTwentyTwo_day16 "github.com/d1r7y/adventofcode/cmd/2022/day16"
	TwentyTwentyTwo_day17 "github.com/d1r7y/adventofcode/cmd/2022/day17"
	TwentyTwentyTwo_day18 "github.com/d1r7y/adventofcode/cmd/2022/day18"
	TwentyTwentyTwo_day20 "github.com/d1r7y/adventofcode/cmd/2022/day20"
	TwentyTwentyTwo_day21 "github.com/d1r7y/adventofcode/cmd/2022/day21"
	"github.com/spf13/cobra"
)

// TwentyTwentyTwo represents the base commands for 2022
var TwentyTwentyTwoCmd = &cobra.Command{
	Use:   "2022",
	Short: "2022 solutions for Advent Of Code (www.adventofcode.com)",
	Long:  ``,
}

func init() {
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day01.Day01Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day02.Day02Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day03.Day03Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day04.Day04Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day05.Day05Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day06.Day06Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day07.Day07Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day08.Day08Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day09.Day09Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day10.Day10Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day11.Day11Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day12.Day12Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day13.Day13Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day14.Day14Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day15.Day15Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day16.Day16Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day17.Day17Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day18.Day18Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day20.Day20Cmd)
	TwentyTwentyTwoCmd.AddCommand(TwentyTwentyTwo_day21.Day21Cmd)
}
