package tcp

import (
	"net"
	"sktools/context"
	"strconv"
)

func Connection() *net.TCPConn {
	port, _ := strconv.Atoi(context.Port)
	addr := net.TCPAddr{IP: net.ParseIP(context.Host), Port: port}
	conn, err := net.DialTCP("tcp", nil, &addr)
	if err != nil {
		panic(err)
	}
	conn.SetKeepAlive(true)
	return conn
}

func Write(conn *net.TCPConn, body string) {

	_, err := conn.Write([]byte(body))
	if err != nil {
		panic(err)
	}
}

func Read(conn *net.TCPConn) {
	var data = make([]byte, 2048)
	_, err := conn.Read(data)
	if err != nil {
		panic(err)
	}
	//println("count:", count)

	println(string(data))
}

func loopRead(conn *net.TCPConn) {
	for {
		Read(conn)
	}
}

func SyncLoopRead(conn *net.TCPConn) {
	go loopRead(conn)
}
