package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of hcfc",
	Long: `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hcfc v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}