//go:build wasm && js

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"syscall/js"
)

var (
	document js.Value
	logger   js.Value
)

func write(format string, args ...any) (n int, err error) {
	s := fmt.Sprintf(format, args...)
	node := document.Call("createElement", "div")
	node.Set("innerHTML", s)
	js.Value(logger).Call("appendChild", node)
	return len(s), nil
}

func init() {
	document = js.Global().Get("document")
	logger = document.Call("getElementById", "target")
}

func main() {
	res, err := http.Get("greet")
	if err != nil {
		write("GET request failed: %v", err)
	} else {
		bytes, _ := io.ReadAll(res.Body)
		res.Body.Close()
		write("GET request success: %s", bytes)
	}

	res, err = http.Post("greet", "text/plain", bytes.NewBufferString("dummy"))
	if err != nil {
		write("POST request failed %v", err)
	} else {
		bytes, _ := io.ReadAll(res.Body)
		res.Body.Close()
		write("GET request success: %s", bytes)
	}

	<-make(chan bool)
}
