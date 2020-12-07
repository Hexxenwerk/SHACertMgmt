package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(rw, "Hello World"); err != nil {
			log.Fatal(err)
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := fmt.Fprintf(rw, "Data: %s\n", d); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatalln(http.ListenAndServe("127.0.0.1:9090", nil))
}
