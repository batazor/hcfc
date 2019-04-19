package generate

import (
	"errors"
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

	// Create secret
	err = p.isConfirm("Create a new secret")
	if err == nil {
		for {
			errSecret := p.setSecretConfig()
			if errSecret != nil {
				return errSecret
			}

			isAddResource := p.isConfirm("Add a secret")
			if isAddResource != nil {
				break
			}

		}
	}

	// Create deployment
	err = p.isConfirm("Create a new deployment")
	if err == nil {
		for {
			errDeployment := p.setDeploymentConfig()
			if errDeployment != nil {
				return errDeployment
			}

			isAddResource := p.isConfirm("Add a deployment")
			if isAddResource != nil {
				break
			}

		}
	}

	// Create service
	err = p.isConfirm("Create a new service")
	if err == nil {
		for {
			errService := p.setServiceConfig()
			if errService != nil {
				return errService
			}

			isAddResource := p.isConfirm("Add a service")
			if isAddResource != nil {
				break
			}

		}
	}

	// Create ingress
	err = p.isConfirm("Create a new ingress")
	if err == nil {
		for {
			errIngress := p.setIngressConfig()
			if errIngress != nil {
				return errIngress
			}

			isAddResource := p.isConfirm("Add a ingress")
			if isAddResource != nil {
				break
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
