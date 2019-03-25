package generate

import (
	"errors"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
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
		errDeployment := p.setDeploymentConfig()
		if errDeployment != nil {
			return errDeployment
		}
	}

	// Create service
	err = p.isConfirm("Create a new service")
	if err == nil {
		errService := p.setServiceConfig()
		if errService != nil {
			return errService
		}
	}

	// Create ingress
	err = p.isConfirm("Create a new ingress")
	if err == nil {
		errIngress := p.setIngressConfig()
		if errIngress != nil {
			return errIngress
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

	// Append new deployment
	p.Deployment = append(p.Deployment, newDeployment)

	return nil
}

func (p *Project) setServiceConfig() error {
	// Init deployment
	newService := Service{}

	// Set type
	prompt := promptui.Prompt{
		Label:   "Type",
		Default: "ClusterIP",
	}
	typeService, err := prompt.Run()
	if err != nil {
		return err
	}
	newService.Type = typeService

	// Append new service
	p.Service = append(p.Service, newService)

	return nil
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
