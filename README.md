# hcfc

generate Helm Chart from yaml config

### Getting start

```
go get -u hcfc

hcfc generate -o ops/Helm value.yaml
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

### ROADMAP

#### v1.0 first version

- [x] Command
  - [x] `generate` is a command generation helm chart
  - [x] `-o` output directory
  - [x] `-f` path to values file `values.yaml`
  - [x] `-t` template directory
- [ ] Generation simple chart from templates
  - [x] Chart.yaml
  - [ ] deployment.yaml
  - [ ] service.yaml
  - [ ] values.yaml # copy `example/values.yaml` and delete `chart` (if f.Name() == 'values.yaml')
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
