# Go Check Updates

`go-check-updates` upgrades your go.mod dependencies to the latest versions, ignoring specified versions.

## Quick Start

1. Install utility using either from [Homebrew](https://brew.sh) or [GitHub](https://github.com/fogfish/go-check-updates).

```bash
## Install using brew
brew install fogfish/tap/go-check-updates

## Alternatively, install from source code
go install github.com/fogfish/go-check-updates@latest
```

2. Run the command in your Golang repository to check for dependency updates

```bash
go-check-updates
```

3. Update dependencies to the latest versions

```bash
go-check-updates -u
```

4. Alternatively, you can update dependency and push changes to your git repository. The utility creates a new branch `go-update-deps`.

```bash
go-check-updates -u --push origin
```

5. Automate workflow with GitHub Actions to create a pull request every time `go-update-deps` branch is pushed. Craft GitHub Action with following command and install it into your workflow.

```bash
go-check-updates generate github > .github/workflows/update-deps.yml
```


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

`go-check-update` is the utility that simplify the workflow of running these commands.


## How To Contribute

`go-check-update` is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


## License

[![See LICENSE](https://img.shields.io/github/license/fogfish/go-check-updates.svg?style=for-the-badge)](LICENSE)
