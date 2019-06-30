package file

import (
        "bytes"
        "log"
        "os"
        "strings"
	"webtailf/mtime"
)

var filename = "/var/log/wifi.log"
var bufSize = 1024

func SetFileName(name string) {
	filename = name
}

func ReadLastNLines(n int) ([]string, error) {
        curr := getFileSize()
        var end int
        count := n
        result := make([]byte, bufSize)
        for count >= 0 && curr != 0 {
                curr -= bufSize
                end = curr + bufSize - 1
                if curr < 0 {
                        curr = 0
                }
                buff := readNBytes(curr, end)
                result = append(buff, result...)
                count -= lineCount(buff)
        }
        lines := byteArrayToMultiLines(result)
	if len(lines) > n {
		l := len(lines)
		lines = lines[(l-n-1):]
	}
	return lines, nil
}

func MonitorFile(out chan []string) {
        size := getFileSize()
        last_modified, err := mtime.Stat(filename)
	if err != nil {
		log.Print("File-Stat-Error:", err)
	}

        go func() {
                for {
                        t, err := mtime.Stat(filename)
                        if err != nil {
				log.Print("File-Stat-Error:", err)
                                continue
                        }
                        if last_modified != t {
                                last_modified = t

                                // read file
                                newSize := getFileSize()
                                if newSize <= size {
                                        continue
                                }
                                content := readNBytes(size, newSize-1)
                                size = newSize
                                out <- byteArrayToMultiLines(content)
                        }
                }
        }()
}

func getFileSize() int {
        fileInfo, err := os.Stat(filename)
        if err != nil {
                log.Fatal("File-Stat-Error:", err)
        }
        return int(fileInfo.Size())
}

func lineCount(lines []byte) int {
        return bytes.Count(lines, []byte("\n"))
}

func byteArrayToMultiLines(bytes []byte) []string {
        lines := string(bytes)
        return strings.Split(lines, "\n")
}

func readNBytes(start int, end int) []byte {
        fh, _ := os.Open(filename)
        defer fh.Close()

        fh.Seek(int64(start), 0)
        size := end - start + 1
        buff := make([]byte, size)

        fh.Read(buff)
        return buff
}
