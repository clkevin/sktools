package action

import (
	"sktools/context"
	"sktools/tcp"
	"os"
	"bufio"
	"strings"
	"sktools/config"
	"time"
)

type Tcp struct {
}

func (a *Tcp) Action() {
	now := time.Now().String()
	println(now, "  tcp action:", context.Host)

	conn := tcp.Connection()
	tcp.SyncLoopRead(conn)
	
	reader := bufio.NewReader(os.Stdin)

	var content string
	var begin = true
	for {
		if begin {
			//sendChar := string(config.SendCharByte[:])
			println("input tcp package,if you want to send please input ^]", " :")
			begin = false
			content = ""
		}
		text, _ := reader.ReadString('\n')
		if text == config.SendChar() {
			request := strings.Replace(content, "\n", "\r\n", -1)
			//httpBody := "GET /spring-tomcat/keepalive HTTP/1.1\r\nHost: 127.0.0.1:8080\r\n\r\n";
			tcp.Write(conn, request)
			begin = true
		} else {
			content = content + text
		}
	}
}

func (a *Tcp) Alias() (string, string) {
	return "tcp", "send tcp package"
}

func init() {
	action := &Tcp{}
	key, _ := action.Alias()
	Register(key, action)
}
