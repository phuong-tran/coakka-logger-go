// Package coakka_logger is the Go connector for the standalone CoAkka native
// logger core.
//
// Install the public module with:
//
//	go get github.com/phuong-tran/coakka-logger-go@v1.2.2
//
// CoAkka Logger gives a host application a small language-native logging API
// while the native core owns queueing, pressure behavior, drain semantics, sink
// behavior, and platform library loading.
//
// Start one logger handle per process or component that owns a logging lane,
// write records with category and message text, then close the handle during
// application shutdown.
//
//	log, err := logger.Start(logger.LoggerSpec{
//		SystemName: "sample",
//		MinLevel:   logger.LevelInfo,
//	}, "")
//	if err != nil {
//		panic(err)
//	}
//	defer log.Close()
//
//	_, _, _ = log.Info("sample", "hello from Go")
//
// Public repositories:
//
//   - https://github.com/phuong-tran/coakka-logger-go
//   - https://github.com/phuong-tran/coakka-runtime-go
//   - https://github.com/phuong-tran/coakka-publish
//   - https://github.com/phuong-tran/coakka-samples
package coakka_logger
