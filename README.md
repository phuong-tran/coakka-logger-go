# CoAkka Logger Go Connector

This package is the Go connector lane for the standalone CoAkka native logger
core.

It keeps the Go-facing API small:

- resolve the native logger shared library
- start and stop one logger handle
- write `category + message` text records
- manually drain records for tests and embedding use cases
- read native info/config/stats snapshots

The native core still owns queueing, pressure behavior, sink behavior, and
lifecycle state. The Go connector only owns host-language ergonomics and dynamic
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
logger/go/coakka-logger-go-1.2.1.tar.gz
```

Public Go module export:

```sh
bash scripts/export-module-repo.sh /tmp/coakka-logger-go-module
```

The exported directory is the root of module
`github.com/phuong-tran/coakka-logger-go`. After that public repository exists
and is tagged as `v1.2.1`, consumers can use `go get` without a local
`replace`.

The archive embeds native logger libraries under:

```text
native/<platform>/libcoakka_logger_core-<native-package-version>.<suffix>
```
