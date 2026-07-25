package coakka_logger

import (
	"fmt"
	"sync"
)

type Logger struct {
	bindings *nativeBindings
	handle   nativeHandle
	spec     LoggerSpec

	mu      sync.Mutex
	stopped bool
	closed  bool
}

func Start(spec LoggerSpec, loggerLibPath string) (*Logger, error) {
	spec = spec.withDefaults()
	resolvedPath, err := ResolveLoggerLibrary(loggerLibPath)
	if err != nil {
		return nil, err
	}
	bindings, err := openNativeBindings(resolvedPath)
	if err != nil {
		return nil, err
	}
	if abi := bindings.getAbiVersion(); abi != LoggerCoreABIVersion {
		bindings.close()
		return nil, fmt.Errorf("unexpected logger ABI version: expected %d got %d", LoggerCoreABIVersion, abi)
	}
	handle, err := bindings.create(spec)
	if err != nil {
		bindings.close()
		return nil, err
	}
	if err := bindings.start(handle); err != nil {
		bindings.destroy(handle)
		bindings.close()
		return nil, err
	}
	return &Logger{bindings: bindings, handle: handle, spec: spec}, nil
}

func ReadInfo(loggerLibPath string) (LoggerInfoSnapshot, error) {
	resolvedPath, err := ResolveLoggerLibrary(loggerLibPath)
	if err != nil {
		return LoggerInfoSnapshot{}, err
	}
	bindings, err := openNativeBindings(resolvedPath)
	if err != nil {
		return LoggerInfoSnapshot{}, err
	}
	defer bindings.close()
	return bindings.readInfo()
}

func (l *Logger) Config() (LoggerConfigSnapshot, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.requireOpenLocked(); err != nil {
		return LoggerConfigSnapshot{}, err
	}
	return l.bindings.config(l.handle)
}

func (l *Logger) Stats() (LoggerStatsSnapshot, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.requireOpenLocked(); err != nil {
		return LoggerStatsSnapshot{}, err
	}
	return l.bindings.stats(l.handle)
}

func (l *Logger) IsEnabled(level int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.closed {
		return false
	}
	return l.bindings.isEnabled(l.handle, level)
}

func (l *Logger) IsEnabledForCategory(category string, level int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.closed {
		return false
	}
	return l.bindings.isEnabledForCategory(l.handle, category, level)
}

func (l *Logger) Log(level int, category string, message string) (uint64, bool, error) {
	if !l.IsEnabledForCategory(category, level) {
		return 0, false, nil
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.requireOpenLocked(); err != nil {
		return 0, false, err
	}
	sequence, err := l.bindings.log(l.handle, level, category, message)
	if err != nil {
		return 0, false, err
	}
	return sequence, true, nil
}

func (l *Logger) Trace(category string, message string) (uint64, bool, error) {
	return l.Log(LevelTrace, category, message)
}

func (l *Logger) Debug(category string, message string) (uint64, bool, error) {
	return l.Log(LevelDebug, category, message)
}

func (l *Logger) Info(category string, message string) (uint64, bool, error) {
	return l.Log(LevelInfo, category, message)
}

func (l *Logger) Warn(category string, message string) (uint64, bool, error) {
	return l.Log(LevelWarn, category, message)
}

func (l *Logger) Error(category string, message string) (uint64, bool, error) {
	return l.Log(LevelError, category, message)
}

func (l *Logger) Fatal(category string, message string) (uint64, bool, error) {
	return l.Log(LevelFatal, category, message)
}

func (l *Logger) Poll() (*LoggerRecordSnapshot, error) {
	return l.AwaitNext(0)
}

func (l *Logger) AwaitNext(timeoutMs uint32) (*LoggerRecordSnapshot, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.requireOpenLocked(); err != nil {
		return nil, err
	}
	return l.bindings.readNext(l.handle, timeoutMs, l.spec.CategoryCapacity, l.spec.MessageCapacity)
}

func (l *Logger) Stop() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.closed || l.stopped {
		return nil
	}
	if err := l.bindings.stop(l.handle); err != nil {
		return err
	}
	l.stopped = true
	return nil
}

func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.closed {
		return nil
	}
	var stopErr error
	if !l.stopped {
		stopErr = l.bindings.stop(l.handle)
		l.stopped = true
	}
	l.bindings.destroy(l.handle)
	l.bindings.close()
	l.closed = true
	return stopErr
}

func (l *Logger) requireOpenLocked() error {
	if l.closed {
		return fmt.Errorf("logger is closed")
	}
	return nil
}
