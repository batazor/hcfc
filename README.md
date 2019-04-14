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

1. ignore `{{ ignore .Values.Cat }} => {{ .Values.Cat }}`
1. text `{{ text .Values.Cat }} => .Values.Cat`
1. toToml
1. toYaml
1. fromYaml
1. toJson
1. fromJson
1. include
1. required
1. tpl
