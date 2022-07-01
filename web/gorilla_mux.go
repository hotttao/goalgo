package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	p := map[string]string{"name": "tsong"}
	json.NewEncoder(w).Encode(p)
}

func GetGorillaMux() {
	r := mux.NewRouter()
	r.HandleFunc("/products/{key}", HelloHandler)
	r.HandleFunc("/articles/{category}/", HelloHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", HelloHandler)

	s := r.Host("www.example.com").Subrouter()
	s.HandleFunc("/products/{key}", HelloHandler)

	s = r.PathPrefix("/products").Subrouter()
	// "/products/"
	s.HandleFunc("/", HelloHandler)
	// "/products/{key}/"
	s.HandleFunc("/{key}/", HelloHandler)
	// "/products/{key}/details"
	s.HandleFunc("/{key}/details", HelloHandler)

}
