package generate

import (
	"github.com/manifoldco/promptui"
	"os"
	"path/filepath"
)

func (p *Project) setChartConfig() error {
	// Set name (default value: current dir)
	defaultNameProject, err := os.Getwd()
	if err != nil {
		return err
	}

	_, file := filepath.Split(defaultNameProject)

	prompt := promptui.Prompt{
		Label:   "Project name",
		Default: file,
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	p.Chart.Name = result

	// Set description
	prompt = promptui.Prompt{
		Label: "Description",
	}

	result, err = prompt.Run()
	if err != nil {
		return err
	}

	p.Chart.Description = result

	// 4. Set version (default version)
	prompt = promptui.Prompt{
		Label:   "Version",
		Default: "0.1.0",
	}

	result, err = prompt.Run()
	if err != nil {
		return err
	}

	p.Chart.Version = result

	return nil
}
