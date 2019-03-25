package cli

import (
	"github.com/spf13/cobra"
	"log"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of hcfc",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("hcfc v1.1.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
