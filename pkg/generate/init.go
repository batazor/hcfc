package generate

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
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

func (p *Project) setDeploymentConfig() error {
	// Init deployment
	newDeployment := Deployment{}

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

func (p *Project) setServiceConfig() error {
	// Init deployment
	newService := Service{}

	// Set type
	setTypeService := promptui.Select{
		Label: "Type",
		Items: []string{"ClusterIP", "NodePort"},
	}
	_, typeService, err := setTypeService.Run()
	if err != nil {
		return err
	}
	newService.Type = typeService

	// Binding to deployment
	listDeployment := []string{}
	for index := range p.Deployment {
		listDeployment = append(listDeployment, fmt.Sprintf("deployment-%s", strconv.Itoa(index)))
	}
	bindingToDeployment := promptui.Select{
		Label: "Binding to deployment",
		Items: listDeployment,
	}
	_, binding, err := bindingToDeployment.Run()
	if err != nil {
		return err
	}

	indexDeployment, _ := strconv.Atoi(binding)
	deployment := p.Deployment[indexDeployment]

	// binding metadata
	newService.Metadata.Labels = deployment.Metadata.Labels
	// binding selector
	newService.Selector = deployment.Metadata.Labels
	// binding port
	for _, port := range deployment.Ports {
		newService.Ports = append(newService.Ports, port)
	}

	// Append new service
	p.Service = append(p.Service, newService)

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

func (p *Project) setIngressConfig() error {
	// Init ingress
	newIngress := Ingress{}

	prompt := promptui.Prompt{
		Label:     "Enable",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		newIngress.Enable = false
	}

	newIngress.Enable = true

	// Set domain
	prompt = promptui.Prompt{
		Label:   "Domain",
		Default: "example.com",
	}
	domain, err := prompt.Run()
	if err != nil {
		return err
	}
	newIngress.Domain = domain

	// Append new service
	p.Ingress = append(p.Ingress, newIngress)

	return nil
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
