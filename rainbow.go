package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/yosssi/ace"
)

func handler(w http.ResponseWriter, r *http.Request) {

	tpl, err := ace.Load("rainbow", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rand.Seed(time.Now().UnixNano())
	messages := []string{"You are awesome,", "Pure awesomeness", "Coolest person alive:", "Literally Awesome:", "The amazing"}
	selectedMessage := messages[rand.Intn(len(messages))]

	templateErr := tpl.Execute(w, map[string]string{"Message": selectedMessage, "Url": r.URL.Path[1:]})
	if templateErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Println("Serving " + selectedMessage + " " + r.URL.Path[1:] + " for " + r.RemoteAddr)
	}
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Starting server!")
	http.ListenAndServe(":8888", nil)
}
