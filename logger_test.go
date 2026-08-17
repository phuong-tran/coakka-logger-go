package coakka_logger

import (
	"path/filepath"
	"runtime"
	"testing"
)

func stagedLoggerLib(t *testing.T) string {
	t.Helper()
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("could not resolve test source path")
	}
	libName := "libcoakka_logger_core.so"
	if runtime.GOOS == "darwin" {
		libName = "libcoakka_logger_core.dylib"
	} else if runtime.GOOS == "windows" {
		libName = "libcoakka_logger_core.dll"
	}
	return filepath.Join(
		filepath.Dir(sourceFile),
		"..",
		"staging",
		"native",
		LoggerNativePackageVersion,
		PlatformID(runtime.GOOS, runtime.GOARCH),
		libName,
	)
}

func TestPlatformID(t *testing.T) {
	tests := []struct {
		goos string
		arch string
		want string
	}{
		{goos: "darwin", arch: "arm64", want: "macos-aarch64"},
		{goos: "macos", arch: "aarch64", want: "macos-aarch64"},
		{goos: "linux", arch: "amd64", want: "linux-x86_64"},
		{goos: "linux", arch: "x86_64", want: "linux-x86_64"},
		{goos: "windows", arch: "amd64", want: "windows-x86_64"},
		{goos: "windows", arch: "arm64", want: "windows-aarch64"},
	}
	for _, tt := range tests {
		if got := PlatformID(tt.goos, tt.arch); got != tt.want {
			t.Fatalf("PlatformID(%q, %q)=%q want %q", tt.goos, tt.arch, got, tt.want)
		}
	}
}

func TestLoggerResourceFileNames(t *testing.T) {
	if got, want := loggerResourceFileNames("linux"), []string{"libcoakka_logger_core-1.2.1+f50756ebff0d.so"}; len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("loggerResourceFileNames(linux)=%v want %v", got, want)
	}
	if got, want := loggerResourceFileNames("windows"), []string{"libcoakka_logger_core-1.2.1+f50756ebff0d.dll"}; len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("loggerResourceFileNames(windows)=%v want %v", got, want)
	}
}

func TestLoggerInfoAndManualDrain(t *testing.T) {
	libPath := stagedLoggerLib(t)
	if _, err := ResolveLoggerLibrary(libPath); err != nil {
		t.Skipf("logger native library not available: %v", err)
	}

	info, err := ReadInfo(libPath)
	if err != nil {
		t.Fatalf("ReadInfo failed: %v", err)
	}
	if info.ABIVersion != LoggerCoreABIVersion {
		t.Fatalf("abi version=%d want %d", info.ABIVersion, LoggerCoreABIVersion)
	}
	if info.GitCommit != LoggerNativeGitCommit {
		linuxGenerationWithoutEmbeddedCommit := runtime.GOOS == "linux" &&
			LoggerNativePackageVersion == "1.2.1+f50756ebff0d" &&
			info.GitCommit == "unknown"
		if !linuxGenerationWithoutEmbeddedCommit {
			t.Fatalf("git commit=%q want %q", info.GitCommit, LoggerNativeGitCommit)
		}
	}

	logger, err := Start(LoggerSpec{SystemName: "go-logger-test", MinLevel: LevelInfo}, libPath)
	if err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer logger.Close()

	if logger.IsEnabled(LevelDebug) {
		t.Fatalf("debug should be disabled by info min level")
	}
	if _, accepted, err := logger.Debug("test", "should stay disabled"); err != nil || accepted {
		t.Fatalf("disabled debug log accepted=%v err=%v", accepted, err)
	}
	sequence, accepted, err := logger.Info("test", "hello from go logger")
	if err != nil {
		t.Fatalf("Info failed: %v", err)
	}
	if !accepted {
		t.Fatalf("info log should be accepted")
	}
	record, err := logger.AwaitNext(1000)
	if err != nil {
		t.Fatalf("AwaitNext failed: %v", err)
	}
	if record == nil {
		t.Fatalf("expected logger record")
	}
	if record.Sequence != sequence {
		t.Fatalf("record sequence=%d want %d", record.Sequence, sequence)
	}
	if record.Level != LevelInfo || record.Category != "test" || record.Message != "hello from go logger" {
		t.Fatalf("unexpected record: %+v", record)
	}
	stats, err := logger.Stats()
	if err != nil {
		t.Fatalf("Stats failed: %v", err)
	}
	if stats.EmittedCount < 1 {
		t.Fatalf("emitted count=%d want >= 1", stats.EmittedCount)
	}
}
