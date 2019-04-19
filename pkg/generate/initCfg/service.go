package initCfg

import (
	"fmt"
	"github.com/batazor/hcfc/pkg/generate"
	"github.com/manifoldco/promptui"
	"strconv"
)

func (p *Project) setServiceConfig() error {
	// Init service
	newService := generate.Service{}

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
