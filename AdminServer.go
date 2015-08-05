package main

import (
    //"fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", adminHandler)
    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })
    panic(http.ListenAndServe(":6868", nil))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "admin/index.htm")
}