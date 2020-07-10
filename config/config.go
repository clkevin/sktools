package config

var Version = "v1.0.0"

var SendCharByte = []byte{29, 10}

func SendChar() string {
	return string(SendCharByte)
}
