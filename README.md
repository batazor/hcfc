# hcfc

generate Helm Chart from yaml config

### Getting start

```
go get -u github.com/batazor/hcfc/cmd/hcfc

hcfc init // Create deploy.yaml
hcfc generate -o ./mychart -f deploy.yaml
```

**value.yaml**

```
chart:
  name: hcfc
  description: ""
  version: 0.1.0
 
deployment:
  image:
    repository: alpine
    tag: latest
 
service:
  type: ClusterIP

ingress:
  enable: true
  domain: example.com
```

### Function template

1. ignore
1. toToml
1. toYaml
1. fromYaml
1. toJson
1. fromJson
1. include
1. required
1. tpl


### ROADMAP

#### v1.0 first version

- [x] Command
  - [x] `generate` is a command generation helm chart
  - [x] `-o` output directory
  - [x] `-f` path to values file `deploy.yaml`
  - [x] `-t` template directory
- [x] Generation simple chart from templates
  - [x] Chart.yaml
  - [x] deployment.yaml
  - [x] service.yaml
  - [x] add support Helm parsing

#### v1.0.1 Add CI/Refactoring

- [x] Add CI
  - [x] Build binary file
  - [x] Build docker image
  - [x] GitHub Action

#### v1.1.0 easy create deploy config

- [x] Use custom logger (for formatting logs)
- [x] Add command `init` - Generate a new `deploy.yaml` config
  - [x] Confirm create a new config
  - [x] Write name, description, version  
  - [x] Generation base deployment, service ,ingress (optional)

#### v1.1.1 dogfooding

- [ ] Use chart generator :-)
- [ ] Generate `values.yaml`

#### v1.2 Support Giltab

- [ ] Add command `init`
  - [ ] Add type `GitLab`
  - [ ] Generate Dockerfile (as plugins)
    - [ ] NodeJS
    - [ ] Simple HTML
    - [ ] Golang
    - [ ] Yii
- [ ] Generation gitlab-ci.yaml
  - [ ] Jobs:
    - [ ] Build Dockerfile
    - [ ] Push to registry
    - [ ] Generate Helm chart
    - [ ] Deploy chart
- [ ] Support ENV
- [ ] Add test example

#### v1.2.1 dogfooding

- [ ] Use gitlab generator :-)

#### v1.3 Support GitHub Action

- [ ] Update command `init`
  - [ ] Add type `GitHub`

#### v1.3.1 dogfooding

- [ ] Use gitlab generator :-)

#### v1.4

- [ ] Support dependencies chart

#### v1.5 Monitoring

- [ ] Support prometheus
  - [ ] Add template
    - [ ] Ping
    - [ ] healthcheck
- [ ] Dashboard for grafana

#### v2.0

- [ ] Interactive create config `deploy.yaml`
