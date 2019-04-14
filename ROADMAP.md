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

- [x] Use chart generator :-)
- [x] Generate `values.yaml`

#### v1.2.0 Improve

- [ ] Generate `README.md`
  - [ ] description
  - [ ] table with ENV variable (name, default value)
- [ ] Skip comments in template file
- [ ] Improve Ingress template
- [ ] Use ENV CI_COMMIT_TAG
  - [ ] Dockerfile
  - [ ] GitHub Action

### v1.3.0

- [ ] Add linters
- [ ] Use k8s/helm structure

### v1.3.1

- [ ] Add test create `deployment.yaml`

### v1.3.2

- [ ] Add test create `deployment`
- [ ] Add test create `service`
- [ ] Add test create `values.yaml`

#### v1.4.0 Monitoring

- [ ] Support prometheus
  - [ ] Add template
    - [ ] Ping
    - [ ] healthcheck
- [ ] Dashboard for grafana

#### v2.0.0 Support Giltab

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

#### v2.0.1 dogfooding

- [ ] Use gitlab generator :-)

#### v3.0.0 Support GitHub Action

- [ ] Update command `init`
  - [ ] Add type `GitHub`

#### v3.0.1 dogfooding

- [ ] Use gitlab generator :-)

#### v4.0.0

- [ ] Support dependencies chart