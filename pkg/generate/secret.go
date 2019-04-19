package generate

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func (p *Project) setSecretConfig() error {
	// Init secret
	newSecret := Secret{}

	// Set metadata
	newSecret.Metadata.Labels = map[string]string{}
	newSecret.Metadata.Labels["app.kubernetes.io/name"] = p.Chart.Name
	newSecret.Metadata.Labels["app.kubernetes.io/instance"] = "{{ .Release.Name }}"
	newSecret.Metadata.Labels["app.kubernetes.io/managed-by"] = "Tiller"
	newSecret.Metadata.Labels["helm.sh/chart"] = fmt.Sprintf("%s-%s", p.Chart.Name, p.Chart.Version)

	// Set name
	setName := promptui.Prompt{
		Label:   "Secret name",
		Default: p.Chart.Name,
	}
	name, err := setName.Run()
	if err != nil {
		return err
	}
	newSecret.Metadata.Name = name

	// Set file
	newSecret.Data = map[string]string{}

	isAddFile := p.isConfirm("Add a file")
	if isAddFile == nil {
		for {
			prompt := promptui.Prompt{
				Label:   "File name",
				Default: "test.txt",
			}
			fileName, err := prompt.Run()
			if err != nil {
				return err
			}
			prompt = promptui.Prompt{
				Label:   "File name",
				Default: "secretValue",
			}
			fileValue, err := prompt.Run()
			if err != nil {
				return err
			}

			newSecret.Data[fileName] = fileValue

			isAddFile = p.isConfirm("Add a file")
			if isAddFile != nil {
				break
			}
		}
	}

	// Append new secret
	p.Secret = append(p.Secret, newSecret)

	return nil
}
