package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(rw, "Hello World"); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatalln(http.ListenAndServe(":9090", nil))
}
