package coakka_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const loggerEnvVar = "COAKKA_LOGGER_LIB"

func ResolveLoggerLibrary(explicitPath string) (string, error) {
	if explicitPath != "" {
		return requireExisting(explicitPath, "explicit loggerLibPath")
	}
	if configured := os.Getenv(loggerEnvVar); configured != "" {
		return requireExisting(configured, "$"+loggerEnvVar)
	}
	if embedded, ok := resolvePackagedNative(); ok {
		return embedded, nil
	}
	for _, candidate := range searchCandidates() {
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}
	return "", fmt.Errorf(
		"native logger library was not found; set %s, pass loggerLibPath explicitly, or package one of %v under logger/go/native/%s",
		loggerEnvVar,
		loggerResourceFileNames(runtime.GOOS),
		PlatformID(runtime.GOOS, runtime.GOARCH),
	)
}

func PlatformID(goos string, goarch string) string {
	return normalizeOS(goos) + "-" + normalizeArch(goarch)
}

func normalizeOS(goos string) string {
	switch goos {
	case "darwin", "macos":
		return "macos"
	case "linux":
		return "linux"
	case "windows":
		return "windows"
	default:
		panic(fmt.Sprintf("unsupported os=%s; supported platforms are macOS, Linux, and Windows", goos))
	}
}

func normalizeArch(goarch string) string {
	switch goarch {
	case "arm64", "aarch64":
		return "aarch64"
	case "amd64", "x86_64":
		return "x86_64"
	default:
		panic(fmt.Sprintf("unsupported arch=%s; supported architectures are aarch64 and x86_64", goarch))
	}
}

func loggerResourceFileNames(goos string) []string {
	versionedBase := "libcoakka_logger_core-" + LoggerNativePackageVersion
	switch normalizeOS(goos) {
	case "macos":
		return []string{versionedBase + ".dylib"}
	case "linux":
		return []string{versionedBase + ".so"}
	case "windows":
		return []string{versionedBase + ".dll"}
	default:
		panic("unsupported platform")
	}
}

func loggerLocalFileNames(goos string) []string {
	switch normalizeOS(goos) {
	case "macos":
		return []string{"libcoakka_logger_core.dylib", "libcoakka_logger_core.so"}
	case "linux":
		return []string{"libcoakka_logger_core.so"}
	case "windows":
		return []string{"libcoakka_logger_core.dll"}
	default:
		panic("unsupported platform")
	}
}

func resolvePackagedNative() (string, bool) {
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", false
	}
	packageDir := filepath.Dir(sourceFile)
	platform := PlatformID(runtime.GOOS, runtime.GOARCH)
	for _, fileName := range loggerResourceFileNames(runtime.GOOS) {
		candidate := filepath.Join(packageDir, "native", platform, fileName)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, true
		}
	}
	return "", false
}

func searchCandidates() []string {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	names := loggerLocalFileNames(runtime.GOOS)
	roots := []string{
		filepath.Join(cwd, "lib"),
		filepath.Join(cwd, "logger", "go", "lib"),
	}
	var out []string
	for _, root := range roots {
		for _, name := range names {
			out = append(out, filepath.Join(root, name))
		}
	}
	return out
}

func requireExisting(path string, source string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(abs); err != nil {
		return "", fmt.Errorf("%s does not exist: %s", source, abs)
	}
	return abs, nil
}
