package generate

type BuildConfig struct {
	Output   string
	Filename string
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
}

type Deployment struct {
	Image Image
}

type Image struct {
	Repository string
	Tag        string
}

type Service struct {
	Type string
}

type Ingress struct {
	Enable bool
	Domain string
}
