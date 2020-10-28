package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rsc.io/quote"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, quote.Go())
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
