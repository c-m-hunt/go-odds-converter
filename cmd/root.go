package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	US       = "us"
	FRACTION = "fraction"
	DECIMAL  = "decimal"
)

const (
	SLASH  = "/"
	HYPHEN = "-"
)

var (
	rootCmd = &cobra.Command{
		Use:   "odds",
		Short: `A CLI to calculate different odds methods`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			oddStr := args[0]
			odds, err := NewOdds(oddStr)
			if err != nil {
				log.Error(err)
				os.Exit(0)
			}
			odds.Display(os.Stdout)
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
