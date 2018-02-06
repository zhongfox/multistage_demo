// github.com/bigwhite/experiments/multi_stage_image_build/isomorphism/httpserver.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Println("  -> hello docker multi-stage")
	w.Write([]byte("Welcome to this website!\n"))
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Webserver start....")
	fmt.Println("  -> listen on port:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
