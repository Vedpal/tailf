package tail

import (
	"strings"
	"webtailf/conn"
	"webtailf/file"
)

func Filetail() {

	out := make(chan []string)
        file.MonitorFile(out)

        for {
                select {
                case lines := <-out:
                        content := strings.Join(lines, "\n")
                        conn.GetConnMgr().Broadcast(content)
                }
        }
}
