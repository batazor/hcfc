package initCfg

import (
	"fmt"
	"github.com/batazor/hcfc/pkg/generate"
	"github.com/manifoldco/promptui"
	"strconv"
)

func (p *Project) setIngressConfig() error {
	// Init ingress
	newIngress := generate.Ingress{
		Domain: []generate.Domain{},
	}

	// Enabled
	prompt := promptui.Prompt{
		Label:     "Enabled",
		IsConfirm: true,
	}
	_, err := prompt.Run()
	if err != nil {
		newIngress.Enabled = false
	}
	newIngress.Enabled = true

	for {
		Domain := generate.Domain{
			Rules: []generate.Rules{},
		}

		// Add host
		getHost := promptui.Prompt{
			Label:   "Host name",
			Default: "example.com",
		}
		host, err := getHost.Run()
		if err != nil {
			return err
		}
		Domain.Host = append(Domain.Host, host)

		// Add secretName
		// Binding to secret
		listSecret := []string{}
		for index := range p.Secret {
			listSecret = append(listSecret, fmt.Sprintf("secret-%s", strconv.Itoa(index)))
		}
		bindingToDeployment := promptui.Select{
			Label: "Binding to secret",
			Items: listSecret,
		}
		_, binding, err := bindingToDeployment.Run()
		if err != nil {
			return err
		}

		indexSecret, _ := strconv.Atoi(binding)
		secret := p.Secret[indexSecret]
		Domain.SecretName = secret.Metadata.Name

		// add rules
		// add path
		// bind to backendService
		for {
			rules := generate.Rules{}

			rules.Host = "TODO: select host"

			getPath := promptui.Prompt{
				Label:   "path",
				Default: "/",
			}
			path, err := getPath.Run()
			if err != nil {
				return err
			}

			rules.Path = []generate.Path{}

			rules.Path = append(rules.Path, generate.Path{
				Path: path,
				Backend: generate.Backend{
					ServiceName: "",
					ServicePort: 80,
				},
			})

			Domain.Rules = append(Domain.Rules, rules)

			isAddRule := p.isConfirm("Add a rule")
			if isAddRule != nil {
				break
			}
		}

		// Save domain
		newIngress.Domain = append(newIngress.Domain, Domain)

		if isAddResource := p.isConfirm("Add a domain"); isAddResource != nil {
			break
		}
	}

	// Append new service
	p.Ingress = append(p.Ingress, newIngress)

	return nil
}
