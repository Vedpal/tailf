package mtime

import (
	"os"
	"time"
)

// Get returns the Last Access Time for the given FileInfo
func Get(fi os.FileInfo) time.Time {
	return mtime(fi)
}

// Stat returns the Last Access Time for the given filename
func Stat(name string) (time.Time, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return time.Time{}, err
	}
	return mtime(fi), nil
}
