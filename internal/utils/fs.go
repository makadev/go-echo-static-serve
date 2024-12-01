package utils

import (
	"os"
	"strings"
)

func GetPathSeparator() string {
	delimString := string(os.PathSeparator)
	if len(delimString) != 1 {
		// we expect only 1 character
		panic("InternalError: 202304111456")
	}
	return delimString
}

func RemoveTrailingSeparators(fn string) string {
	return strings.TrimRight(fn, GetPathSeparator())
}

func RemoveLeadingSeparators(fn string) string {
	return strings.TrimLeft(fn, GetPathSeparator())
}

func AddTrailingSeparator(fn string) string {
	return RemoveTrailingSeparators(fn) + GetPathSeparator()
}

func PathExists(name string) (bool, bool) {
	stat, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, false
	}
	return true, stat.IsDir()
}

func FileExists(name string) bool {
	res, isDir := PathExists(name)
	return res && !isDir
}

func DirectoryExists(name string) bool {
	res, isDir := PathExists(name)
	return res && isDir
}
