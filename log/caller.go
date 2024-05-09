package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	callerInitOnce     sync.Once
	minimumCallerDepth int
	logrusPackage      string
)

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func CallerPrettyfier(*runtime.Frame) (function string, file string) {
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "CallerPrettyfier") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	basePath := GetGoPath()

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)
		if pkg == logrusPackage {
			f, _ = frames.Next()
			return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, basePath), f.Line)
		}
		if pkg != "github.com/sirupsen/logrus" {
			return "", fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, basePath), f.Line)
		}
	}

	return "", ""
}

func GetGoPath() string {
	path := os.Getenv("BUILD_REPO_WS")
	if path != "" && path != " " {
		return path
	}
	path = os.Getenv("GOPATH")
	if path != "" && path != " " {
		return path
	}

	return ""
}
