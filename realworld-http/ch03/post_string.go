package main

import (
	"log"
	"net/http"
	"strings"
)

// 3.8 POST 메서드로 임의의 바디 전송, 예제 3-9
func main() {
	reader := strings.NewReader("text")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
