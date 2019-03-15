# hcfc

generate Helm Chart from yaml config

### Getting start

```
go get -u hcfc

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

### Function

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

#### v1.1 Add CI/Refactoring

- [ ] Use logger `zap`
- [ ] Add CI
  - [ ] Build binary file
  - [ ] Build docker image
  - [ ] GitHub Action
  - [ ] Coverage

#### v1.2 Support Giltab

- [ ] Generation gitlab-ci.yaml
- [ ] Support ENV
- [ ] Add test example

#### v1.3

- [ ] Support dependencies chart
