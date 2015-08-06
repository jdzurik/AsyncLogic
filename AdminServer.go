package main

import (
	//"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", adminHandler)
	http.HandleFunc("/Shared/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	panic(http.ListenAndServe(":6868", nil))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Admin/Index.htm")
}
