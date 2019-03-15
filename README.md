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

#### v1.0

- [ ] Command
  - [ ] `generate` is a command generation helm chart
  - [ ] `-o` output directory
- [ ] Generation simple chart
  - [ ] Chart.yaml
  - [ ] values.yaml
  - [ ] deployment.yaml
  - [ ] service.yaml

#### v1.1

- [ ] Generation ingress
- [ ] Support dependencies chart

#### v1.2

- [ ] Generation gitlab-ci.yaml

#### v1.3

- [ ] Support custom template (go-modules? middleware?)
