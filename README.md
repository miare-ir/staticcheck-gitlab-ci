# staticcheck-gitlab-ci

This repo is an adaptor that converts the [staticcheck](https://pkg.go.dev/honnef.co/go/tools/cmd/staticcheck) output to a format that is recognized by the gitlab-ci.


## Installation
```bash
go install github.com/miare-ir/staticcheck-gitlab-ci@latest
```

## Usage
This application expects the staticcheck result as json in stdin and outputs the converted result as json. 


After installation run the below at the desired directory:
```bash
staticcheck -f json ./... | staticcheck-gitlab-ci
```

Check the example for a full integration. 

## Example
Somewhere in your `gitlab-ci.yml` you'll need to add sth like this:

```yaml
stages:
  - checks

code_quality:
  stage: checks
  script:
    - go install honnef.co/go/tools/cmd/staticcheck@v0.4.2
    - go install github.com/miare-ir/staticcheck-gitlab-ci@latest
    - staticcheck -f json ./... | staticcheck-gitlab-ci > staticcheck-report.json
  allow_failure: true
  artifacts:
    when: always
    reports:
      codequality: staticcheck-report.json
```

By doing so whenever the pipeline is ran, the Code Quality Report is put on the merge request:

![img](https://raw.githubusercontent.com/miare-ir/staticcheck-gitlab-ci/main/screenshots/gitlab-ci-mr-example.png)
