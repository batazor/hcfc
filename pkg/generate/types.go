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
	ApiVersion  string
	AppVersion  string
}

type Deployment struct {
	Replicas     int
	Image        Image
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

type Service struct {
	Type string
	Port int
}

type Ingress struct {
	Enable bool
	Domain string
}
