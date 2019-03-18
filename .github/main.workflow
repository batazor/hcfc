workflow "Build and Publish" {
  on = "push"
  resolves = "Publish"
}

action "Lint" {
  uses = "actions/action-builder/shell@master"
  args = "lint"
}

action "Test" {
  uses = "actions/action-builder/shell@master"
  args = "test"
}

action "build docker image" {
  needs = ["Lint", "Test"]
  uses = "actions/action-builder/docker@master"
  args = "build"
}

action "Publish Filter" {
  needs = ["Build"]
  uses = "actions/bin/filter@master"
  args = "branch master"
}

action "Docker Login" {
  needs = ["Publish Filter"]
  uses = "actions/docker/login@master"
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "Publish" {
  needs = ["Docker Login"]
  uses = "actions/action-builder/docker@master"
  runs = "make"
  args = "publish"
}
