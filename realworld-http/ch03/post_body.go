package main

import (
	"log"
	"net/http"
	"os"
)

// 3.8 POST 메서드로 임의의 바디 전송, 예제 3-8
// curl -T post_body.go -H "Content-Type: text/plain" http://localhost:18888
func main() {
	file, err := os.Open("post_body.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
