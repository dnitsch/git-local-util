[![Go Report Card](https://goreportcard.com/badge/github.com/dnitsch/git-local-util)](https://goreportcard.com/report/github.com/dnitsch/git-local-util)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=bugs)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=coverage)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)

# git-local-util

Collection of local git operations for increased productivity :shrug:

## Installation

Major platform binaries [here](https://github.com/dnitsch/git-local-util/releases)

*nix binary

```bash
curl -L https://github.com/dnitsch/git-local-util/releases/latest/download/git-local-util-linux -o git-local-util
```

MacOS binary

```bash
curl -L https://github.com/dnitsch/git-local-util/releases/latest/download/git-local-util-darwin -o git-local-util
```

```bash
chmod +x git-local-util
sudo mv git-local-util /usr/local/bin
```

Windows

```pwsh
iwr -Uri "https://github.com/dnitsch/git-local-util/releases/latest/download/git-local-util-windows" -OutFile "git-local-util"
```

Download specific version:

```bash
curl -L https://github.com/dnitsch/git-local-util/releases/download/v0.1.0/git-local-util-`uname -s` -o git-local-util
```

## Usage

`git-local-util --help`

`git-local-util migrate --help`

## Migrate

Migrates `remote "origin"` from one Url to another - useful when a repo or a large collection of repos are being moved to another URL - either within the same provider or another git compliant provider.

### Example

Simple migration

```bash
git-local-util migrate -d "./test" -f "find/old-origin-part" -r "replace/origin-part" --verbose
```

#### find

Find argument can be any continuos string i.e. `singleword` or `url/part/1`

#### replace

Replace argument can be any continuos string i.e. `singleword` or `url/part/1`
