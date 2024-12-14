/*
Copyright © 2021-2024 Cameron Esfahani
*/

package TwentyTwentyOne

import (
	TwentyTwentyOne_day15 "github.com/d1r7y/adventofcode/cmd/2021/day15"
	"github.com/spf13/cobra"
)

// TwentyTwentyOne represents the base commands for 2021
var TwentyTwentyOneCmd = &cobra.Command{
	Use:   "2021",
	Short: "2021 solutions for Advent Of Code (www.adventofcode.com)",
	Long:  ``,
}

func init() {
	TwentyTwentyOneCmd.AddCommand(TwentyTwentyOne_day15.Day15Cmd)
}
