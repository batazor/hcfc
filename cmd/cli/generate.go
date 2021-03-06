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
	Long: `
As example:
hcfc generate -o ./ops/Helm/mychart -f deploy.yaml
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := generate.Build(config); err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	// Init CLI command
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&config.Filename, "filename", "f", "./deploy.yaml", "config file")
	generateCmd.Flags().StringVarP(&config.Template.Output, "output", "o", "./", "output directory")
	generateCmd.Flags().StringVarP(&config.Template.Path, "template", "t", "./templates", "templates file")
	generateCmd.MarkFlagRequired("output")

	// Create store templates
	config.Templates.Init("templates", config.Template.Path)
}
