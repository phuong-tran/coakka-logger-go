# CoAkka Logger Go Connector

[![CI](https://github.com/phuong-tran/coakka-logger-go/actions/workflows/go-ci.yml/badge.svg)](https://github.com/phuong-tran/coakka-logger-go/actions/workflows/go-ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/phuong-tran/coakka-logger-go.svg)](https://pkg.go.dev/github.com/phuong-tran/coakka-logger-go)
[![Version](https://img.shields.io/badge/version-v1.2.5-blue)](https://github.com/phuong-tran/coakka-logger-go/tree/v1.2.5)
[![Release](https://img.shields.io/badge/release-v1.2.5-informational)](https://github.com/phuong-tran/coakka-logger-go/releases/tag/v1.2.5)
[![License](https://img.shields.io/badge/license-Apache--2.0-green)](LICENSE)
[![Funding](https://img.shields.io/badge/funding-Ko--fi-ff5f5f)](https://ko-fi.com/phuongnamtran)

Go module:

```sh
go get github.com/phuong-tran/coakka-logger-go@v1.2.5
```

This package is the Go connector for the standalone CoAkka native logger core.
It embeds native logger generation `1.2.1+f50756ebff0d` for macOS, Linux, and
Windows.

## License

The Go connector source is Apache-2.0 licensed. The bundled native logger
libraries under `native/` use the CoAkka Public Artifact Preview terms in
`NATIVE-LICENSE.md`.

## New To CoAkka Logger

CoAkka Logger gives a host application a small language-native logging API
while the native core owns queueing, pressure behavior, drain semantics, sink
behavior, and platform library loading.

Use these public repositories to orient first:

| Repository | Use it for | Link |
| --- | --- | --- |
| `coakka-samples` | Runnable examples and code you can inspect first. | https://github.com/phuong-tran/coakka-samples |
| `coakka-publish` | Released packages, native archives, manifests, checksums, compatibility matrix, and release notes. | https://github.com/phuong-tran/coakka-publish |
| `coakka-logger-go` | Public Go module source for this package. | https://github.com/phuong-tran/coakka-logger-go |
| `coakka-runtime-go` | Public Go runtime module source. | https://github.com/phuong-tran/coakka-runtime-go |

Run the matching sample:

```sh
git clone https://github.com/phuong-tran/coakka-samples.git
cd coakka-samples
bash run.sh logger go basic
```

Read the deeper package docs:

- [Why CoAkka Logger matters](docs/coakka-logger.md)
- [CoAkka ecosystem map](docs/coakka-ecosystem.md)

Try the module without cloning any CoAkka repo:

```sh
mkdir coakka-logger-go-first-run
cd coakka-logger-go-first-run
go mod init coakka-logger-go-first-run
go get github.com/phuong-tran/coakka-logger-go@v1.2.5
```

```go
package main

import (
	"fmt"

	logger "github.com/phuong-tran/coakka-logger-go"
)

func main() {
	log, err := logger.Start(logger.LoggerSpec{
		SystemName: "first-user-logger",
		MinLevel:   logger.LevelInfo,
	}, "")
	if err != nil {
		panic(err)
	}
	defer log.Close()

	sequence, accepted, err := log.Info("first.user", `{"hello":"logger"}`)
	if err != nil {
		panic(err)
	}
	record, err := log.AwaitNext(1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(sequence, accepted, record.Category)
}
```

It keeps the Go-facing API small:

- resolve the native logger shared library
- start and stop one logger handle
- write `category + message` text records
- manually drain records for tests and embedding use cases
- read native info/config/stats snapshots

The native core still owns queueing, pressure behavior, sink behavior, and
lifecycle state. The Go connector owns host-language ergonomics and dynamic
library loading.

## Verify

From this directory:

```sh
bash scripts/stage-logger-natives.sh
go test ./...
bash scripts/smoke-packaged-package.sh
```

If the local Go toolchain is managed by Homebrew on macOS, this workspace has
also been tested with:

```sh
export GOROOT=/opt/homebrew/opt/go/libexec
export PATH="$GOROOT/bin:$PATH:$HOME/go/bin"
```

## Package

```sh
bash scripts/package-release.sh
```

The package archive is written to:

```text
logger/go/coakka-logger-go-1.2.5.tar.gz
```

Public Go module export:

```sh
bash scripts/export-module-repo.sh /tmp/coakka-logger-go-module
```

The exported directory is the root of public module
`github.com/phuong-tran/coakka-logger-go`.

The archive embeds native logger libraries under:

```text
native/<platform>/libcoakka_logger_core-<native-package-version>.<suffix>
```
