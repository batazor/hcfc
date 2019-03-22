package cli

import (
	"github.com/batazor/hcfc/pkg/generate"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	config generate.BuildConfig
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new helm chart",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generate.Build(config); err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&config.Filename, "filename", "f", "./deploy.yaml", "config file")
	generateCmd.Flags().StringVarP(&config.Output, "output", "o", "./", "output directory")
	generateCmd.Flags().StringVarP(&config.Template, "template", "t", "./templates", "templates file")
	generateCmd.MarkFlagRequired("output")
}
