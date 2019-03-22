package cli

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hcfc",
	Short: "Generate a new helm chart",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Try: hcfc --help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
