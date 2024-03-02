package main

import (
	"io"
	"net/http"
)

func main() {
	r, err := http.Get("http://localhost:8080/cotacao")
	if err != nil {
	}

	defer r.Body.Close()

	test, _ := io.ReadAll(r.Body)

	println(string(test))
}
