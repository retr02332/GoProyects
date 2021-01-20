package main

import (
	"fmt"
	"io"
	"net/http"
)

type handler struct{}

func (handler) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return 0, nil
}

func main() {
	response, err := http.Get("https://facebook.com")
	if err != nil {
		fmt.Println(err)
	}
	hand := handler{}
	io.Copy(hand, response.Body)
}
