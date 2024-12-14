/*
Copyright © 2021-2024 Cameron Esfahani
*/

package cmd

import (
	"os"

	TwentyTwentyOne "github.com/d1r7y/adventofcode/cmd/2021"
	TwentyTwentyTwo "github.com/d1r7y/adventofcode/cmd/2022"
	TwentyTwentyThree "github.com/d1r7y/adventofcode/cmd/2023"
	TwentyTwentyFour "github.com/d1r7y/adventofcode/cmd/2024"
	"github.com/spf13/cobra"
)

var verbosity int = 0
var inputPath string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "advent",
	Short: "Solutions for Advent Of Code (www.adventofcode.com)",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "verbose output")
	RootCmd.PersistentFlags().StringVarP(&inputPath, "input", "i", "", "Input file")

	RootCmd.AddCommand(TwentyTwentyOne.TwentyTwentyOneCmd)
	RootCmd.AddCommand(TwentyTwentyTwo.TwentyTwentyTwoCmd)
	RootCmd.AddCommand(TwentyTwentyThree.TwentyTwentyThreeCmd)
	RootCmd.AddCommand(TwentyTwentyFour.TwentyTwentyFourCmd)
}
