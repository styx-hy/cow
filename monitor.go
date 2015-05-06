package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	gohttp "net/http" // conflict with http variable in config.go
	_ "net/http/pprof"
	"runtime"
	"runtime/pprof"
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

func CountGoRoutine(w gohttp.ResponseWriter, r *gohttp.Request) {
	fmt.Fprintf(w, "%d", runtime.NumGoroutine())
}

func MonitorProfile(w gohttp.ResponseWriter, r *gohttp.Request) {
	prof := pprof.Lookup("goroutine")
	prof.WriteTo(w, 1)
}
func monitor() {
	go func() {
		log.Println(gohttp.ListenAndServe("localhost:6060", nil))
	}()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/numgoroutine", CountGoRoutine)
	router.HandleFunc("/profile", MonitorProfile)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)
	log.Fatal(gohttp.ListenAndServe(":8088", router))
}
