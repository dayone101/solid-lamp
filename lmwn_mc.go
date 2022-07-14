package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	fmt.Println(sd)

	asString := string(sd)
	fmt.Println(asString)

	for _, ch := range asString {
		whatIsIt = string(ch) + whatIsIt
	}

	fmt.Println(whatIsIt)
}
