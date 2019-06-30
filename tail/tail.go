package tail

import (
	"strings"
	"tailf/conn"
	"tailf/file"
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
