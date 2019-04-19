package initCfg

import "github.com/batazor/hcfc/pkg/generate"

type Project struct {
	generate.Project
}

type configPromptui struct {
	isCreateMessage string
	isAddMessage    string
	setConfigMethod func() error
}
