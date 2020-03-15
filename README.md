# tftpdd

TFTP server daemon.

[![pipeline status](https://gitlab.com/pojntfx/tftpdd/badges/master/pipeline.svg)](https://gitlab.com/pojntfx/tftpdd/commits/master)

## Overview

`tftpd` is a management daemon with a CLI for [pin/tftp](https://github.com/pin/tftp), an excellent TFTP library for Go. It is intended to be used as the TFTP server in PXE boot setups, and as such all write capabilities have been removed. It is built of two components:

- `tftpdd`, an TFTP server management daemon with a gRPC interface
- `tftpdctl`, a CLI for `tftpdd`

## Installation

### Prebuilt Binaries

Prebuilt binaries are available on the [releases page](https://github.com/pojntfx/tftpdd/releases/latest).

### Go Package

A Go package [is available](https://pkg.go.dev/github.com/pojntfx/tftpdd).

### Docker Image

A Docker image is available on [Docker Hub](https://hub.docker.com/r/pojntfx/tftpdd).

### Helm Chart

A Helm chart is available in [@pojntfx's Helm chart repository](https://pojntfx.github.io/charts/).

## Usage

### Daemons

You may also set the flags by setting env variables in the format `TFTPDD_[FLAG]` (i.e. `TFTPDD_TFTPDD_CONFIGFILE=examples/tftpdd.yaml`) or by using a [configuration file](examples/tftpdd.yaml).

```bash
% tftpdd --help
tftpdd is the TFTP server management daemon.

Find more information at:
https://pojntfx.github.io/tftpdd/

Usage:
  tftpdd [flags]

Flags:
  -h, --help                           help for tftpdd
  -f, --tftpdd.configFile string       Configuration file to use.
  -l, --tftpdd.listenHostPort string   TCP listen host:port. (default "localhost:1340")
```

### Client CLIs

You may also set the flags by setting env variables in the format `TFTPD_[FLAG]` (i.e. `TFTPD_TFTPD_CONFIGFILE=examples/tftpd.yaml`) or by using a [configuration file](examples/tftpd.yaml).

```bash
% tftpdctl --help
tftpdctl manages tftpdd, the TFTP server management daemon.

Find more information at:
https://pojntfx.github.io/tftpdd/

Usage:
  tftpdctl [command]

Available Commands:
  apply       Apply a TFTP server
  delete      Delete one or more TFTP server(s)
  get         Get one or all TFTP server(s)
  help        Help about any command

Flags:
  -h, --help   help for tftpdctl

Use "tftpdctl [command] --help" for more information about a command.
```

## License

tftpdd (c) 2020 Felicitas Pojtinger

SPDX-License-Identifier: AGPL-3.0
