# hcfc

generate Helm Chart from yaml config

### Getting start

```
go get -u github.com/batazor/hcfc

hcfc generate -o ./mychart -f ./example/values.yaml
```

**value.yaml**

```

projectName: example

deployment:
  scale: 1
  image:
    repository: batazor/example
    tag: latest

service:
  type: NodePort
  port: 5672

ingress:
  enable: true
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
  - [x] `-f` path to values file `values.yaml`
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

#### v1.1.0 dogfooding

- [ ] Use chart generator :-)

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

- [ ] Interactive create config `values.yaml`
