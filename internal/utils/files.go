package utils

import (
	"runtime"
)

func GetPath() string {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	return path
}
