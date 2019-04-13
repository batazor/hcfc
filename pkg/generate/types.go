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
}

type Chart struct {
	Name        string
	Description string
	Version     string
	ApiVersion  string
	AppVersion  string
}

type Deployment struct {
	Name         string
	Replicas     int
	Image        Image
	Ports        []Port
	Resources    interface{}
	NodeSelector interface{}
	Affinity     interface{}
	Tolerations  interface{}
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
	Type string
	Port int
}

type Ingress struct {
	Enable bool
	Domain string
}
