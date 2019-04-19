package generate

import "github.com/manifoldco/promptui"

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
