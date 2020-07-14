package action

import "sktools/context"

type Action interface {
	Action()
	Alias() (string, string)
}

var actions = make(map[string]Action)

func GetAction() Action {
	var result = actions[context.Action]
	if result == nil {
		println("invalid param ", context.Action)
		result = actions["help"]
	}
	return result
}

func Register(key string, action Action) {
	actions[key] = action
}

func Help() string {
	var result string
	for e := range actions {
		k, u := actions[e].Alias()
		result = result + " [" + k + " - " + u + "]" + "\r\n"
	}
	return result
}
