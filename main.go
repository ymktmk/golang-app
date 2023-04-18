package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	hellowHandler := func(w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "Hello, World\n")
	}
	http.HandleFunc("/", hellowHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
