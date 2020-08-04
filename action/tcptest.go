package action

import (
	"sktools/tcp"
	"sktools/context"
	"time"
)

type Tcptest struct {
}

var channels = make(chan string, context.Connection)

func (a *Tcptest) Action() {
	println("tcptest")
	for i := 0; i < context.Connection; i++ {
		go testConn()
		//time.Sleep(time.Duration(context.ConnIntervalTime) * time.Second)
		time.Sleep(time.Duration(context.ConnIntervalTime)*time.Millisecond)
	}
	//println("sleep")
	//time.Sleep((time.Duration(context.Count)+5)*time.Second )
	for i := 0; i < context.Connection; i++ {
		var str = <-channels
		println(str)
	}
}

func testConn() {
	conn := tcp.Connection()
	println(conn.LocalAddr().String(), "open connection!!!")
	index := 0
	for true {
		index ++
		println(conn.LocalAddr().String(), "write:", context.Content)
		tcp.Write(conn, context.Content)
		if context.Count != -1 && index >= context.Count {
			break
		}
		sleepTime := time.Duration(context.KeepaliveIntervalTime) * time.Second
		time.Sleep(sleepTime)
	}
	conn.Close()
	println("close connection")
	channels <- conn.LocalAddr().String() + " close connection"
}

func (a *Tcptest) Alias() (string, string) {
	return "tcptest", "send tcp test package"
}

func init() {
	action := &Tcptest{}
	key, _ := action.Alias()
	Register(key, action)
}
