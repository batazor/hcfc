package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hcfc",
	Short: "Generate a new helm chart",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Try: hcfc --help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
