// ConfigController.go
package ConfigController

import (
	//"fmt"
	"encoding/json"
	"net/http"
)

type config_struct struct {
	Test string
}

func handelConfig(w http.ResponseWriter, r *http.Request) {
	http.Handle("/Config/save", SaveConfig)
}

func SaveConfig(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c config_struct
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}
