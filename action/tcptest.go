package action

import (
	"sktools/tcp"
	"sktools/context"
	"time"
)

type Tcptest struct {
}

func (a *Tcptest) Action() {
	println("tcptest")
	for i := 0; i < context.Connection; i++ {
		go testConn()
	}
	println("sleep")
	time.Sleep(10 * time.Second)
}

func testConn() {
	println("open connection!!!")
	conn := tcp.Connection()
	index := 0
	for true {
		index ++
		tcp.Write(conn, context.Content)
		if index > 10 {
			break
		}
		time.Sleep(1* time.Second)
	}
}

func (a *Tcptest) Alias() (string, string) {
	return "tcptest", "send tcp test package"
}

func init() {
	action := &Tcptest{}
	key, _ := action.Alias()
	Register(key, action)
}
