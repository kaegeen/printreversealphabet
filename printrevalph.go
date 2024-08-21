package main

import (
	"fmt"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i := len(alphabet) - 1; i >= 0; i-- {
		fmt.Print(string(alphabet[i]))
	}

}
