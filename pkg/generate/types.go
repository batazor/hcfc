package generate

import "github.com/batazor/hcfc/pkg/templates"

type BuildConfig struct {
	Filename  string
	Template  Template
	Templates templates.Template
}

type Template struct {
	Filename string
	Path     string
	Output   string
	Values   Project
	Template string
}

type Project struct {
	Chart      Chart
	Deployment []Deployment
	Service    []Service
	Ingress    []Ingress
	Secret     []Secret
}

type Chart struct {
	Name        string
	Description string
	Version     string
	ApiVersion  string
	AppVersion  string
}

type Deployment struct {
	Metadata     Metadata
	ENV          map[string]interface{}
	Replicas     int
	Image        Image
	Ports        []Port
	Resources    interface{}
	NodeSelector interface{}
	Affinity     interface{}
	Tolerations  interface{}
}

type Metadata struct {
	Name        string
	Labels      map[string]string
	Annotations map[string]string
}

type Image struct {
	Repository string
	Tag        string
	PullPolicy string
}

type Port struct {
	Name     string
	Port     int
	Protocol string
}

type Service struct {
	Metadata Metadata
	Type     string
	Ports    []Port
	Selector map[string]string
}

type Ingress struct {
	Enabled  bool
	Metadata Metadata
	Domain   []Domain
}

type Domain struct {
	Host       []string
	SecretName string
	Rules      []Rules
}

type Rules struct {
	Host string
	Path []Path
}

type Path struct {
	Path    string
	Backend Backend
}

type Backend struct {
	ServiceName string
	ServicePort int
}

type Secret struct {
	Metadata Metadata
	Data     map[string]string // map[fileName]fileValue
}
