# git-local-util

Collection of local git operations for increased productivity :shrug:

## Migrate

Migrates `remote "origin"` from one Url to another - useful when a repo or a large collection of repos are being moved to another URL - either within the same provider or another git compliant provider. 

### Example

Simple migration

```bash
git-local-util migrate -f "old-org-name" -r "new-org" --verbose
```
