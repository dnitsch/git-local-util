[![Go Report Card](https://goreportcard.com/badge/github.com/dnitsch/git-local-util)](https://goreportcard.com/report/github.com/dnitsch/git-local-util)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=bugs)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=dnitsch_git-local-util&metric=coverage)](https://sonarcloud.io/summary/new_code?id=dnitsch_git-local-util)

# git-local-util

Collection of local git operations for increased productivity :shrug:

## Migrate

Migrates `remote "origin"` from one Url to another - useful when a repo or a large collection of repos are being moved to another URL - either within the same provider or another git compliant provider. 

### Example

Simple migration

```bash
git-local-util migrate -f "old-org-name" -r "new-org" --verbose
```
