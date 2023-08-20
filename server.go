package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/haruyama480/go1210-request-streaming-failed/static"
)

func main() {
	var port = flag.Int("port", 10000, "port")
	var tls = flag.Bool("tls", false, "use tls")
	flag.Parse()

	handlerfn := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("request recv: %s %s\n", req.Method, req.URL.Path)

		if strings.Contains(req.URL.Path, ".wasm") {
			w.Header().Set("content-type", "application/wasm")
		}

		if req.URL.Path == "/greet" {
			fmt.Fprintln(w, "hello")
			return
		}

		http.FileServer(http.FS(static.Assets)).ServeHTTP(w, req)
	})

	addr := fmt.Sprintf("localhost:%d", *port)

	var err error
	log.Printf("start serving")
	if *tls {
		log.Printf("https://%s\n", addr)
		err = http.ListenAndServeTLS(addr, "./insecure/cert.pem", "./insecure/key.pem", handlerfn)
	} else {
		log.Printf("http://%s\n", addr)
		err = http.ListenAndServe(addr, handlerfn)
	}
	if err != nil {
		panic(err)
	}
}
