package mtime

import (
	"os"
	"syscall"
	"time"
)

func mtime(fi os.FileInfo) time.Time {
	return time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastWriteTime.Nanoseconds())
}
