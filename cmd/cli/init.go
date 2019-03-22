package cli

import (
	"github.com/batazor/hcfc/pkg/generate"
	"github.com/spf13/cobra"
	"log"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize new configuration",
	Run:   initConfig,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initConfig(cmd *cobra.Command, args []string) {
	newProject := generate.Project{}

	err := newProject.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
}
