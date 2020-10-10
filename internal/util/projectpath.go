package util

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return path.Join(basepath, "..", "..")
}
