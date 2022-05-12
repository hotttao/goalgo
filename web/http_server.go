package web

import (
	"fmt"
	"log"
	"net/http"
)

func StartWebServer() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", "tsong")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
