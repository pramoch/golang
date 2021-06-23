package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// #1
	// bs := make([]byte, 19999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// #2
	// io.Copy(os.Stdout, resp.Body)

	// #3
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Len = ", len(bs))

	return len(bs), nil
}
