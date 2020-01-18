package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

// 3.15, 3-20
func main() {
	src := "한글도메인"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", src, ascii)
}
