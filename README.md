[![build](https://github.com/martoc/ipstack-cli/actions/workflows/build.yml/badge.svg?branch=main&event=push)](https://github.com/martoc/ipstack-cli/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/martoc/ipstack-cli/branch/main/graph/badge.svg?token=S06JCJYGHM)](https://codecov.io/gh/martoc/ipstack-cli)
![Go Version](https://img.shields.io/github/go-mod/go-version/martoc/ipstack-cli/main)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![slack](https://img.shields.io/badge/slack-general-brightgreen.svg?logo=slack)](https://app.slack.com/messages/T8L8AAD3M/C8LBHLSVA)

# ipstack-cli

A simple CLI to get information about an IP address using the [ipstack](https://ipstack.com/) API.

## Installation

### Go

```sh
go install github.com/martoc/ipstack-cli@latest
```

## Usage

### Go installation

```sh
ipstack-cli -h
```

### Docker

```sh
export ACCESS_KEY=<from https://ipstack.com/dashboard>

docker run -it -e ACCESS_KEY=$ACCESS_KEY martoc/ipstack-cli:latest get-coordinates --ip "10.0.0.0"

docker run -it martoc/ipstack-cli:latest get-coordinates --ip "10.0.0.0" --access-key $ACCESS_KEY
```

[Documentation >>](./docs/index.md)
