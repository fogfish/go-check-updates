# Go Check Updates

`go-check-updates` upgrades your go.mod dependencies to the latest versions, ignoring specified versions.

## Quick Start


## Inspiration

Golang support easy way to inspect new version of modules:

```bash
go list -u \
  -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' \
  -m all

## github.com/aws/aws-cdk-go/awscdk/v2: v2.87.0 -> v2.103.1
## github.com/aws/aws-sdk-go-v2: v1.19.0 -> v1.21.2
## github.com/aws/aws-sdk-go-v2/config: v1.18.28 -> v1.19.1
## ...
```

With following command, it is possible to update go.mod
```bash
go list -u \
  -f "{{if (and (not (or .Main .Indirect)) .Update)}}go get -d {{.Path}}@{{.Update.Version}} ; {{end}}" \
  -m all | sh

## go get -d github.com/aws/aws-cdk-go/awscdk/v2@v2.103.1 ;
## go get -d github.com/aws/aws-sdk-go-v2@v1.21.2 ;
## go get -d github.com/aws/aws-sdk-go-v2/config@v1.19.1 ;
## ...
```

The command line utility just simple way of running these commands



