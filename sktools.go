package main

import (
	"flag"
	"sktools/context"
	"sktools/action"
)

func main() {
	action.GetAction().Action()
}

func init() {
	var actionUsage =  action.Help()
	flag.StringVar(&context.Action, "a", "tcp",actionUsage)
	flag.StringVar(&context.Host, "host", "127.0.0.1", "host ip")
	flag.StringVar(&context.Port, "port", "80", "host port")
	flag.IntVar(&context.Connection, "c", 10, "connection ")
	flag.Parse()
}
