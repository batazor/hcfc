package initCfg

import (
	"github.com/batazor/hcfc/pkg/generate"
	"github.com/manifoldco/promptui"
)

func (p *Project) setIngressConfig() error {
	// Init ingress
	newIngress := generate.Ingress{}

	prompt := promptui.Prompt{
		Label:     "Enable",
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		newIngress.Enabled = false
	}

	newIngress.Enabled = true

	// Set domain
	//prompt = promptui.Prompt{
	//	Label:   "Domain",
	//	Default: "example.com",
	//}
	//domain, err := prompt.Run()
	//if err != nil {
	//	return err
	//}
	//newIngress.Domain = domain

	// Append new service
	p.Ingress = append(p.Ingress, newIngress)

	return nil
}
