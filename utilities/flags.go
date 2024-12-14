/*
Copyright © 2021-2024 Cameron Esfahani
*/

package utilities

import "github.com/spf13/cobra"

func GetInputPath(cmd *cobra.Command) string {
	if cmd == nil {
		return ""
	}

	inputPath, err := cmd.Flags().GetString("input")

	if err != nil {
		return ""
	}

	return inputPath
}

func GetVerbosity(cmd *cobra.Command) int {
	if cmd == nil {
		return 0
	}

	verbosity, err := cmd.Flags().GetCount("verbose")

	if err != nil {
		return 0
	}

	return verbosity
}
