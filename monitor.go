package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	gohttp "net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func Index(w gohttp.ResponseWriter, r *gohttp.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w gohttp.ResponseWriter, r *gohttp.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w gohttp.ResponseWriter, r *gohttp.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintf(w, "Todo show:", todoID)
}

func monitor() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)
	log.Fatal(gohttp.ListenAndServe(":8080", router))
}
