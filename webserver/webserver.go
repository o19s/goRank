package goRank

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/events", EventHandler)
	r.HandleFunc("/searches/{search}", SearchHandler)
	r.HandleFunc("/init", InitHandler)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*/*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	port := os.Getenv("GORANK_PORT")
	if port == "" {
		port = "8000"
	}
	address := os.Getenv("GORANK_ADDR")
	address = address + ":" + port
	fmt.Println("listening on:", address)

	srv := &http.Server{
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(handlers.RecoveryHandler()(r)),
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "goRank!")
}

func InitHandler(w http.ResponseWriter, r *http.Request) {
	InitStorage()
	fmt.Fprint(w, "Maybe?")
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	summary := make(map[string]int)

	events, err := FindEventsForSearch(vars["search"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	for _, element := range events {
		summary[element.Item] = summary[element.Item] + 1
	}
	js, err := json.Marshal(summary)
	fmt.Fprint(w, string(js))
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	// ignore hugh writes, DDOS bad!
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &event); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	Save(event)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
