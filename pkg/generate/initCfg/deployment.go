package initCfg

import (
	"fmt"
	"github.com/batazor/hcfc/pkg/generate"
	"github.com/manifoldco/promptui"
	"strconv"
)

func (p *Project) setDeploymentConfig() error {
	// Init deployment
	newDeployment := generate.Deployment{}

	// Set metadata
	newDeployment.Metadata.Labels = map[string]string{}
	newDeployment.Metadata.Labels["app.kubernetes.io/name"] = p.Chart.Name
	newDeployment.Metadata.Labels["app.kubernetes.io/instance"] = "{{ .Release.Name }}"
	newDeployment.Metadata.Labels["app.kubernetes.io/managed-by"] = "Tiller"
	newDeployment.Metadata.Labels["helm.sh/chart"] = fmt.Sprintf("%s-%s", p.Chart.Name, p.Chart.Version)

	// Set repository
	prompt := promptui.Prompt{
		Label:   "Repository",
		Default: "alpine",
	}
	repository, err := prompt.Run()
	if err != nil {
		return err
	}
	newDeployment.Image.Repository = repository

	// Set tag
	prompt = promptui.Prompt{
		Label:   "Tag",
		Default: "latest",
	}
	tag, err := prompt.Run()
	if err != nil {
		return err
	}
	newDeployment.Image.Tag = tag

	// Create ENV
	newDeployment.ENV = map[string]interface{}{}
	isAddENV := p.isConfirm("Add a ENV")
	if isAddENV == nil {
		for {
			setNameENV := promptui.Prompt{
				Label:   "Name ENV",
				Default: "test",
			}
			nameENV, err := setNameENV.Run()
			if err != nil {
				return err
			}

			setValueENV := promptui.Prompt{
				Label:   fmt.Sprintf("Value ENV (name: %s)", nameENV),
				Default: "test",
			}
			valueENV, err := setValueENV.Run()
			if err != nil {
				return err
			}

			newDeployment.ENV[nameENV] = valueENV

			isAddENV = p.isConfirm("Add a ENV")
			if isAddENV != nil {
				break
			}
		}
	}

	// Create port
	isAddResource := p.isConfirm("Add a port")
	if isAddResource == nil {
		for {
			port, err := p.addPort()
			if err != nil {
				return err
			}

			newDeployment.Ports = append(newDeployment.Ports, port)

			isAddResource = p.isConfirm("Add a port")
			if isAddResource != nil {
				break
			}
		}
	}

	// Append new deployment
	p.Deployment = append(p.Deployment, newDeployment)

	return nil
}

func (p *Project) addPort() (generate.Port, error) {
	var err error
	port := generate.Port{}

	// Set portName
	setPortName := promptui.Prompt{
		Label:   "Port name",
		Default: "http",
	}
	port.Name, err = setPortName.Run()
	if err != nil {
		return port, err
	}

	// Set port
	setPortInt := promptui.Prompt{
		Label:   "Port",
		Default: "80",
	}
	port1, err := setPortInt.Run()
	if err != nil {
		return port, err
	}
	port.Port, err = strconv.Atoi(port1)
	if err != nil {
		return port, err
	}

	// Set protocol
	protocol := promptui.Select{
		Label: "protocol",
		Items: []string{"TCP", "UDP"},
	}
	_, port.Protocol, err = protocol.Run()
	if err != nil {
		return port, err
	}

	return port, nil
}
