package utils

func PrintAsBytes(content string) {
	bytes := []byte(content)
	for e := range bytes {
		print(bytes[e], " ")
	}
	println("")
}
