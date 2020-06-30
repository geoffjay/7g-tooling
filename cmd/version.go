package cmd

import (
	"fmt"

	"github.com/geoffjay/7g-cli/pkg"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of 7g-cli",
	Long:  `7Geese CLI version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.VERSION)
	},
}