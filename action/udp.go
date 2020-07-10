package action

import (
	"sktools/context"
)

type Udp struct {
}

func (a *Udp) Action() {
	println("Udp action:", context.Host)
}

func (a *Udp) Alias() (string, string) {
	return "udp", "send Udp package"
}


func init() {
	action := &Udp{}
	key, _ := action.Alias()
	Register(key, action)

}
