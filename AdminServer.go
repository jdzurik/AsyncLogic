package main

import (
	//"fmt"
	"net/http"
	"AsyncLogic/Config"
	
)

func main() {
	
var config = ConfigObject.LoadConfig() 
	http.HandleFunc("/", adminHandler)
	http.HandleFunc("/Shared/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	panic(http.ListenAndServe(":"+config.Admin.Port, nil))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Admin/Index.htm")
}

