package cli

import (
	"fmt"
	"github.com/batazor/hcfc/pkg/generate"
	"github.com/spf13/cobra"
	"os"
)

var (
	filename string
	output   string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new helm chart",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generate.Build(filename, output); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&filename, "filename", "f", "./values.yaml", "config file")
	generateCmd.Flags().StringVarP(&output, "output", "o", "./", "output directory")
	generateCmd.MarkFlagRequired("output")
}
