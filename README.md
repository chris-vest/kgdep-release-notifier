# kgdep-release-notfier

[![Go Report Card](https://goreportcard.com/badge/github.com/chris-vest/kgdep-release-notifier)](https://goreportcard.com/report/github.com/chris-vest/kgdep-release-notifier)

Program to get annotations with the key `repo` and print the values to a file. 

## Next steps

Either:

1) Fork [release notifier](https://github.com/justwatchcom/github-releases-notifier) and add to its codebase to allow option to read configuration from file
2) Run as a sidecar, allow [release notifier](https://github.com/justwatchcom/github-releases-notifier) to pick it up - either configMap or reading from file on shared volume