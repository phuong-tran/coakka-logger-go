# CoAkka Logger Go Connector

Go module:

```sh
go get github.com/phuong-tran/coakka-logger-go@v1.2.2
```

This package is the Go connector for the standalone CoAkka native logger core.
It embeds native logger generation `1.2.1+f50756ebff0d` for macOS, Linux, and
Windows.

## New To CoAkka Logger

CoAkka Logger gives a host application a small language-native logging API
while the native core owns queueing, pressure behavior, drain semantics, sink
behavior, and platform library loading.

Use these public repositories to orient first:

- `https://github.com/phuong-tran/coakka-logger-go`
- `https://github.com/phuong-tran/coakka-runtime-go`
- `https://github.com/phuong-tran/coakka-publish`
- `https://github.com/phuong-tran/coakka-samples`

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
logger/go/coakka-logger-go-1.2.2.tar.gz
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
