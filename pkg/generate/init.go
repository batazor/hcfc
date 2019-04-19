package generate

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
)

func (p *Project) Init() error {
	err := p.isConfirm("Create a new config")
	if err != nil {
		return err
	}

	err = p.setChartConfig()
	if err != nil {
		return err
	}

	pipelinePromptui := []configPromptui{
		{
			isCreateMessage: "Create a new secret",
			isAddMessage:    "Add a secret",
			setConfigMethod: p.setSecretConfig,
		},
		{
			isCreateMessage: "Create a new deployment",
			isAddMessage:    "Add a deployment",
			setConfigMethod: p.setDeploymentConfig,
		},
		{
			isCreateMessage: "Create a new service",
			isAddMessage:    "Add a service",
			setConfigMethod: p.setServiceConfig,
		},
		{
			isCreateMessage: "Create a new ingress",
			isAddMessage:    "Add a ingress",
			setConfigMethod: p.setIngressConfig,
		},
	}

	// Create resource
	for _, cfg := range pipelinePromptui {
		errCreate := p.isConfirm(cfg.isCreateMessage)
		if errCreate == nil {
			fmt.Println("errCreate", errCreate)
			for {
				if errCreateConfig := cfg.setConfigMethod(); errCreateConfig != nil {
					return errCreateConfig
				}

				if isAddResource := p.isConfirm(cfg.isAddMessage); isAddResource != nil {
					break
				}
			}
		}
	}

	// Save result
	err = p.saveConfig()
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) isConfirm(label string) error {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		return errors.New("Enter 'Y' for continue\n")
	}

	return nil
}

func (p *Project) addPort() (Port, error) {
	var err error
	port := Port{}

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

func (p *Project) saveConfig() error {
	d, err := yaml.Marshal(&p)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("deploy.yaml", d, 0644)
	if err != nil {
		return err
	}

	return nil
}
