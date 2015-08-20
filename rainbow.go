package main

import (
	"fmt"
	"net/http"

	"github.com/yosssi/ace"
)

func handler(w http.ResponseWriter, r *http.Request) {

	tpl, err := ace.Load("rainbow", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateErr := tpl.Execute(w, map[string]string{"Url": r.URL.Path[1:]})
	if templateErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Println("Serving " + r.URL.Path[1:] + " for " + r.RemoteAddr)
	}
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Starting server!")
	http.ListenAndServe(":8888", nil)
}
