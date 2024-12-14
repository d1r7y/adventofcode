/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyThree

import (
	TwentyTwentyThree_day01 "github.com/d1r7y/adventofcode/cmd/2023/day01"
	TwentyTwentyThree_day02 "github.com/d1r7y/adventofcode/cmd/2023/day02"
	TwentyTwentyThree_day03 "github.com/d1r7y/adventofcode/cmd/2023/day03"
	TwentyTwentyThree_day04 "github.com/d1r7y/adventofcode/cmd/2023/day04"
	TwentyTwentyThree_day05 "github.com/d1r7y/adventofcode/cmd/2023/day05"
	TwentyTwentyThree_day06 "github.com/d1r7y/adventofcode/cmd/2023/day06"
	TwentyTwentyThree_day07 "github.com/d1r7y/adventofcode/cmd/2023/day07"
	TwentyTwentyThree_day08 "github.com/d1r7y/adventofcode/cmd/2023/day08"
	TwentyTwentyThree_day09 "github.com/d1r7y/adventofcode/cmd/2023/day09"
	TwentyTwentyThree_day10 "github.com/d1r7y/adventofcode/cmd/2023/day10"
	TwentyTwentyThree_day11 "github.com/d1r7y/adventofcode/cmd/2023/day11"
	TwentyTwentyThree_day12 "github.com/d1r7y/adventofcode/cmd/2023/day12"
	TwentyTwentyThree_day13 "github.com/d1r7y/adventofcode/cmd/2023/day13"
	TwentyTwentyThree_day14 "github.com/d1r7y/adventofcode/cmd/2023/day14"
	TwentyTwentyThree_day15 "github.com/d1r7y/adventofcode/cmd/2023/day15"
	TwentyTwentyThree_day16 "github.com/d1r7y/adventofcode/cmd/2023/day16"
	TwentyTwentyThree_day17 "github.com/d1r7y/adventofcode/cmd/2023/day17"
	TwentyTwentyThree_day18 "github.com/d1r7y/adventofcode/cmd/2023/day18"
	TwentyTwentyThree_day20 "github.com/d1r7y/adventofcode/cmd/2023/day20"
	TwentyTwentyThree_day21 "github.com/d1r7y/adventofcode/cmd/2023/day21"
	"github.com/spf13/cobra"
)

// TwentyTwentyThree represents the base commands for 2023
var TwentyTwentyThreeCmd = &cobra.Command{
	Use:   "2023",
	Short: "2023 solutions for Advent Of Code (www.adventofcode.com)",
	Long:  ``,
}

func init() {
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day01.Day01Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day02.Day02Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day03.Day03Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day04.Day04Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day05.Day05Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day06.Day06Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day07.Day07Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day08.Day08Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day09.Day09Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day10.Day10Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day11.Day11Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day12.Day12Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day13.Day13Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day14.Day14Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day15.Day15Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day16.Day16Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day17.Day17Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day18.Day18Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day20.Day20Cmd)
	TwentyTwentyThreeCmd.AddCommand(TwentyTwentyThree_day21.Day21Cmd)
}
