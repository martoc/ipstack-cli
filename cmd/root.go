package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(getCoordinates)
}

var rootCmd = &cobra.Command{
	Use:   "ipstack-cli",
	Short: "Interacts with the IPStack API",
	Long:  `Interacts with the IPStack API`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
