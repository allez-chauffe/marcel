package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("swagger-ui"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":3000", nil)
}
