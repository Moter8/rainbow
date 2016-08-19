package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/yosssi/ace"
)

var i = 1

func handler(w http.ResponseWriter, r *http.Request) {

	tpl, err := ace.Load("rainbow", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rand.Seed(time.Now().UnixNano())
	messages := []string{
		"You are awesome,",
		"Pure awesomeness",
		"Coolest person alive:",
		"Literally Awesome:",
		"The amazing",
		"YOLO 420 -",
		"Knows how to tie their shoes:",
		"Makes delicious cakes:",
	}
	selectedMessage := messages[rand.Intn(len(messages))]

	templateErr := tpl.Execute(w, map[string]string{"Message": selectedMessage, "Counter": strconv.Itoa(i), "Url": r.URL.Path[1:]})
	if templateErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Println("Serving " + selectedMessage + " " + r.URL.Path[1:] + " for " + r.RemoteAddr)
		if r.URL.Path[1:] != "favicon.ico" {
			i++
		}
	}
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Starting server!")
	http.ListenAndServe(":8888", nil)
}
