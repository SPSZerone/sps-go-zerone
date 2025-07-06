package os

import (
	"runtime"
)

func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

func IsFreeBSD() bool {
	return runtime.GOOS == "freebsd"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
