package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use: "generate",
	Short: "Generate a new helm chart",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world!!!!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("output", "o", "./", "output directory")
	generateCmd.MarkFlagRequired("output")
}