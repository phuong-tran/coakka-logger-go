# Consuming The Go Logger Package

The current Go logger artifact is a source package plus embedded native logger
libraries. The module path is already fixed as
`github.com/phuong-tran/coakka-logger-go`, but the public Go module repository
must exist and be tagged before users can install it with `go get`.

Until that repository is opened, extract the archive and use a local `replace`.

Example consumer `go.mod`:

```go
module my-logger-consumer

go 1.22

require github.com/phuong-tran/coakka-logger-go v0.0.0

replace github.com/phuong-tran/coakka-logger-go => ./coakka-logger-go-1.2.1
```

Example:

```go
package main

import logger "github.com/phuong-tran/coakka-logger-go"

func main() {
	log, err := logger.Start(logger.LoggerSpec{
		SystemName: "sample",
		MinLevel:   logger.LevelInfo,
	}, "")
	if err != nil {
		panic(err)
	}
	defer log.Close()

	_, _, _ = log.Info("sample", "hello from Go")
}
```

Library resolution order:

- explicit path passed to `Start(spec, loggerLibPath)` or `ReadInfo(loggerLibPath)`
- `$COAKKA_LOGGER_LIB`
- packaged native library under `native/<platform>/`
- local fallback candidates under `lib/` and `logger/go/lib/`

Current packaged platforms:

- `macos-aarch64`
- `linux-aarch64`
- `linux-x86_64`

After the public module repository is opened and tagged, the local `replace`
line goes away and consumers should use:

```sh
go get github.com/phuong-tran/coakka-logger-go@v1.2.1
```
