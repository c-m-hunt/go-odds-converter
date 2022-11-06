package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

const (
	US = "us"
	FRACTION = "fraction"
	DECIMAL = "decimal"
)

const (
	SLASH = "/"
	HYPHEN = "-"
)

var (
	rootCmd = &cobra.Command{
		Use: "odds",
		Short: `A CLI to calculate different odds methods`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			oddStr := args[0]
			odds := NewOdds(oddStr)
			odds.Display()
		},
	}
)

func detectOddsType(odd string) string {
	if strings.Contains(odd, SLASH) {
		return FRACTION
	} else if strings.Contains(odd, HYPHEN) {
		return US
	} else {
		return DECIMAL
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}