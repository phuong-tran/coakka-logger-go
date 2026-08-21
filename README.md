# CoAkka Logger Go Connector

<p align="center">
  <img src="https://raw.githubusercontent.com/phuong-tran/coakka-samples/main/docs/assets/brand/coakka-logo.png" alt="CoAkka" width="480">
</p>

[![CI](https://github.com/phuong-tran/coakka-logger-go/actions/workflows/go-ci.yml/badge.svg)](https://github.com/phuong-tran/coakka-logger-go/actions/workflows/go-ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/phuong-tran/coakka-logger-go.svg)](https://pkg.go.dev/github.com/phuong-tran/coakka-logger-go)
[![Version](https://img.shields.io/badge/version-v1.2.6-blue)](https://github.com/phuong-tran/coakka-logger-go/tree/v1.2.6)
[![Release](https://img.shields.io/badge/release-v1.2.6-informational)](https://github.com/phuong-tran/coakka-logger-go/releases/tag/v1.2.6)
[![License: file-scoped](https://img.shields.io/badge/license-file--scoped-blue)](PACKAGE-LICENSE.md)
[![Funding](https://img.shields.io/badge/funding-Ko--fi-ff5f5f)](https://ko-fi.com/phuongnamtran)

Go module:

```sh
go get github.com/phuong-tran/coakka-logger-go@v1.2.6
```

This package is the Go connector for the standalone CoAkka native logger core.
It embeds native logger generation `1.2.1+f50756ebff0d` for macOS, Linux, and
Windows.
The module compatibility floor is Go `1.18`; use a currently supported Go
release for production builds. The unchanged Linux native payload requires
glibc `2.34` or newer independently of the Go version. That Linux generation
self-reports `git=unknown`; its source checkpoint remains pinned and verified
through the native manifest and payload digests.

Public package links:

| Link | Purpose |
| --- | --- |
| [pkg.go.dev/coakka-logger-go](https://pkg.go.dev/github.com/phuong-tran/coakka-logger-go@v1.2.6) | Go API reference for the current module version. |
| [GitHub Release v1.2.6](https://github.com/phuong-tran/coakka-logger-go/releases/tag/v1.2.6) | Source module release with bundled native libraries. |
| [Logger sample](https://github.com/phuong-tran/coakka-samples/tree/main/logger/go/basic) | Runnable bounded logger sample. |
| [Compatibility matrix](https://github.com/phuong-tran/coakka-publish/blob/main/docs/compatibility-matrix.md) | Current native generation and package-manager status. |

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
go get github.com/phuong-tran/coakka-logger-go@v1.2.6
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
go test ./...
```

If the local Go toolchain is managed by Homebrew on macOS, this workspace has
also been tested with:

```sh
export GOROOT=/opt/homebrew/opt/go/libexec
export PATH="$GOROOT/bin:$PATH:$HOME/go/bin"
```

This public module repository is already the exported Go module root. It does
not contain the internal release-packaging scripts used by the central CoAkka
release workspace.

The embedded native logger libraries live under:

```text
native/<platform>/libcoakka_logger_core-<native-package-version>.<suffix>
```

To verify the module as a consumer, create a clean module and install the
released tag:

```bash
mkdir coakka-logger-go-consumer
cd coakka-logger-go-consumer
go mod init coakka-logger-go-consumer
go get github.com/phuong-tran/coakka-logger-go@v1.2.6
```

Run the official sample for an end-to-end package check:

```bash
git clone https://github.com/phuong-tran/coakka-samples.git
cd coakka-samples
bash run.sh logger go basic
```

Release packaging, native staging, checksum capture, and module export are
owned by the central release pipeline and the public artifact surface:

- `coakkaCoreNativeDev`
- `coakka-publish`
- `coakka-samples`

## License

**Free for application use, including commercial and production use.**

Connector source, generated bindings, type declarations, examples, and package
documentation use the [Apache License, Version 2.0](https://github.com/phuong-tran/coakka-samples/blob/main/LICENSE).
Bundled native files use the [CoAkka Native Artifact License 1.2](https://github.com/phuong-tran/coakka-samples/blob/main/NATIVE-LICENSE.md).
Those native terms permit ordinary application and SaaS use but require a
separate agreement to sell or offer CoAkka itself as managed runtime or
infrastructure.

See [CoAkka Package Licensing](https://github.com/phuong-tran/coakka-samples/blob/main/docs/package-licensing.md)
for the file-scope map. This repository also carries offline `LICENSE`,
`NATIVE-LICENSE.md`, `PACKAGE-LICENSE.md`, and `NOTICE` copies.
