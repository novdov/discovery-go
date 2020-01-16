package main

import (
	"log"
	"net/http"
	"net/url"
)

// curl -d test=value http://localhost:18888
func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
