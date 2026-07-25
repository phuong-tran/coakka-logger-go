# Consuming The Go Logger Package

The current Go logger package is a public Go module plus embedded native
logger libraries:

```sh
go get github.com/phuong-tran/coakka-logger-go@v1.2.5
```

Example consumer `go.mod`:

```go
module my-logger-consumer

go 1.22

require github.com/phuong-tran/coakka-logger-go v1.2.5
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
- `windows-aarch64`
- `windows-x86_64`

Repository:

- `https://github.com/phuong-tran/coakka-logger-go`
