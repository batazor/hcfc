workflow "Build and Publish" {
  on = "release"
  resolves = [
    "Docker Publish",
    "release darwin/amd64",
    "release windows/amd64",
    "release linux/amd64",
  ]
}

action "Docker Tag" {
  needs = ["Build"]
  uses = "actions/docker/tag@master"
  args = "hcfc batazor/hcfc --no-latest"
}

action "Publish Filter" {
  uses = "actions/bin/filter@master"
  args = "tag"
}

action "Build" {
  needs = ["Publish Filter"]
  uses = "actions/docker/cli@master"
  args = "build -t hcfc ."
}

action "Docker Login" {
  needs = ["Publish Filter"]
  uses = "actions/docker/login@master"
  secrets = [
    "DOCKER_PASSWORD",
    "DOCKER_USERNAME",
  ]
}

action "Docker Publish" {
  needs = ["Docker Tag", "Docker Login"]
  uses = "actions/docker/cli@master"
  args = "push batazor/hcfc"
}

action "release darwin/amd64" {
  uses = "batazor/actions/golang/release@master"
  env = {
    GOOS = "darwin"
    GOARCH = "amd64"
    CMD_PATH = "cmd/hcfc/main.go"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release windows/amd64" {
  uses = "batazor/actions/golang/release@master"
  env = {
    GOOS = "windows"
    GOARCH = "amd64"
    CMD_PATH = "cmd/hcfc/main.go"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release linux/amd64" {
  uses = "batazor/actions/golang/release@master"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
    CMD_PATH = "cmd/hcfc/main.go"
  }
  secrets = ["GITHUB_TOKEN"]
}
